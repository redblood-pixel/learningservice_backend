package hash

import (
	"crypto/sha1"
	"fmt"
)

type Config struct {
	PasswordSalt string `mapstructure:"password_salt"`
}

type PasswordHasher struct {
	passwordSalt []byte
}

func NewHasher(cfg *Config) *PasswordHasher {
	return &PasswordHasher{
		passwordSalt: []byte(cfg.PasswordSalt),
	}
}

func (h *PasswordHasher) GetHash(password string) (string, error) {
	hash := sha1.New()

	if _, err := hash.Write([]byte(password)); err != nil {
		return "", fmt.Errorf("error occured while getting hash: %v", err)
	}
	return fmt.Sprintf("%x", hash.Sum(h.passwordSalt)), nil
}

//
