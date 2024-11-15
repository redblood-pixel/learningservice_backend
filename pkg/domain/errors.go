package domain

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user with such username or email is already exists")
	ErrNotAuthorized     = errors.New("user is not authorized")
)
