package errcode

import (
	"errors"
	"net/http"
)

func CheckErrorAddUsers(thisError error) (int, error) {
	if errors.Is(thisError, ErrNotFound) {
		return http.StatusNotFound, ErrNotFound
	} else if errors.Is(thisError, ErrUsernameEmpty) {
		return http.StatusBadRequest, ErrUsernameEmpty
	} else if errors.Is(thisError, ErrEmailEmpty) {
		return http.StatusBadRequest, ErrEmailEmpty
	} else if errors.Is(thisError, ErrPasswordEmpty) {
		return http.StatusBadRequest, ErrPasswordEmpty
	}

	return http.StatusInternalServerError, ErrInternalServer
}

func ErrorUpdateUsersCheck(thisError error) (int, error) {
	if errors.Is(thisError, ErrNotFound) {
		return http.StatusNotFound, ErrNotFound
	} else if errors.Is(thisError, ErrUsernameEmpty) {
		return http.StatusBadRequest, ErrUsernameEmpty
	} else if errors.Is(thisError, ErrEmailEmpty) {
		return http.StatusBadRequest, ErrEmailEmpty
	} else if errors.Is(thisError, ErrPasswordEmpty) {
		return http.StatusBadRequest, ErrPasswordEmpty
	} else if errors.Is(thisError, ErrIDEmpty) {
		return http.StatusBadRequest, ErrIDEmpty
	}

	return http.StatusInternalServerError, ErrInternalServer
}

func ErrorDeleteUsersCheck(thisError error) (int, error) {
	if errors.Is(thisError, ErrNotFound) {
		return http.StatusNotFound, ErrServiceUnavailable
	}
	return http.StatusInternalServerError, ErrInternalServer
}

func ErrorGetUserCheck(thisError error) (int, error) {
	if errors.Is(thisError, ErrNotFound) {
		return http.StatusNotFound, ErrNotFound
	} else if errors.Is(thisError, ErrUsernameEmpty) {
		return http.StatusBadRequest, ErrUsernameEmpty
	} else if errors.Is(thisError, ErrPasswordEmpty) {
		return http.StatusBadRequest, ErrPasswordEmpty
	}
	return http.StatusInternalServerError, ErrInternalServer
}
