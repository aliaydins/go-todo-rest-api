package user

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrPasswordIncorrect = errors.New("password incorrect")
)