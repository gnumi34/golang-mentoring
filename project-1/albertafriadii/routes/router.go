package routes

import (
	"net/http"

	_ "github.com/albertafriadii/tree/fix/albertafriadii/docs"
	controller "github.com/albertafriadii/tree/fix/albertafriadii/pkg/users/delivery/http"
	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Go Echo Library Management
// @version 1.0
// @description a simple Go library management with echo framework
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:1323
// @BasePath /
// @Schemes http

type RouteList struct {
	UsersController controller.UserController
}

func (r RouteList) RouteUsers(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/user/get", r.UsersController.GetUser)
	e.POST("/user", r.UsersController.CreateUser)
	e.PUT("/user/:user_id", r.UsersController.UpdateUser)
	e.DELETE("/user/:user_id", r.UsersController.DeleteUser)
}
