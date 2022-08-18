package http

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gnumi34/golang-mentoring/project-1/gnumi34/cmd/config"
	"github.com/gnumi34/golang-mentoring/project-1/gnumi34/pkg/domain/common"
	"github.com/gnumi34/golang-mentoring/project-1/gnumi34/pkg/domain/users"
	"github.com/gnumi34/golang-mentoring/project-1/gnumi34/pkg/helpers"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UsersHTTPHandler struct {
	common.BaseHTTPHandler
	UseCase users.UseCase
}

func NewUsersHTTPHandler(appGroup *echo.Group, uc users.UseCase) {
	handler := &UsersHTTPHandler{
		UseCase: uc,
	}

	usersGroup := appGroup.Group("/users")
	usersGroup.GET("", handler.FindAll)
	usersGroup.POST("/login", handler.ValidateUser)
	usersGroup.POST("", handler.AddUser)
	usersGroup.PUT("/:id", handler.UpdateUser)
	usersGroup.DELETE("/:id", handler.DeleteUser)
	return
}

// FindAll godoc
// @Summary      Find All Users
// @Description  Find All Users saved in the DB
// @Tags         users
// @Produce      json
// @Success      200  {object}  common.Response{message=common.DataSuccess,data=[]users.User,code=200}
// @Failure      400  {object}  common.Response{message=common.ValidationError,data=[]users.User,code=400}
// @Failure      404  {object}  common.Response{message=common.DataFailed,errors=common.RecordNotFound,code=404}
// @Failure      500  {object}  common.Response{message=common.DataFailed,errors=common.UnknownError,code=500}
// @Router       /users [get]
func (h *UsersHTTPHandler) FindAll(ctx echo.Context) error {
	res, err := h.UseCase.FindAll(ctx.Request().Context())
	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return h.ResponseJSON(ctx, common.DataFailed, nil, common.RecordNotFound, http.StatusNotFound)
		}

		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusInternalServerError)
	}

	return h.ResponseJSON(ctx, common.DataSuccess, users.ToMultipleResponse(res), nil, http.StatusOK)
}

// ValidateUser godoc
// @Summary      Validate User
// @Description  Validate User Account
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        jsonBody   body      users.ValidateUserRequest  true  "Validate User Request Body"
// @Success      200  {object}  common.Response{message=common.DataSuccess,data=valid user,code=200}
// @Failure      400  {object}  common.Response{message=common.ValidationError,code=400}
// @Failure      404  {object}  common.Response{message=common.DataFailed,errors=common.RecordNotFound,code=404}
// @Failure      500  {object}  common.Response{message=common.DataFailed,errors=common.UnknownError,code=500}
// @Router       /login [post]
func (h *UsersHTTPHandler) ValidateUser(ctx echo.Context) error {
	var request users.ValidateUserRequest
	var isValid bool
	var err error
	valid := ctx.Get("validator").(*config.CustomValidator)

	err = ctx.Bind(&request)
	if err != nil {
		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusUnprocessableEntity)
	}

	err = ctx.Validate(&request)
	if err != nil {
		if valErr, ok := err.(validator.ValidationErrors); ok {
			return h.ResponseJSON(ctx, common.ValidationError, nil, valErr.Translate(valid.Translator), http.StatusBadRequest)
		}

		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, err.Error(), http.StatusNotFound)
	}

	isValid, err = h.UseCase.ValidateUser(ctx.Request().Context(), &request)
	if err != nil {
		log.Println(err.Error())

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return h.ResponseJSON(ctx, common.DataFailed, nil, common.RecordNotFound, http.StatusNotFound)
		}

		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusInternalServerError)
	}

	if !isValid {
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.PasswordNotMatch, http.StatusNotFound)
	}

	return h.ResponseJSON(ctx, common.DataSuccess, "valid user", nil, http.StatusOK)
}

// AddUser godoc
// @Summary      Update User by ID
// @Description  Update user data by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        jsonBody   body      users.AddUserRequest  true  "Add User Request Body"
// @Success      201  {object}  common.Response{message=common.DataSuccess,data=users.UserResponse,code=201}
// @Failure      400  {object}  common.Response{message=common.ValidationError,code=400}
// @Failure      404  {object}  common.Response{message=common.DataFailed,errors=common.RecordNotFound,code=404}
// @Failure      500  {object}  common.Response{message=common.DataFailed,errors=common.UnknownError,code=500}
// @Router       /users [post]
func (h *UsersHTTPHandler) AddUser(ctx echo.Context) error {
	var request users.AddUserRequest
	var user *users.User
	var err error
	valid := ctx.Get("validator").(*config.CustomValidator)

	err = ctx.Bind(&request)
	if err != nil {
		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusUnprocessableEntity)
	}

	err = ctx.Validate(&request)
	if err != nil {
		if valErr, ok := err.(validator.ValidationErrors); ok {
			return h.ResponseJSON(ctx, common.ValidationError, nil, valErr.Translate(valid.Translator), http.StatusBadRequest)
		}

		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusNotFound)
	}

	isValidPassword := helpers.IsPasswordOK(request.Password1)
	if !isValidPassword {
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.InvalidPassword, http.StatusBadRequest)
	}

	user, err = h.UseCase.CreateUser(ctx.Request().Context(), request.ToUserDomain())
	if err != nil {
		if errors.Is(err, common.ErrUserAlreadyCreated) {
			return h.ResponseJSON(ctx, common.DataFailed, nil, common.UserAlreadyCreated, http.StatusBadRequest)
		}
		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusInternalServerError)
	}

	return h.ResponseJSON(ctx, common.DataSuccess, user.ToResponse(), nil, http.StatusCreated)
}

// UpdateUser godoc
// @Summary      Update User by ID
// @Description  Update user data by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Param        jsonBody   body      users.UpdateUserRequest  true  "Update User Request Body"
// @Success      200  {object}  common.Response{message=common.DataSuccess,data=users.UserResponse,code=200}
// @Failure      400  {object}  common.Response{message=common.ValidationError,code=400}
// @Failure      404  {object}  common.Response{message=common.DataFailed,errors=common.RecordNotFound,code=404}
// @Failure      500  {object}  common.Response{message=common.DataFailed,errors=common.UnknownError,code=500}
// @Router       /users/{id} [put]
func (h *UsersHTTPHandler) UpdateUser(ctx echo.Context) error {
	var request users.UpdateUserRequest
	var id int
	var err error

	idStr := ctx.Param("id")
	id, err = strconv.Atoi(idStr)
	if err != nil {
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.InvalidUserID, http.StatusBadRequest)
	}

	standAloneValidator := ctx.Get("validator").(*config.CustomValidator)
	err = ctx.Bind(&request)
	if err != nil {
		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusUnprocessableEntity)
	}

	err = ctx.Validate(&request)
	if err != nil {
		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusBadRequest)
	}

	if uint(id) != request.ID {
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.InvalidUserID, http.StatusBadRequest)
	}

	if request.ExistingPassword != "" {
		if request.Password1 == "" {
			return h.ResponseJSON(ctx, common.ValidationError, nil, common.PasswordNotFilled, http.StatusBadRequest)
		}

		if request.Password1 != request.Password2 {
			return h.ResponseJSON(ctx, common.ValidationError, nil, common.PasswordNotSame, http.StatusBadRequest)
		}

		isValidPassword := helpers.IsPasswordOK(request.Password1)
		if !isValidPassword {
			return h.ResponseJSON(ctx, common.ValidationError, nil, common.InvalidPassword, http.StatusBadRequest)
		}
	} else {
		// To prevent updating the password when the existing password is not filled
		request.Password1, request.Password2 = "", ""
	}

	if request.UserName != "" {
		err = (standAloneValidator).Validator.Var(&request.UserName, "max=50")
		if err != nil {
			valErr := err.(validator.ValidationErrors)
			return h.ResponseJSON(ctx, common.ValidationError, nil, valErr.Translate(standAloneValidator.Translator), http.StatusBadRequest)
		}
	}

	if request.Email != "" {
		err = (standAloneValidator).Validator.Var(&request.Email, "email")
		if err != nil {
			valErr := err.(validator.ValidationErrors)
			return h.ResponseJSON(ctx, common.ValidationError, nil, valErr.Translate(standAloneValidator.Translator), http.StatusBadRequest)
		}
	}

	err = h.UseCase.UpdateUser(ctx.Request().Context(), request.ExistingPassword, request.ToUserDomain())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return h.ResponseJSON(ctx, common.DataFailed, nil, common.RecordNotFound, http.StatusNotFound)
		}

		if errors.Is(err, common.ErrPasswordNotMatch) {
			return h.ResponseJSON(ctx, common.DataFailed, nil, common.PasswordNotMatch, http.StatusBadRequest)
		}

		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusInternalServerError)
	}

	return h.ResponseJSON(ctx, common.DataSuccess, nil, nil, http.StatusOK)
}

// DeleteUser godoc
// @Summary      Delete User by ID
// @Description  Delete User by ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  common.Response{message=common.DataSuccess,code=200}
// @Failure      400  {object}  common.Response{message=common.ValidationError,code=400}
// @Failure      404  {object}  common.Response{message=common.DataFailed,errors=common.RecordNotFound,code=404}
// @Failure      500  {object}  common.Response{message=common.DataFailed,errors=common.UnknownError,code=500}
// @Router       /users/{id} [delete]
func (h *UsersHTTPHandler) DeleteUser(ctx echo.Context) error {
	var id int
	var err error

	idStr := ctx.Param("id")
	id, err = strconv.Atoi(idStr)
	if err != nil {
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.InvalidUserID, http.StatusBadRequest)
	}

	err = h.UseCase.DeleteUser(ctx.Request().Context(), &users.User{
		Model: gorm.Model{
			ID: uint(id),
		},
	})
	if err != nil {
		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusInternalServerError)
	}

	return h.ResponseJSON(ctx, common.DataSuccess, nil, nil, http.StatusOK)
}
