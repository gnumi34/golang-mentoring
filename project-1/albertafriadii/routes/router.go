package routes

import (
	"net/http"

	_ "github.com/albertafriadii/tree/fix/albertafriadii/docs"
	"github.com/albertafriadii/tree/fix/albertafriadii/pkg/config"
	controller "github.com/albertafriadii/tree/fix/albertafriadii/pkg/users/delivery/http"
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

// @host localhost:1323
// @BasePath /
// @Schemes http

type RouteList struct {
	UsersController controller.UserController
}

func (r RouteList) RouteUsers(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"message": "Hello World!"})
	})
	e.GET("/protected/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"message": "This is Protected Page"})
	}, config.IsAuthenticated)

	g := e.Group("/user")
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${method}], host=${host}${path}, status=${status}, latency_human=${latency_human}\n",
	}))

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	g.GET("/get", r.UsersController.GetUser)
	g.POST("/login", r.UsersController.LoginUser)
	g.POST("/create", r.UsersController.CreateUser)
	g.PUT("/update/:user_id", r.UsersController.UpdateUser, config.IsAuthenticated)
	g.DELETE("/delete/:user_id", r.UsersController.DeleteUser, config.IsAuthenticated)
}
