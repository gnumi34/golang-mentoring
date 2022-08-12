package config

import (
	"errors"
	"net/http"
)

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFound            = errors.New("Data not found")
	ErrUserIDEmpty         = errors.New("UserID was empty")
	ErrUsernameEmpty       = errors.New("Required Username")
	ErrEmailEmpty          = errors.New("Required Email")
	ErrPasswordEmpty       = errors.New("Required Password")
)

func ErrGetUserCheck(err error) (int, error) {
	if errors.Is(err, ErrNotFound) {
		return http.StatusNotFound, ErrNotFound
	} else if errors.Is(err, ErrUsernameEmpty) {
		return http.StatusBadRequest, ErrUsernameEmpty
	} else if errors.Is(err, ErrPasswordEmpty) {
		return http.StatusBadRequest, ErrPasswordEmpty
	}
	return http.StatusInternalServerError, ErrInternalServerError
}

func ErrCreateUserCheck(err error) (int, error) {
	if errors.Is(err, ErrNotFound) {
		return http.StatusNotFound, ErrNotFound
	} else if errors.Is(err, ErrUsernameEmpty) {
		return http.StatusBadRequest, ErrUsernameEmpty
	} else if errors.Is(err, ErrInternalServerError) {
		return http.StatusBadRequest, ErrEmailEmpty
	} else if errors.Is(err, ErrPasswordEmpty) {
		return http.StatusBadRequest, ErrPasswordEmpty
	}
	return http.StatusInternalServerError, ErrInternalServerError
}

func ErrUpdateUserCheck(err error) (int, error) {
	if errors.Is(err, ErrNotFound) {
		return http.StatusNotFound, ErrNotFound
	} else if errors.Is(err, ErrUserIDEmpty) {
		return http.StatusBadRequest, ErrUserIDEmpty
	} else if errors.Is(err, ErrUsernameEmpty) {
		return http.StatusBadRequest, ErrUsernameEmpty
	} else if errors.Is(err, ErrInternalServerError) {
		return http.StatusBadRequest, ErrEmailEmpty
	} else if errors.Is(err, ErrPasswordEmpty) {
		return http.StatusBadRequest, ErrPasswordEmpty
	}
	return http.StatusInternalServerError, ErrInternalServerError
}

func ErrDeleteUserCheck(err error) (int, error) {
	if errors.Is(err, ErrNotFound) {
		return http.StatusNotFound, ErrNotFound
	}
	return http.StatusInternalServerError, ErrInternalServerError
}
