package main

import (
	"golang-mentoring/project-1/albertafriadii/handlers"
	"golang-mentoring/project-1/albertafriadii/repository"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

func main() {
	_, err := repository.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Routes
	e.POST("/user/create", handlers.CreateUser)
	e.PUT("/user/update", handlers.UpdateUser)
	e.DELETE("/user/delete", handlers.DeleteUser)

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}
