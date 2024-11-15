package repository

import (
	"fmt"
	"time"

	"github.com/redblood-pixel/learning-service-go/pkg/domain"
	gorm_repo "github.com/redblood-pixel/learning-service-go/pkg/repository/gorm"
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

type Repositories struct {
	Users Users
	Dict  Dictionary
}

func NewRepositories(db interface{}) (*Repositories, error) {
	switch db := db.(type) {
	case *domain.Database:
		return &Repositories{
			Users: postgres_repo.NewUsersRepository(db),
			Dict:  postgres_repo.NewDictRepository(db),
		}, nil
	case *domain.DatabaseORM:
		return &Repositories{
			Users: gorm_repo.NewUsersRepository(db),
			Dict:  gorm_repo.NewDictRepository(db),
		}, nil
	default:
		return nil, fmt.Errorf("can not create repositories: invalid db type - %T", db)
	}

}
