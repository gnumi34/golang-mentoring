package http

import (
	"errors"
	"golang-mentoring/project-1/Asgun-alt/pkg/common/controller"
	"golang-mentoring/project-1/Asgun-alt/pkg/domain/auth"
	"golang-mentoring/project-1/Asgun-alt/pkg/helper/errcode"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	handler controller.BaseController
	UseCase auth.UseCase
}

func NewAuthController(appGroup *echo.Group, uc auth.UseCase) {
	handler := &AuthHandler{
		UseCase: uc,
	}

	authGroup := appGroup.Group("/auth")
	authGroup.POST("/login", handler.Login)
}

// Login godoc
// @Summary      Login
// @Description  If user is exists in the database, Generate and RETURN user token.
// @Param 		 User body auth.LoginUserRequest true "Login"
// @Tags         Login
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /api/auth/login [post]
func (h *AuthHandler) Login(c echo.Context) error {
	var req auth.LoginUserRequest
	var res *auth.Response

	err := c.Bind(&req)
	if err != nil {
		return h.handler.ErrorResponse(c, http.StatusBadRequest, errors.New("bind data error"))
	}

	if req.Username == "" {
		return errcode.ErrUsernameEmpty
	}
	if req.Password == "" {
		return errcode.ErrPasswordEmpty
	}

	ctx := c.Request().Context()
	res, err = h.UseCase.Login(ctx, &req)
	if err != nil {
		errCode, errMessage := errcode.ErrorUnauthorizedCheck(err)
		return h.handler.ErrorResponse(c, errCode, errMessage)
	}
	return h.handler.SuccessDataResponse(c, res)
}
