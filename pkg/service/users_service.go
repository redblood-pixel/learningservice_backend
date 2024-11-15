package service

import (
	"time"

	"github.com/redblood-pixel/learning-service-go/internal/hash"
	"github.com/redblood-pixel/learning-service-go/internal/tokenutil"
	"github.com/redblood-pixel/learning-service-go/pkg/domain"
	"github.com/redblood-pixel/learning-service-go/pkg/repository"
	"github.com/sirupsen/logrus"
)

type UsersService struct {
	UsersRepository repository.Users
	TokenManager    *tokenutil.TokenManager
	Hasher          *hash.PasswordHasher
}

func NewUsersService(repo repository.Users, hasher *hash.PasswordHasher, tokenManager *tokenutil.TokenManager) *UsersService {
	return &UsersService{
		UsersRepository: repo,
		Hasher:          hasher,
		TokenManager:    tokenManager,
	}
}

func (us *UsersService) SignUp(signupInput domain.SignupInput) (domain.TokensResponse, error) {

	var err error
	signupInput.Password, err = us.Hasher.GetHash(signupInput.Password)
	if err != nil {
		return domain.TokensResponse{}, err
	}

	id, err := us.UsersRepository.Create(signupInput.Username, signupInput.Email, signupInput.Password)
	if err != nil {
		return domain.TokensResponse{}, err
	}

	return us.createSession(id)
}

func (us *UsersService) SignIn(signinInput domain.SigninInput) (domain.TokensResponse, error) {

	var err error
	signinInput.Password, err = us.Hasher.GetHash(signinInput.Password)
	if err != nil {
		return domain.TokensResponse{}, err
	}

	id, err := us.UsersRepository.FindUserByEmail(signinInput.Email, signinInput.Password)
	if err != nil {
		return domain.TokensResponse{}, err
	}

	return us.createSession(id)
}

func (us *UsersService) Refresh(refreshInput domain.RefreshInput) (domain.TokensResponse, error) {

	// Get session and delete it if exists
	res, err := us.UsersRepository.RemoveSession(refreshInput.RefreshToken)
	if err != nil {
		logrus.Error("Refresh userservice error: remove session")
		return domain.TokensResponse{}, err
	}

	// Validate session
	if res.ExpiresAt.Before(time.Now()) {
		logrus.Error("Refresh userservice error: session expired")
		return domain.TokensResponse{}, domain.ErrNotAuthorized
	}

	// create new session
	return us.createSession(res.UserId)
}

func (us *UsersService) createSession(id int) (domain.TokensResponse, error) {
	var (
		err    error
		tokens domain.TokensResponse
	)

	tokens.AccessToken, err = us.TokenManager.CreateAccessToken(id)
	if err != nil {
		return domain.TokensResponse{}, err
	}

	tokens.RefreshToken = us.TokenManager.CreateRefreshToken()

	err = us.UsersRepository.CreateSession(tokens.RefreshToken, id, us.TokenManager.RefreshTTL())
	if err != nil {
		return domain.TokensResponse{}, err
	}

	return tokens, nil
}
