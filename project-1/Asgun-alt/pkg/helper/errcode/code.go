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
	} else if errors.Is(thisError, ErrWrongPassword) {
		return http.StatusBadRequest, ErrWrongPassword
	}

	return http.StatusInternalServerError, ErrInternalServer
}

func ErrorDeleteUsersCheck(thisError error) (int, error) {
	if errors.Is(thisError, ErrNotFound) {
		return http.StatusNotFound, ErrServiceUnavailable
	} else if errors.Is(thisError, ErrUserNotFound) {
		return http.StatusNotFound, ErrUserNotFound
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
	} else if errors.Is(thisError, ErrUserNotFound) {
		return http.StatusNotFound, ErrUserNotFound
	} else if errors.Is(thisError, ErrWrongPassword) {
		return http.StatusForbidden, ErrWrongPassword
	}

	return http.StatusInternalServerError, ErrInternalServer
}

func ErrorUnauthorizedCheck(thisError error) (int, error) {
	if errors.Is(thisError, ErrUnauthorized) {
		return http.StatusUnauthorized, ErrUnauthorized
	} else if errors.Is(thisError, ErrUsernameEmpty) {
		return http.StatusBadRequest, ErrUsernameEmpty
	} else if errors.Is(thisError, ErrPasswordEmpty) {
		return http.StatusBadRequest, ErrPasswordEmpty
	} else if errors.Is(thisError, ErrWrongPassword) {
		return http.StatusBadRequest, ErrWrongPassword
	} else if errors.Is(thisError, ErrRecordNotFound) {
		return http.StatusNotFound, ErrRecordNotFound
	} else if errors.Is(thisError, ErrTokenClaims) {
		return http.StatusBadRequest, ErrTokenClaims
	}

	return http.StatusInternalServerError, ErrInternalServer
}

func CheckErrorBookUsecase(thisError error) (int, error) {
	switch {
	case errors.Is(thisError, ErrNotFound):
		return http.StatusNotFound, ErrNotFound
	case errors.Is(thisError, ErrRecordNotFound):
		return http.StatusNotFound, ErrRecordNotFound
	case errors.Is(thisError, ErrIDEmpty):
		return http.StatusNotFound, ErrIDEmpty
	case errors.Is(thisError, ErrStockUnavailable):
		return http.StatusAccepted, ErrStockUnavailable
	case errors.Is(thisError, ErrLendRequestNotFound):
		return http.StatusNotFound, ErrLendRequestNotFound
	case errors.Is(thisError, ErrBookNotFound):
		return http.StatusNotFound, ErrBookNotFound
	default:
		return http.StatusInternalServerError, ErrInternalServer
	}
}
