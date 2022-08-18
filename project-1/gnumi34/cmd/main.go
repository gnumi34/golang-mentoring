package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gnumi34/golang-mentoring/project-1/gnumi34/cmd/config"
	"github.com/gnumi34/golang-mentoring/project-1/gnumi34/pkg/domain/users"
	customMiddleware "github.com/gnumi34/golang-mentoring/project-1/gnumi34/pkg/middleware"
	usersHTTPHandler "github.com/gnumi34/golang-mentoring/project-1/gnumi34/pkg/users/handler/http"
	usersDBRepository "github.com/gnumi34/golang-mentoring/project-1/gnumi34/pkg/users/repository/db"
	usersUseCase "github.com/gnumi34/golang-mentoring/project-1/gnumi34/pkg/users/usecase"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// @title Library Management API
// @version 0.0.1-alpha
// @description This is an API application to manage the library.
// @termsOfService http://example.com

// @contact.name Library API Support
// @contact.url http://example.com/support
// @contact.email support@example.com

// @host localhost:8081
// @BasePath /api

// @produce json
func main() {
	var err error
	var db *gorm.DB

	e := echo.New()
	e.Debug = true
	err = config.InitConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	postgresConfig := config.NewDBConfig(
		viper.GetString("database.provider"),
		viper.GetString("database.db_name"),
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.hostname"),
		viper.GetString("database.port"),
		viper.GetString("database.timezone"),
	)
	db, err = postgresConfig.InitDB()
	if err != nil {
		panic(fmt.Errorf("fatal error init DB: %w", err))
	}

	e.Use(customMiddleware.RecoverMiddleware())
	validator := config.NewCustomValidator()
	e.Validator = validator
	e.Use(func(handle echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Set("validator", validator)
			return handle(ctx)
		}
	})
	e.Use(customMiddleware.CORSMiddleware)

	// Init Handlers
	e.GET("", func(ctx echo.Context) error {
		return ctx.JSON(200, echo.Map{
			"message": "hello world",
		})
	})
	api := e.Group("/api")
	{
		InitUserHandler(api, db)
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", viper.GetString("app.hostname"), viper.GetString("app.port")),
		Handler:      e,
		WriteTimeout: 3 * time.Minute,
		ReadTimeout:  3 * time.Minute,
		IdleTimeout:  5 * time.Minute,
	}
	go func() {
		if err := e.StartServer(srv); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(fmt.Sprintf("Error start server %v\n", err))
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		panic(fmt.Sprintf("Server forced to shutdown %v\n", err))
	}
	if <-ctx.Done(); true {
		fmt.Println("timeout of 5 seconds.")
	}
}

func InitUserHandler(appGroup *echo.Group, db *gorm.DB, middlewares ...echo.MiddlewareFunc) {
	var dbRepository users.DBRepository = usersDBRepository.NewUsersDBRepository(db)
	var useCase users.UseCase = usersUseCase.NewUsersUseCase(dbRepository)

	if len(middlewares) > 0 {
		appGroup.Use(middlewares...)
	}
	usersHTTPHandler.NewUsersHTTPHandler(appGroup, useCase)
}
