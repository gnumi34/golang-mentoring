package errcode

import "errors"

var (
	ErrInternalServer     = errors.New("something gone wrong, contact administrator")
	ErrIDEmpty            = errors.New("id is empty")
	ErrNotFound           = errors.New("data not found")
	ErrEmailEmpty         = errors.New("email cannot be empty")
	ErrUsernameEmpty      = errors.New("username cannot be empty")
	ErrPasswordEmpty      = errors.New("password cannot be empty")
	ErrNoUser             = errors.New("there is no such user")
	ErrServiceUnavailable = errors.New("server is temporarily unable to handle a request")
)
