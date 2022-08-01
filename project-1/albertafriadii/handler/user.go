package handler

import (
	"golang-mentoring/project-1/albertafriadii/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func (h *Handler) CreateUser(c echo.Context) (err error) {
	// Bind
	u := &model.User{UserId: uuid.New()}
	if err = c.Bind(u); err != nil {
		return
	}

	// Validate
	if u.Username == "" || u.Password == "" || u.RePassword == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid email or password"}
	}

	if u.Password != u.RePassword {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Password not match"}
	}

	if len(u.Password) < 8 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Must be at least 8 letters, alphanumeric + symbol, has at least 1 uppercase letter, has at least 1 number, and has at least 1 symbol."}
	}

	// Save user
	err = h.DB.Debug().Create(&u).Error
	if err != nil {
		return
	}
	return c.JSON(http.StatusCreated, u)
}
