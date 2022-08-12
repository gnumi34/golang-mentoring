package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseController struct{}

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

func (c *BaseController) SuccessDataResponse(ctx echo.Context, data interface{}) error {
	response := DataResponse{}
	response.Data = data
	return ctx.JSON(http.StatusOK, response)
}

func (c *BaseController) SuccessOkResponse(ctx echo.Context) error {
	response := OkResponse{}
	response.Message = "OK"
	return ctx.JSON(http.StatusOK, response)
}

func (c *BaseController) SuccessMessageResponse(ctx echo.Context, message string) error {
	response := OkResponse{}
	response.Message = message
	return ctx.JSON(http.StatusOK, response)
}

func (c *BaseController) ErrorResponse(ctx echo.Context, status int, errMessage error) error {
	response := ErrResponse{}
	response.Error = errMessage.Error()
	return ctx.JSON(status, response)
}

func (c *BaseController) ErrorValidationResponse(ctx echo.Context, status int, errMessage interface{}) error {
	response := ErrResponse{}
	response.Error = errMessage
	return ctx.JSON(status, response)
}
