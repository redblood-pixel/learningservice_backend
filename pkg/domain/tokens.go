package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
	ID int
	jwt.RegisteredClaims
}

type RefreshToken struct {
	ID        string
	UserID    int
	IssuedAt  time.Time
	ExpiresAt time.Time
}
