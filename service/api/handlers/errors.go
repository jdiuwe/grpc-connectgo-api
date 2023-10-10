package handlers

import "errors"

var (
	ErrNilServerConfig     = errors.New("nil server config")
	ErrNotImplemented      = errors.New("not implemented")
	ErrUserExists          = errors.New("user with specified email already exists")
	ErrFailedToHashPass    = errors.New("failed to hash password")
	ErrUserNotFound        = errors.New("user not found")
	ErrBadPassword         = errors.New("incorrect password")
	ErrVerifyEmailNotFound = errors.New("verification code not found")
)
