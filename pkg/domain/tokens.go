package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
	Id int
	jwt.RegisteredClaims
}

type RefreshToken struct {
	Id        string
	UserId    int
	IssuedAt  time.Time
	ExpiresAt time.Time
}
