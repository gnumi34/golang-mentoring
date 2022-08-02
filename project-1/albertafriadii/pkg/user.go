package pkg

import (
	"encoding/json"
	"fmt"
	"golang-mentoring/project-1/albertafriadii/model"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type Response struct {
	Message string      `json:"message"`
	Error   error       `json:"error"`
	Data    interface{} `json:"data"`
}

func (h *Handler) CreateUser(c echo.Context) error {

	response := new(Response)

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		response.Message = "Wrong Details"
		response.Error = err
		response.Data = nil
		return err
	}

	u := model.User{UserId: uuid.New()}
	err = json.Unmarshal(body, &u)
	if err != nil {
		response.Message = "Wrong Details"
		response.Error = err
		response.Data = nil
		return err
	}

	err = u.Validate("")
	if err != nil {
		response.Message = "Wrong Details"
		response.Error = err
		response.Data = nil
		return err
	}

	userCreated, err := u.SaveAUser(h.DB)
	if err != nil {
		response.Message = "Failed to create data user"
		response.Error = err
		response.Data = nil
		return err
	} else {
		response.Message = "Succesfully to create data user"
		response.Error = nil
		response.Data = userCreated
	}
	c.Response().Header().Set("Location", fmt.Sprintf("%s%s%d", c.Request().Host, c.Request().RequestURI, userCreated.UserId.ID()))

	return c.JSON(http.StatusCreated, response)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	response := new(Response)
	u := model.User{}

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		response.Message = "Wrong Details"
		response.Error = err
		response.Data = nil
		return err
	}

	err = json.Unmarshal(body, &u)
	if err != nil {
		response.Message = "Wrong Details"
		response.Error = err
		response.Data = nil
		return err
	}

	err = u.Validate("update")
	if err != nil {
		response.Message = "Wrong Details"
		response.Error = err
		response.Data = nil
		return err
	}

	users, err := u.UpdateAUser(h.DB, c.Param(u.UserId.String()))
	if err != nil {
		response.Message = "Failed"
		response.Error = err
		response.Data = nil
		return err
	} else {
		response.Message = "Successfully"
		response.Error = nil
		response.Data = users
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	response := new(Response)
	u := new(model.User)

	users, err := u.DeleteAUser(h.DB, c.Param(u.UserId.String()))
	if err != nil {
		response.Message = "Failed"
		response.Error = err
		response.Data = nil
		return err
	} else {
		response.Message = "Successfully"
		response.Error = nil
		response.Data = users
	}

	return c.JSON(http.StatusOK, response)
}
