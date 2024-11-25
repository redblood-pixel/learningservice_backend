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

func (ur *UsersRepository) Create(username, email, password string) (userID int, err error) {

	q := fmt.Sprintf("INSERT INTO %s (username, email, password_hash) VALUES ('%s', '%s', '%s') RETURNING id",
		UserTable, username, email, password)

	row := ur.db.QueryRow(q)
	if err = row.Scan(&userID); err != nil {
		return 0, err
	}

	return
}

func (ur *UsersRepository) FindUserByEmail(email, password string) (userID int, err error) {

	q := fmt.Sprintf("SELECT id FROM %s WHERE email='%s' AND password_hash='%s'", UserTable, email, password)

	if err = ur.db.Get(&userID, q); err != nil {
		return 0, err
	}
	return userID, nil
}

func (ur *UsersRepository) CreateSession(sessionId string, userID int, refreshTTL time.Duration) error {

	q := fmt.Sprintf("INSERT INTO %s (id, user_id, issued_at, expired_at) VALUES ('%s', %d, '%s', '%s') RETURNING id",
		TokenTable, sessionId, userID, time.Now().Format(domain.TimeFormat), time.Now().Add(refreshTTL).Format(domain.TimeFormat))

	_, err := ur.db.Exec(q)
	return err
}

func (ur *UsersRepository) RemoveSession(sessionId string) (domain.RefreshToken, error) {

	res := domain.RefreshToken{}
	q := fmt.Sprintf("DELETE FROM %s WHERE id='%s' RETURNING *", TokenTable, sessionId)
	if err := ur.db.QueryRow(q).Scan(&res.ID, &res.UserID, &res.IssuedAt, &res.ExpiresAt); err != nil {
		logrus.Error(err)
		return domain.RefreshToken{}, domain.ErrNotAuthorized
	}
	fmt.Println(res)
	return res, nil
}
