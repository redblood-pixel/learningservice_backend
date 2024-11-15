package gorm_repo

import (
	"fmt"

	"github.com/redblood-pixel/learning-service-go/pkg/config"
	"github.com/redblood-pixel/learning-service-go/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg *config.PostgresDB) (*domain.DatabaseORM, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &domain.DatabaseORM{DB: db}, nil
}
