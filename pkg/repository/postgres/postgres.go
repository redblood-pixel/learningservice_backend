package postgres_repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redblood-pixel/learning-service-go/pkg/domain"
	"github.com/sirupsen/logrus"
)

type Config struct {
	DBHost     string `mapstructure:"db_host"`
	DBPort     string `mapstructure:"db_port"`
	DBName     string `mapstructure:"db_name"`
	DBUser     string `mapstructure:"db_user"`
	DBPassword string `mapstructure:"db_password"`
	DBSSLMode  string `mapstructure:"db_sslmode"`
}

func NewPostgresDB(cfg *Config) (*domain.Database, error) {

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser,
		cfg.DBName, cfg.DBPassword, cfg.DBSSLMode))

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return &domain.Database{DB: db}, nil
}
