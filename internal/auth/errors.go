package auth

import "errors"

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrInvalidAddress = errors.New("invalid address")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrUserExists     = errors.New("user exists")
)
