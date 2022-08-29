package domain

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string      `json:"message,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessDataResponse(c echo.Context, data interface{}) error {
	r := Response{}
	r.Data = data
	return c.JSON(http.StatusOK, r)
}

func OkResponse(c echo.Context) error {
	r := Response{}
	r.Message = "OK"
	return c.JSON(http.StatusOK, r)
}

func SuccessMessageResponse(c echo.Context, message string) error {
	r := Response{}
	r.Message = message
	return c.JSON(http.StatusOK, r)
}

func ErrResponse(c echo.Context, status int, err error) error {
	r := Response{}
	r.Error = err.Error()
	return c.JSON(http.StatusInternalServerError, r)
}

func ErrValidResponse(c echo.Context, status int, err interface{}) error {
	r := Response{}
	r.Error = err
	return c.JSON(http.StatusInternalServerError, r)
}
