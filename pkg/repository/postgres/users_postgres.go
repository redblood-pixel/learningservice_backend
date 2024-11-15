package postgres_repo

import (
	"fmt"
	"time"

	"github.com/redblood-pixel/learning-service-go/pkg/domain"
	"github.com/sirupsen/logrus"
)

type UsersRepository struct {
	db *domain.Database
}

func NewUsersRepository(db *domain.Database) *UsersRepository {
	return &UsersRepository{db: db}
}

func (ur *UsersRepository) Create(username, email, password string) (int, error) {
	var id int

	q := fmt.Sprintf("INSERT INTO users (username, email, password_hash) VALUES ('%s', '%s', '%s') RETURNING id",
		username, email, password)

	row := ur.db.QueryRow(q)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (ur *UsersRepository) FindUserByEmail(email, password string) (int, error) {
	var id int

	q := fmt.Sprintf("SELECT (id) FROM users WHERE email='%s' AND password_hash='%s'", email, password)

	if err := ur.db.Get(&id, q); err != nil {
		return 0, err
	}
	return id, nil
}

func (ur *UsersRepository) CreateSession(sessionId string, userId int, refreshTTL time.Duration) error {

	q := fmt.Sprintf("INSERT INTO tokens (id, user_id, issued_at, expired_at) VALUES ('%s', %d, '%s', '%s') RETURNING id",
		sessionId, userId, time.Now().Format(domain.TimeFormat), time.Now().Add(refreshTTL).Format(domain.TimeFormat))

	_, err := ur.db.Exec(q)
	return err
}

func (ur *UsersRepository) RemoveSession(sessionId string) (domain.RefreshToken, error) {

	res := domain.RefreshToken{}
	q1 := fmt.Sprintf("DELETE FROM tokens WHERE id='%s' RETURNING *", sessionId)
	if err := ur.db.QueryRow(q1).Scan(&res.Id, &res.UserId, &res.IssuedAt, &res.ExpiresAt); err != nil {
		logrus.Error(err)
		return domain.RefreshToken{}, domain.ErrNotAuthorized
	}
	fmt.Println(res)
	return res, nil
}
