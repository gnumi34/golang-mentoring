package common

import "errors"

var (
	ErrUserAlreadyCreated = errors.New(UserAlreadyCreated)
	ErrRecordNotFound     = errors.New(RecordNotFound)
	ErrPasswordNotMatch   = errors.New(PasswordNotMatch)
	ErrPasswordNotSame    = errors.New(PasswordNotSame)
	ErrInvalidPassword    = errors.New(InvalidPassword)
	ErrInvalidUserID      = errors.New(InvalidUserID)
)
