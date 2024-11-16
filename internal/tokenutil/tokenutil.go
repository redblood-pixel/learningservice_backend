package tokenutil

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redblood-pixel/learning-service-go/pkg/domain"
)

type Config struct {
	AccessTokenTTL  time.Duration `mapstructure:"access_token_expiry_time"`
	RefreshTokenTTL time.Duration `mapstructure:"refresh_token_expiry_time"`
	SigningKey      string        `mapstructure:"signing_key"`
}

type TokenManager struct {
	signingKey        []byte
	accessExpireTime  time.Duration
	refreshExpireTime time.Duration
}

func NewTokenManager(cfg *Config) *TokenManager {
	return &TokenManager{
		accessExpireTime:  cfg.AccessTokenTTL,
		refreshExpireTime: cfg.RefreshTokenTTL,
		signingKey:        []byte(cfg.SigningKey),
	}
}

func (tm *TokenManager) CreateAccessToken(userId int) (string, error) {
	claims := domain.JWTCustomClaims{
		Id: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tm.accessExpireTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	t, err := token.SignedString(tm.signingKey)
	if err != nil {
		return "", err
	}

	return t, nil
}

func (tm *TokenManager) ParseAccessToken(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.JWTCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected singing method")
		}
		return tm.signingKey, nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*domain.JWTCustomClaims); ok && token.Valid {
		return claims.Id, nil
	}
	return 0, fmt.Errorf("access token not valid")
}

func (tm *TokenManager) CreateRefreshToken() string {
	return uuid.New().String()
}

func (tm *TokenManager) RefreshTTL() time.Duration {
	return tm.refreshExpireTime
}
