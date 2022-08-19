package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/app/middlewares"
	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/app/routes"
	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/helper"
	userController "github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/user/controllers/users"
	userRepo "github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/user/repository/users"
	userUseCase "github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/user/service/users"
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

func init() {
	viper.SetConfigFile(`config.json`)
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}
}

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

	// Configure middleware with the custom claims type
	configJWT := middlewares.ConfigJWT{
		SecretKey:       viper.GetString(`jwt.secretkey`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
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
	usersUseCaseInterface := userUseCase.NewUserUseCase(usersRepoInterface, &configJWT)
	usersUseControllerInterface := userController.NewUserController(usersUseCaseInterface)

	initRoutes := routes.RouteControllerList{
		UsersController: *usersUseControllerInterface,
		JWTConfig:       configJWT.Init(),
	}

	initRoutes.RoutesUser(e)
	fmt.Println(e.Start(viper.GetString(`server.address`)))
}
