package errcode

import (
	"errors"
	"net/http"
)

func CheckErrorUserUsecase(thisError error) (int, error) {
	switch {
	case errors.Is(thisError, ErrIDEmpty):
		return http.StatusBadRequest, ErrIDEmpty
	case errors.Is(thisError, ErrNotFound):
		return http.StatusNotFound, ErrNotFound
	case errors.Is(thisError, ErrUsernameEmpty):
		return http.StatusBadRequest, ErrUsernameEmpty
	case errors.Is(thisError, ErrPasswordEmpty):
		return http.StatusBadRequest, ErrPasswordEmpty
	case errors.Is(thisError, ErrEmailEmpty):
		return http.StatusBadRequest, ErrEmailEmpty
	case errors.Is(thisError, ErrUserNotFound):
		return http.StatusNotFound, ErrUserNotFound
	case errors.Is(thisError, ErrWrongPassword):
		return http.StatusForbidden, ErrWrongPassword
	default:
		return http.StatusInternalServerError, ErrInternalServer
	}
}

func CheckErrorUnauthorized(thisError error) (int, error) {
	switch {
	case errors.Is(thisError, ErrUnauthorized):
		return http.StatusUnauthorized, ErrUnauthorized
	case errors.Is(thisError, ErrUsernameEmpty):
		return http.StatusBadRequest, ErrUsernameEmpty
	case errors.Is(thisError, ErrPasswordEmpty):
		return http.StatusBadRequest, ErrPasswordEmpty
	case errors.Is(thisError, ErrWrongPassword):
		return http.StatusBadRequest, ErrWrongPassword
	case errors.Is(thisError, ErrRecordNotFound):
		return http.StatusNotFound, ErrRecordNotFound
	case errors.Is(thisError, ErrTokenClaims):
		return http.StatusBadRequest, ErrTokenClaims
	default:
		return http.StatusInternalServerError, ErrInternalServer
	}
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
