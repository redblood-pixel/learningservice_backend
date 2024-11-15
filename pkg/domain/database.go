package domain

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type Database struct {
	*sqlx.DB
}

type DatabaseORM struct {
	*gorm.DB
}
