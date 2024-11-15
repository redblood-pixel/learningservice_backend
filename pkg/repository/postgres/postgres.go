package postgres_repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redblood-pixel/learning-service-go/pkg/config"
	"github.com/redblood-pixel/learning-service-go/pkg/domain"
)

func NewPostgresDB(cfg *config.Cfg) (*domain.Database, error) {

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.PostgresDB.DBHost, cfg.PostgresDB.DBPort, cfg.PostgresDB.DBUser,
		cfg.PostgresDB.DBName, cfg.PostgresDB.DBPassword, cfg.PostgresDB.DBSSLMode))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return nil, nil
}
