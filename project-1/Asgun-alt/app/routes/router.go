package routes

import (
	"net/http"

	_ "golang-mentoring/project-1/Asgun-alt/docs"
	customMiddleware "golang-mentoring/project-1/Asgun-alt/pkg/middleware"
	users "golang-mentoring/project-1/Asgun-alt/pkg/users/controllers/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type RouteControllerList struct {
	UsersController users.UserController
}

func (controller RouteControllerList) RoutesUser(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())

	JWTConfig := customMiddleware.NewJWTMiddlewareConfig()
	jwtMiddleware := middleware.JWTWithConfig(JWTConfig)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello the program is functioning properly, welcome to the user routes.")
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/users", controller.UsersController.AddUser)
	e.POST("/users/get-user", controller.UsersController.GetUser, jwtMiddleware)
	e.PUT("/users", controller.UsersController.UpdateUser, jwtMiddleware)
	e.DELETE("users/:id", controller.UsersController.DeleteUser, jwtMiddleware)

	e.GET("/user/protected", controller.UsersController.Protected, jwtMiddleware)
}
