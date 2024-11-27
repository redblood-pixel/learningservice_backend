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
	CreateSession(sessionID string, userId int, refreshTTL time.Duration) error
	RemoveSession(sessionID string) (domain.RefreshToken, error)
}

type Dictionary interface {
	GetAll() []domain.Word
	Get(wordID int) (domain.Word, error)
	Create(word domain.CreateWordRequest) error
	Update(word domain.Word) error
	Delete(wordID int) error
}

type Group interface {
	GetAll() []domain.Group
	Get(groupID int) (domain.Group, error)
	Create(word domain.CreateGroupRequest) (int, error)
	Update(word domain.Group) error
	Delete(groupID int) error
	WordsInGroup(groupID int) ([]domain.Word, error)
	GroupsOfUser(userID int) ([]domain.Group, error)
}

type Repository struct {
	Users Users
	Dict  Dictionary
	Group Group
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
