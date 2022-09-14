package main

import (
	"fmt"
	"log"
	"os"

	"golang-mentoring/project-1/albertafriadii/pkg/config"
	userController "golang-mentoring/project-1/albertafriadii/pkg/users/delivery/http"
	userRepo "golang-mentoring/project-1/albertafriadii/pkg/users/repository"
	userUsecase "golang-mentoring/project-1/albertafriadii/pkg/users/usecase"
	"golang-mentoring/project-1/albertafriadii/routes"

	bookController "golang-mentoring/project-1/albertafriadii/pkg/books/delivery/http"
	bookRepo "golang-mentoring/project-1/albertafriadii/pkg/books/repository"
	bookUsecase "golang-mentoring/project-1/albertafriadii/pkg/books/usecase"

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

	userRepoInterface := userRepo.NewUserRepositroy(db)
	userUsecaseInterface := userUsecase.NewUserUsecase(userRepoInterface)
	userControllerInterface := userController.NewUserController(userUsecaseInterface)

	bookRepoInterface := bookRepo.NewBookDBRepository(db)
	bookUsecaseInterface := bookUsecase.NewBookUsecase(bookRepoInterface, userRepoInterface)
	bookControllerInterface := bookController.NewBookController(bookUsecaseInterface)

	initRoutes := routes.RouteList{
		UsersController: *userControllerInterface,
		BooksContoller:  *bookControllerInterface,
	}

	initRoutes.RouteUsers(e)
	fmt.Println(e.Start(":1323"))
}
