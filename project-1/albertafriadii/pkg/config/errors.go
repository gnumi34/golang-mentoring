package config

import (
	"errors"
	"net/http"
)

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFound            = errors.New("Data not found")
	ErrIDEmpty             = errors.New("ID was empty")
	ErrUsernameEmpty       = errors.New("Required Username")
	ErrEmailEmpty          = errors.New("Required Email")
	ErrPasswordEmpty       = errors.New("Required Password")
	ErrPasswordNotMatch    = errors.New("Password not match")
	ErrFieldEmpty          = errors.New("All field must be filled")
	ErrBookAlreadyCreated  = errors.New("Book Already Created")
	ErrStockBook           = errors.New("Book out of stock")
	ErrLendRequest         = errors.New("Lend requst not found")
)

func ErrGetUserCheck(err error) (int, error) {
	switch {
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound, ErrNotFound
	case errors.Is(err, ErrUsernameEmpty):
		return http.StatusBadRequest, ErrUsernameEmpty
	case errors.Is(err, ErrPasswordEmpty):
		return http.StatusBadRequest, ErrPasswordEmpty
	case errors.Is(err, ErrPasswordNotMatch):
		return http.StatusBadRequest, ErrPasswordNotMatch
	default:
		return http.StatusInternalServerError, ErrInternalServerError
	}
}

func ErrCreateUserCheck(err error) (int, error) {
	switch {
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound, ErrNotFound
	case errors.Is(err, ErrUsernameEmpty):
		return http.StatusBadRequest, ErrUsernameEmpty
	case errors.Is(err, ErrEmailEmpty):
		return http.StatusBadRequest, ErrEmailEmpty
	case errors.Is(err, ErrPasswordEmpty):
		return http.StatusBadRequest, ErrPasswordEmpty
	default:
		return http.StatusInternalServerError, ErrInternalServerError
	}
}

func ErrUpdateUserCheck(err error) (int, error) {
	switch {
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound, ErrNotFound
	case errors.Is(err, ErrIDEmpty):
		return http.StatusBadRequest, ErrIDEmpty
	case errors.Is(err, ErrUsernameEmpty):
		return http.StatusBadRequest, ErrUsernameEmpty
	case errors.Is(err, ErrInternalServerError):
		return http.StatusBadRequest, ErrEmailEmpty
	case errors.Is(err, ErrPasswordEmpty):
		return http.StatusBadRequest, ErrPasswordEmpty
	default:
		return http.StatusInternalServerError, ErrInternalServerError
	}
}

func ErrDeleteCheck(err error) (int, error) {
	if errors.Is(err, ErrNotFound) {
		return http.StatusNotFound, ErrNotFound
	}
	return http.StatusInternalServerError, ErrInternalServerError
}

func ErrInputBookCheck(err error) (int, error) {
	if errors.Is(err, ErrFieldEmpty) {
		return http.StatusBadRequest, ErrFieldEmpty
	} else if errors.Is(err, ErrBookAlreadyCreated) {
		return http.StatusBadRequest, ErrBookAlreadyCreated
	}
	return http.StatusInternalServerError, ErrInternalServerError
}

func ErrBorrowBookCheck(err error) (int, error) {
	switch {
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound, ErrNotFound
	case errors.Is(err, ErrIDEmpty):
		return http.StatusNotFound, ErrIDEmpty
	case errors.Is(err, ErrStockBook):
		return http.StatusNotAcceptable, ErrStockBook
	case errors.Is(err, ErrLendRequest):
		return http.StatusNotFound, ErrLendRequest
	default:
		return http.StatusInternalServerError, ErrInternalServerError
	}
}
