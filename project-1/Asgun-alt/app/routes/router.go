package routes

import (
	"net/http"

	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/controllers/usersctrl"
	_ "github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Go Echo Library Management
// @version 1.0
// @description a simple Go library management with echo framework
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8000
// @BasePath /
// @Schemes http

type RouteControllerList struct {
	UsersController usersctrl.UserController
}

func (controller RouteControllerList) RoutesUser(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello the program is functioning properly, welcome to the user routes.")
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/users/get-user", controller.UsersController.GetUser)
	e.POST("/users", controller.UsersController.AddUsers)
	e.PUT("/users", controller.UsersController.UpdateUsers)
	e.DELETE("users/:id", controller.UsersController.DeleteUsers)
}
