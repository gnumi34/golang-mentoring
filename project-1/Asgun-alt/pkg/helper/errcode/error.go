package errcode

import "errors"

var (
	ErrInternalServer     = errors.New("something gone wrong, contact administrator")
	ErrIDEmpty            = errors.New("id is empty")
	ErrNotFound           = errors.New("data not found")
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailEmpty         = errors.New("email cannot be empty")
	ErrUsernameEmpty      = errors.New("username cannot be empty")
	ErrPasswordEmpty      = errors.New("password cannot be empty")
	ErrNoUser             = errors.New("there is no such user")
	ErrServiceUnavailable = errors.New("server is temporarily unable to handle a request")
	ErrUnauthorized       = errors.New("error unauthorized")
	ErrRecordNotFound     = errors.New("record not found")
	ErrWrongPassword      = errors.New("wrong password")
	ErrTokenClaims        = errors.New("error token claims")
)
