package errcode

import "errors"

var (
	// general
	ErrInternalServer     = errors.New("something gone wrong, contact administrator")
	ErrIDEmpty            = errors.New("id is empty")
	ErrNotFound           = errors.New("data not found")
	ErrServiceUnavailable = errors.New("server is temporarily unable to handle a request")
	ErrUnauthorized       = errors.New("error unauthorized")
	ErrRecordNotFound     = errors.New("record not found")
	ErrTokenClaims        = errors.New("error token claims")

	// users
	ErrUserNotFound  = errors.New("user not found")
	ErrEmailEmpty    = errors.New("email cannot be empty")
	ErrUsernameEmpty = errors.New("username cannot be empty")
	ErrPasswordEmpty = errors.New("password cannot be empty")
	ErrNoUser        = errors.New("there is no such user")
	ErrWrongPassword = errors.New("wrong password")

	// books
	ErrBookNotFound        = errors.New("book not found")
	ErrStockUnavailable    = errors.New("book is out of stock")
	ErrLendRequestNotFound = errors.New("lend request not found")
	ErrMaxStockBookLimit   = errors.New("stock exceeded max book limit")
)
