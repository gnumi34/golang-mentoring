package http

import (
	"errors"
	"log"
	"net/http"

	"github.com/gnumi34/golang-mentoring/project-1/gnumi34/cmd/config"
	"github.com/gnumi34/golang-mentoring/project-1/gnumi34/pkg/domain/auth"
	"github.com/gnumi34/golang-mentoring/project-1/gnumi34/pkg/domain/common"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthHTTPHandler struct {
	common.BaseHTTPHandler
	UseCase auth.UseCase
}

func NewAuthHTTPHandler(appGroup *echo.Group, uc auth.UseCase) {
	handler := &AuthHTTPHandler{
		UseCase: uc,
	}

	usersGroup := appGroup.Group("/auth")
	usersGroup.POST("/login", handler.ValidateUser)
	return
}

// ValidateUser godoc
// @Summary      Validate User
// @Description  Validate User Account
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        jsonBody   body      auth.ValidateUserRequest  true  "Validate User Request Body"
// @Success      200  {object}  common.Response{message=common.DataSuccess,data=valid user,code=200}
// @Failure      400  {object}  common.Response{message=common.ValidationError,code=400}
// @Failure      404  {object}  common.Response{message=common.DataFailed,errors=common.RecordNotFound,code=404}
// @Failure      500  {object}  common.Response{message=common.DataFailed,errors=common.UnknownError,code=500}
// @Router       /login [post]
func (h *AuthHTTPHandler) ValidateUser(ctx echo.Context) error {
	var request auth.ValidateUserRequest
	var resp *auth.Response
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

	resp, err = h.UseCase.ValidateUser(ctx.Request().Context(), &request)
	if err != nil {
		if errors.Is(err, common.ErrPasswordNotMatch) {
			return h.ResponseJSON(ctx, common.DataFailed, nil, common.PasswordNotMatch, http.StatusBadRequest)
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return h.ResponseJSON(ctx, common.DataFailed, nil, common.RecordNotFound, http.StatusNotFound)
		}

		log.Println(err.Error())
		return h.ResponseJSON(ctx, common.DataFailed, nil, common.UnknownError, http.StatusInternalServerError)
	}

	return h.ResponseJSON(ctx, common.DataSuccess, resp, nil, http.StatusOK)
}
