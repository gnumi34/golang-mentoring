package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/app/routes"
	usercontroller "github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/controllers/ucontroller"
	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/repository"
	userRepo "github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/repository/users"
	userUseCase "github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/service/uservice"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func main() {
	DBconfig := &repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	var db *gorm.DB = DBconfig.InitDB()

	repository.DbMigrate(db)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	e := echo.New()

	usersRepoInterface := userRepo.NewDBUserRepository(db)
	usersUseCaseInterface := userUseCase.NewUserUsecCase(usersRepoInterface, timeoutContext)
	usersUseControllerInterface := usercontroller.NewUserController(usersUseCaseInterface)

	initRoutes := routes.RouteControllerList{
		UsersController: *usersUseControllerInterface,
	}

	initRoutes.RoutesUser(e)
	fmt.Println(e.Start(":8000"))
}
