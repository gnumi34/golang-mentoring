package main

import (
	"fmt"
	"golang-mentoring/project-1/albertafriadii/handler"

	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

func main() {
	h := handler.Handler{}
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	// Database connection
	h.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	// Routes
	e.POST("/create", h.CreateUser)

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}
