package handlers

import (
	"encoding/json"
	"fmt"
	"golang-mentoring/project-1/albertafriadii/model"
	"golang-mentoring/project-1/albertafriadii/repository"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

var userRepository repository.UserRepository

func init() {
	userRepository = repository.NewUserRepository()
}

func CreateUser(c echo.Context) error {

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		errorResp := model.Response{
			Message: "Wrong Details",
			Error:   err,
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, errorResp)
	}

	u := &model.User{UserId: uuid.New().String()}
	err = json.Unmarshal(body, &u)
	if err != nil {
		errorResp := model.Response{
			Message: "Wrong Details",
			Error:   err,
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, errorResp)
	}

	num, _ := regexp.Compile(`[0-9]+?`)
	resNum := num.MatchString(u.Password)
	sym, _ := regexp.Compile(`[^a-zA-Z0-9]+?`)
	resSym := sym.MatchString(u.Password)
	upper, _ := regexp.Compile(`[A-Z]+?`)
	resUpper := upper.MatchString(u.Password)

	if u.Username == "" || u.Password == "" || u.RePassword == "" || u.Email == "" {
		errorResp := model.Response{
			Message: "Required Username, Password_1 and Password_2",
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

	if len(u.Password) < 8 || resSym == false || resNum == false || resUpper == false {
		errorResp := model.Response{
			Message: "Must be at least 8 letters, alphanumeric + symbol, has at least 1 uppercase letter, has at least 1 number, and has at least 1 symbol.",
			Error:   err,
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, errorResp)
	}

	userCreated, err := userRepository.SaveAUser(*u)
	if err != nil {
		errorResp := model.Response{
			Message: "Failed to create data user",
			Error:   err,
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, errorResp)
	}
	resp := model.Response{
		Message: "Successfully to create data user",
		Error:   err,
		Data:    nil,
	}
	c.Response().Header().Set("Location", fmt.Sprintf("%s%s%s", c.Request().Host, c.Request().RequestURI, userCreated))
	return c.JSON(http.StatusCreated, resp)
}

func UpdateUser(c echo.Context) error {
	u := &model.User{}

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		errorResp := model.Response{
			Message: "Not Acceptable",
			Error:   err,
			Data:    nil,
		}
		return c.JSON(http.StatusNotAcceptable, errorResp)
	}

	err = json.Unmarshal(body, &u)
	if err != nil {
		errorResp := model.Response{
			Message: "Not Acceptable",
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

	if u.Username == "" || u.Password == "" || u.RePassword == "" || u.Email == "" {
		errorResp := model.Response{
			Message: "Required Username, Password_1 and Password_2",
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

	if len(u.Password) < 8 || resSym == false || resNum == false || resUpper == false {
		errorResp := model.Response{
			Message: "Must be at least 8 letters, alphanumeric + symbol, has at least 1 uppercase letter, has at least 1 number, and has at least 1 symbol.",
			Error:   err,
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, errorResp)
	}

	updateUser, err := userRepository.UpdateAUser(*u, u.UserId)
	if err != nil {
		errorResp := model.Response{
			Message: "Error in updating user information",
			Error:   err,
			Data:    nil,
		}
		return c.JSON(http.StatusBadGateway, errorResp)
	}
	resp := model.Response{
		Message: "Error in updating user information",
		Error:   err,
		Data:    updateUser,
	}
	return c.JSON(http.StatusOK, resp)
}

func DeleteUser(c echo.Context) error {
	u := &model.User{}

	deleteUser, err := userRepository.GetUser(c.Param(u.UserId))
	if userRepository.DeleteAUser(*u, u.UserId) != nil {
		errorResp := model.Response{
			Message: "Error in updating user information",
			Error:   err,
			Data:    nil,
		}
		return c.JSON(http.StatusBadGateway, errorResp)
	}

	resp := model.Response{
		Message: "Error in updating user information",
		Error:   err,
		Data:    deleteUser,
	}
	return c.JSON(http.StatusOK, resp)
}
