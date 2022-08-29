package main

import (
	"fmt"
	"log"
	"os"

	"github.com/albertafriadii/tree/featured/albert-jwt-auth/pkg/config"
	"github.com/albertafriadii/tree/featured/albert-jwt-auth/pkg/users/delivery/http"
	"github.com/albertafriadii/tree/featured/albert-jwt-auth/pkg/users/repository"
	"github.com/albertafriadii/tree/featured/albert-jwt-auth/pkg/users/usecase"
	"github.com/albertafriadii/tree/featured/albert-jwt-auth/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	DBConfig := &config.Config{
		DB_Host:     os.Getenv("DB_HOST"),
		DB_User:     os.Getenv("DB_USER"),
		DB_Password: os.Getenv("DB_PASSWORD"),
		DB_Name:     os.Getenv("DB_NAME"),
		DB_Port:     os.Getenv("DB_PORT"),
	}
	var db *gorm.DB = DBConfig.IntializeDB()
	e := echo.New()

	userRepoInterface := repository.NewUserRepositroy(db)
	userUsecaseInterface := usecase.NewUserUsecase(userRepoInterface)
	userUseControllerInterface := http.NewUserController(userUsecaseInterface)

	initRoutes := routes.RouteList{
		UsersController: *userUseControllerInterface,
	}

	initRoutes.RouteUsers(e)
	fmt.Println(e.Start(":1323"))
}
