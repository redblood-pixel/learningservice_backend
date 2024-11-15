package repository

import (
	"time"

	"github.com/redblood-pixel/learning-service-go/pkg/domain"
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

func NewRepositories(db *domain.Database) *Repositories {
	return &Repositories{
		Users: NewUsersRepository(db),
		Dict:  NewDictRepository(db),
	}
}
