package main

import (
	"fmt"
	"net/http"
	"os"

	"golang-mentoring/project-1/Asgun-alt/app/routes"
	"golang-mentoring/project-1/Asgun-alt/cmd/config"
	authHTTPHandler "golang-mentoring/project-1/Asgun-alt/pkg/auth/controller/http"
	authRepository "golang-mentoring/project-1/Asgun-alt/pkg/auth/repository/db"
	authUseCase "golang-mentoring/project-1/Asgun-alt/pkg/auth/service"
	"golang-mentoring/project-1/Asgun-alt/pkg/domain/auth"
	"golang-mentoring/project-1/Asgun-alt/pkg/helper"
	userController "golang-mentoring/project-1/Asgun-alt/pkg/users/controllers/http"
	userRepo "golang-mentoring/project-1/Asgun-alt/pkg/users/repository/db"
	userUseCase "golang-mentoring/project-1/Asgun-alt/pkg/users/service/users"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
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

func main() {
	DBconfig := &helper.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	var db *gorm.DB = DBconfig.InitDB()

	e := echo.New()

	err := config.InitConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8000"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete},
	}))

	usersRepoInterface := userRepo.NewDBUserRepository(db)
	usersUseCaseInterface := userUseCase.NewUserUseCase(usersRepoInterface)
	usersUseControllerInterface := userController.NewUserController(usersUseCaseInterface)

	initRoutes := routes.RouteControllerList{
		UsersController: *usersUseControllerInterface,
	}

	api := e.Group("/api")
	InitAuthHandler(api, db)

	initRoutes.RoutesUser(e)
	fmt.Println(e.Start(viper.GetString(`server.address`)))
}

func InitAuthHandler(appGroup *echo.Group, db *gorm.DB) {
	var dbRepository auth.Repository = authRepository.NewDBAuthRepository(db)
	var useCase auth.UseCase = authUseCase.NewAuthUseCase(dbRepository)
	authHTTPHandler.NewAuthController(appGroup, useCase)
}
