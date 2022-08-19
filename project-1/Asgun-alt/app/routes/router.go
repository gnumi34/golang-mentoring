package routes

import (
	"net/http"

	_ "github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/docs"
	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/user/controllers/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type RouteControllerList struct {
	UsersController users.UserController
	JWTConfig       middleware.JWTConfig
}

func (controller RouteControllerList) RoutesUser(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	jwtMiddleware := middleware.JWTWithConfig(controller.JWTConfig)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello the program is functioning properly, welcome to the user routes.")
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/login", controller.UsersController.Login)

	e.POST("/users/get-user", controller.UsersController.GetUser)
	e.POST("/users", controller.UsersController.AddUser)
	e.PUT("/users", controller.UsersController.UpdateUser)
	e.DELETE("users/:id", controller.UsersController.DeleteUser)

	e.GET("/user/protected", controller.UsersController.Protected, jwtMiddleware)
}
