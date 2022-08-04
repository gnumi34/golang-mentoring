package handlers

import (
	"encoding/json"
	"fmt"
	"golang-mentoring/project-1/albertafriadii/model"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func (h *Handler) CreateUser(c echo.Context) error {

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		errorResp := model.Response{
			Message: "Error getting information",
			Error:   err,
			Data:    nil,
		}
		return c.JSON(http.StatusNotAcceptable, errorResp)
	}

	u := model.User{UserId: uuid.New().String()}
	err = json.Unmarshal(body, &u)
	if err != nil {
		errorResp := model.Response{
			Message: "Error getting information",
			Error:   err,
			Data:    nil,
		}
		return c.JSON(http.StatusNotAcceptable, errorResp)
	}

	num, _ := regexp.Compile(`[0-9]+?`)
	resNum := num.MatchString(u.Password)
	sym, _ := regexp.Compile(`[^a-zA-Z0-9]+?`)
	resSym := sym.MatchString(u.Password)
	upper, _ := regexp.Compile(`[A-Z]+?`)
	resUpper := upper.MatchString(u.Password)

	if u.Username == "" {
		errorResp := model.Response{
			Message: "Required Username",
			Error:   err,
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, errorResp)
	}

	if u.Email == "" {
		errorResp := model.Response{
			Message: "Required Email",
			Error:   err,
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, errorResp)
	}
	if u.Password == "" || u.RePassword == "" {
		errorResp := model.Response{
			Message: "Required Password",
			Error:   err,
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, errorResp)
	}

	if u.Password != u.RePassword {
		errorResp := model.Response{
			Message: "Password not match",
			Error:   err,
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, errorResp)
	}

	if len(u.Password) < 8 || !resSym || !resNum || !resUpper {
		errorResp := model.Response{
			Message: "Must be at least 8 letters, alphanumeric + symbol, has at least 1 uppercase letter, has at least 1 number, and has at least 1 symbol.",
			Error:   err,
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, errorResp)
	}

	userCreated, err := u.SaveAUser(h.DB)
	if err != nil {
		errorResp := model.Response{
			Message: "Failed to create user data",
			Error:   err,
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, errorResp)
	}

	resp := model.Response{
		Message: "Succesfully to create user data",
		Error:   nil,
		Data:    userCreated,
	}
	c.Response().Header().Set("Location", fmt.Sprintf("%s%s%s", c.Request().Host, c.Request().RequestURI, userCreated.UserId))

	return c.JSON(http.StatusCreated, resp)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	u := new(model.User)
	// c.Bind(u)

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		errorResp := model.Response{
			Message: "Error getting request Information",
			Error:   err,
			Data:    nil,
		}
		return c.JSON(http.StatusNotAcceptable, errorResp)
	}

	err = json.Unmarshal(body, &u)
	if err != nil {
		errorResp := model.Response{
			Message: "Error getting information",
			Error:   err,
			Data:    nil,
		}
		return c.JSON(http.StatusNotAcceptable, errorResp)
	}

	if u.Username == "" {
		errorResp := model.Response{
			Message: "Required Username",
			Error:   err,
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, errorResp)
	}

	updateUser, err := u.UpdateAUser(h.DB, c.Param("user_id"))
	if err != nil {
		errorResp := model.Response{
			Message: "Failed to update user data",
			Error:   err,
			Data:    nil,
		}
		return c.JSON(http.StatusNotAcceptable, errorResp)
	}

	resp := model.Response{
		Message: "Succesfully to update user data",
		Error:   nil,
		Data:    updateUser,
	}
	c.Response().Header().Set("Location", fmt.Sprintf("%s%s%s", c.Request().Host, c.Request().RequestURI, updateUser.UserId))

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	u := model.User{}

	deleteUser, err := u.DeleteAUser(h.DB, c.Param("user_id"))
	if err != nil {
		errorResp := model.Response{
			Message: "Failed to delete user data",
			Error:   err,
			Data:    nil,
		}
		return c.JSON(http.StatusNotAcceptable, errorResp)
	}
	resp := model.Response{
		Message: "Succesfully to delete user data",
		Error:   nil,
		Data:    deleteUser,
	}
	c.Response().Header().Set("Location", fmt.Sprintf("%s%s%s", c.Request().Host, c.Request().RequestURI, deleteUser.UserId))

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUser(c echo.Context) error {
	u := model.User{}

	getUser, err := u.GetUser(h.DB, c.Param("user_id"))
	if err != nil {
		errorResp := model.Response{
			Message: "Failed to get user data",
			Error:   err,
			Data:    nil,
		}
		return c.JSON(http.StatusNotAcceptable, errorResp)
	}
	resp := model.Response{
		Message: "Succesfully to get user data",
		Error:   nil,
		Data: map[string]interface{}{
			"username": u.Username,
			"password": u.Password,
		},
	}
	c.Response().Header().Set("Location", fmt.Sprintf("%s%s%s", c.Request().Host, c.Request().RequestURI, getUser.UserId))

	return c.JSON(http.StatusOK, resp)
}
