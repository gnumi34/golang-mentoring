package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type DataResponse struct {
	Data interface{} `json:"data"`
}

type ErrResponse struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

type OkResponse struct {
	Message interface{} `json:"message"`
}

func SuccessDataResponse(c echo.Context, data interface{}) error {
	response := DataResponse{}
	response.Data = data
	return c.JSON(http.StatusOK, response)
}

func SuccessOkResponse(c echo.Context) error {
	response := OkResponse{}
	response.Message = "OK"
	return c.JSON(http.StatusOK, response)
}

func ErrorResponse(c echo.Context, status int, errMessage error) error {
	response := ErrResponse{}
	response.Error = errMessage.Error()
	return c.JSON(http.StatusInternalServerError, response)
}
