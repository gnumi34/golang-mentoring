package users

import (
	"errors"
	"net/http"

	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/common/controller"
	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/domain/request"
	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/helper/errcode"
	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/helper/validate"
	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/user/service/users"
	"github.com/labstack/echo/v4"
)

var (
	badRequest = http.StatusBadRequest
)

type UserController struct {
	handler controller.BaseController
	usecase users.UsersUsecaseInterface
}

func NewUserController(userUsecase users.UsersUsecaseInterface) *UserController {
	return &UserController{usecase: userUsecase}
}

// GetUser godoc
// @Summary      Show an account
// @Description  validate username and password, if user is exists in the database RETURN valid user
// @Param 		 User body request.GetUser true "validate user"
// @Tags         GetUser
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /users/get-user [post]
func (uc *UserController) GetUser(c echo.Context) error {
	req := request.GetUser{}
	bindErr := c.Bind(&req)
	if bindErr != nil {
		return uc.handler.ErrorResponse(c, badRequest, errors.New("bind data error"))
	}
	ctx := c.Request().Context()

	userData, err := uc.usecase.GetUser(ctx, req.ToGetUserDomain())
	if err != nil {
		errCode, errMessage := errcode.ErrorGetUserCheck(err)
		return uc.handler.ErrorResponse(c, errCode, errMessage)
	}

	if userData.Username == "" || userData.Password == "" {
		return uc.handler.ErrorResponse(c, http.StatusNotFound, errors.New("invalid user or user has been deleted"))
	}
	return uc.handler.SuccessMessageResponse(c, "valid user")
}

// AddUser godoc
// @Summary      Add user
// @Description  Add new user to the database, ID is generated by the API, password  is saved with BCrypt hash after passess validation.
// @Param 		 User body request.AddUser true "Add User"
// @Tags         Add new user
// @Accept       json
// @Produce      json
// @Success      200  {object}  users.UsersDomain
// @Router       /users/ [post]
func (uc *UserController) AddUser(c echo.Context) error {
	req := request.AddUser{}
	bindErr := c.Bind(&req)
	if bindErr != nil {
		return uc.handler.ErrorResponse(c, badRequest, errors.New("bind data error"))
	}
	ctx := c.Request().Context()

	validErr := validate.Validator(req)
	if validErr != nil {
		return uc.handler.ErrorValidationResponse(c, badRequest, validErr)
	}
	if !validate.ValidatePassword(req.Password_1) {
		return uc.handler.ErrorResponse(c, badRequest, errors.New("password must have 1 uppercase letter, 1 number, 1 special character"))
	}

	_, result := uc.usecase.AddUser(ctx, req.ToUserDomain())
	if result != nil {
		errCode, errMessage := errcode.CheckErrorAddUsers(result)
		return uc.handler.ErrorResponse(c, errCode, errMessage)
	}
	return uc.handler.SuccessOkResponse(c)
}

// UpdateUser godoc
// @Summary      Update User
// @Description  Update the user to the database.
// @Param 		 User body request.UpdateUser true "Update User"
// @Tags         Update User
// @Accept       json
// @Produce      json
// @Success      200  {object}  users.UsersDomain
// @Router       /users/ [put]
func (uc *UserController) UpdateUser(c echo.Context) error {
	req := request.UpdateUser{}
	bindErr := c.Bind(&req)
	if bindErr != nil {
		return uc.handler.ErrorResponse(c, badRequest, errors.New("bind data error"))
	}

	if req.Password_1 != "" && !validate.ValidatePassword(req.Password_1) {
		return uc.handler.ErrorResponse(c, badRequest, errors.New("password must have 1 uppercase letter, 1 number, 1 special character"))
	}

	ctx := c.Request().Context()
	_, result := uc.usecase.UpdateUser(ctx, req.ToUpdateUserDomain())

	if result != nil {
		errCode, errMessage := errcode.ErrorUpdateUsersCheck(result)
		return uc.handler.ErrorResponse(c, errCode, errMessage)
	}
	return uc.handler.SuccessOkResponse(c)

}

// DeleteUser godoc
// @Summary      Delete User
// @Description  Soft delete user by passing the user ID.
// @Param 		 id path string true "delete user"
// @Tags         (Soft) Delete User
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /users/{id} [delete]
func (uc *UserController) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	ctx := c.Request().Context()
	result := uc.usecase.DeleteUser(ctx, id)
	if result != nil {
		errCode, errMessage := errcode.ErrorDeleteUsersCheck(result)
		return uc.handler.ErrorResponse(c, errCode, errMessage)
	}
	return uc.handler.SuccessOkResponse(c)
}
