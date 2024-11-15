package hash

import (
	"crypto/sha1"
	"fmt"
)

type PasswordHasher struct {
	passwordSalt []byte
}

func NewHasher(salt string) *PasswordHasher {
	return &PasswordHasher{
		passwordSalt: []byte(salt),
	}
}

func (h *PasswordHasher) GetHash(password string) (string, error) {
	hash := sha1.New()

	if _, err := hash.Write([]byte(password)); err != nil {
		return "", fmt.Errorf("Error occured while getting hash: %v", err)
	}
	return fmt.Sprintf("%x", hash.Sum(h.passwordSalt)), nil
}

//
