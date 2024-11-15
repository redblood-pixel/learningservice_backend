package gorm_repo

import (
	"time"

	"github.com/redblood-pixel/learning-service-go/pkg/domain"
)

type UsersRepository struct {
	db *domain.DatabaseORM
}

func NewUsersRepository(db *domain.DatabaseORM) *UsersRepository {
	return &UsersRepository{db: db}
}

func (ur *UsersRepository) Create(username, email, password string) (int, error) {
	return 0, nil
}

func (ur *UsersRepository) FindUserByEmail(username, password string) (int, error) {
	return 0, nil
}

func (ur *UsersRepository) CreateSession(sessionId string, userId int, refreshTTL time.Duration) error {
	return nil
}

func (ur *UsersRepository) RemoveSession(sessionId string) (domain.RefreshToken, error) {
	return domain.RefreshToken{}, nil
}
