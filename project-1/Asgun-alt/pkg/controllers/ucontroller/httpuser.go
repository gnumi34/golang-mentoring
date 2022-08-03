package ucontroller

import (
	"errors"
	"net/http"

	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/controllers"
	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/controllers/ucontroller/request"
	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/helper/err"
	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/helper/validate"
	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/service/uservice"
	"github.com/labstack/echo/v4"
)

var (
	badRequest = http.StatusBadRequest
)

type UserController struct {
	usecase uservice.UsersUsecaseInterface
}

func NewUserController(userUsecase uservice.UsersUsecaseInterface) *UserController {
	return &UserController{usecase: userUsecase}
}

// Validate whether the user exists within the database
func (controller *UserController) GetUser(c echo.Context) error {
	req := request.GetUser{}
	bindErr := c.Bind(&req)
	if bindErr != nil {
		return controllers.ErrorResponse(c, badRequest, errors.New("data bind error"))
	}
	ctx := c.Request().Context()

	userData, result := controller.usecase.GetUser(ctx, req.ToGetUserDomain())
	if result != nil {
		errCode, errMessage := err.ErrorGetUserCheck(result)
		return controllers.ErrorResponse(c, errCode, errMessage)
	}

	if userData.Username == "" && userData.Password == "" {
		return controllers.SuccessMessageResponse(c, "invalid user")
	}
	return controllers.SuccessMessageResponse(c, "valid user")
}

func (controller *UserController) AddUsers(c echo.Context) error {
	req := request.AddUsers{}
	bindErr := c.Bind(&req)
	if bindErr != nil {
		return controllers.ErrorResponse(c, badRequest, errors.New("data bind error"))
	}
	ctx := c.Request().Context()

	if validate.MustNotBeEmpty(req.Username) || validate.MustNotBeEmpty(req.Email) {
		return controllers.ErrorResponse(c, badRequest, errors.New("username or email cannot be empty"))
	}
	if !validate.ValidateEmail(req.Email) {
		return controllers.ErrorResponse(c, badRequest, errors.New("invalid email"))
	}
	if req.Password_1 != req.Password_2 {
		return controllers.ErrorResponse(c, badRequest, errors.New("password didn't match"))
	}
	if !validate.ValidatePassword(req.Password_1) {
		return controllers.ErrorResponse(c, badRequest, errors.New("password must have 1 uppercase letter, 1 number, 1 special character"))
	}

	_, result := controller.usecase.AddUsers(ctx, req.ToUserDomain())
	if result != nil {
		errCode, errMessage := err.CheckErrorAddUsers(result)
		return controllers.ErrorResponse(c, errCode, errMessage)
	}
	return controllers.SuccessOkResponse(c)
}

func (controller *UserController) UpdateUsers(c echo.Context) error {
	req := request.UpdateUsers{}
	bindErr := c.Bind(&req)
	if bindErr != nil {
		return controllers.ErrorResponse(c, badRequest, errors.New("data bind error"))
	}

	ctx := c.Request().Context()
	_, result := controller.usecase.UpdateUsers(ctx, req.ToUpdateUserDomain())

	if result != nil {
		errCode, errMessage := err.ErrorUpdateUsersCheck(result)
		return controllers.ErrorResponse(c, errCode, errMessage)
	}
	return controllers.SuccessOkResponse(c)
}

func (controller *UserController) DeleteUsers(c echo.Context) error {
	id := c.Param("id")

	ctx := c.Request().Context()
	result := controller.usecase.DeleteUsers(ctx, id)
	if result != nil {
		errCode, errMessage := err.ErrorDeleteUsersCheck(result)
		return controllers.ErrorResponse(c, errCode, errMessage)
	}

	return controllers.SuccessOkResponse(c)
}
