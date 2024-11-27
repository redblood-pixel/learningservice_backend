package service

import (
	"github.com/redblood-pixel/learning-service-go/internal/hash"
	"github.com/redblood-pixel/learning-service-go/internal/tokenutil"
	"github.com/redblood-pixel/learning-service-go/pkg/domain"
	"github.com/redblood-pixel/learning-service-go/pkg/repository"
)

type Users interface {
	SignUp(user domain.SignupInput) (domain.TokensResponse, error)
	SignIn(user domain.SigninInput) (domain.TokensResponse, error)
	Refresh(user domain.RefreshInput) (domain.TokensResponse, error)
}

type Dictionary interface {
	GetAllWords() []domain.Word
	GetWord(id int) (domain.Word, error)
	CreateWord(word domain.CreateWordRequest) error
	UpdateWord(word domain.Word) error
	DeleteWord(wordId int) error
}

type Group interface {
	// models
	GetAllGroups() []domain.Group
	GetGroup(groupID int) (domain.Group, error)
	CreateGroup(domain.CreateGroupRequest) (int, error)
	UpdateGroup(group domain.Group) error
	DeleteGroup(groupID int) error
	GetWordsInGroup(groupID int) ([]domain.Word, error)
	GetGroupsOfUser(userID int) ([]domain.Group, error)
}

type Dependencies struct {
	Repos        *repository.Repository
	Hasher       *hash.PasswordHasher
	TokenManager *tokenutil.TokenManager
}

type Service struct {
	Users
	Dictionary
	Group
}

func NewService(d Dependencies) *Service {
	usersService := NewUsersService(d.Repos.Users, d.Hasher, d.TokenManager)
	dictService := NewDictService(d.Repos.Dict)
	groupService := NewGroupService(d.Repos.Group)
	return &Service{
		Users:      usersService,
		Dictionary: dictService,
		Group:      groupService,
	}
}
