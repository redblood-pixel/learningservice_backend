package repository

import (
	"fmt"
	"time"

	"github.com/redblood-pixel/learning-service-go/pkg/domain"
	postgres_repo "github.com/redblood-pixel/learning-service-go/pkg/repository/postgres"
)

type Users interface {
	Create(username, email, password string) (int, error)
	FindUserByEmail(username, password string) (int, error)
	CreateSession(sessionId string, userId int, refreshTTL time.Duration) error
	RemoveSession(sessionId string) (domain.RefreshToken, error)
}

type Dictionary interface {
	GetAll() []domain.Word
	Get(id int) (domain.Word, error)
	Create(word domain.CreateWordRequest) error
	Update(word domain.Word) error
	Delete(wordId int) error
}

type Repository struct {
	Users Users
	Dict  Dictionary
}

func NewRepository(db interface{}) (*Repository, error) {
	switch db := db.(type) {
	case *domain.Database:
		return &Repository{
			Users: postgres_repo.NewUsersRepository(db),
			Dict:  postgres_repo.NewDictRepository(db),
		}, nil
	default:
		return nil, fmt.Errorf("can not create repositories: invalid db type - %T", db)
	}

}
