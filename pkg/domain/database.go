package domain

import (
	"github.com/jmoiron/sqlx"
)

type Database struct {
	*sqlx.DB
}
