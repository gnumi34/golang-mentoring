package routes

import (
	"net/http"

	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/controllers/ucontroller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteControllerList struct {
	UsersController ucontroller.UserController
}

func (controller RouteControllerList) RoutesUser(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello the program is functioning properly, welcome to the user routes.")
	})
	e.GET("/users/get-user", controller.UsersController.GetUser)
	e.POST("/users", controller.UsersController.AddUsers)
	e.PUT("/users", controller.UsersController.UpdateUsers)
	e.DELETE("users/:id", controller.UsersController.DeleteUsers)
}
