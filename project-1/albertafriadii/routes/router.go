package routes

import (
	"net/http"

	_ "golang-mentoring/project-1/albertafriadii/docs"

	bookController "golang-mentoring/project-1/albertafriadii/pkg/books/delivery/http"
	"golang-mentoring/project-1/albertafriadii/pkg/config"
	userController "golang-mentoring/project-1/albertafriadii/pkg/users/delivery/http"

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
	UsersController userController.UserController
	BooksContoller  bookController.BookController
}

func (r RouteList) RouteUsers(e *echo.Echo) {

	jwtConfig := config.NewJWTMiddlewareConfig()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"message": "Hello World!"})
	})
	e.GET("/protected/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"message": "This is Protected Page"})
	}, middleware.JWTWithConfig(jwtConfig))
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	g := e.Group("/user")
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${method}], host=${host}${path}, status=${status}, latency_human=${latency_human}\n",
	}))

	g.POST("/login", r.UsersController.LoginUser)
	g.POST("/create", r.UsersController.CreateUser)
	g.GET("/list", r.UsersController.FindAll)
	g.PUT("/update/:user_id", r.UsersController.UpdateUser, middleware.JWTWithConfig(jwtConfig))
	g.DELETE("/delete/:user_id", r.UsersController.DeleteUser, middleware.JWTWithConfig(jwtConfig))

	gb := e.Group("/book")
	gb.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${method}], host=${host}${path}, status=${status}, latency_human=${latency_human}\n",
	}))
	gb.GET("/list", r.BooksContoller.ListBook)
	gb.POST("/create", r.BooksContoller.CreateBook)
	gb.PUT("/update/:book_id", r.BooksContoller.UpdateBook)
	gb.DELETE("/delete/:book_id", r.BooksContoller.DeleteBook)
	gb.POST("/borrow", r.BooksContoller.BorrowBooks)
	gb.POST("/lend", r.BooksContoller.LendBooks)
	gb.PUT("/lend_approval", r.BooksContoller.LendApproval)
	gb.PUT("/return/:book_id", r.BooksContoller.ReturnBooks)
	gb.GET("/history_borrow/:user_id", r.BooksContoller.BorrowBookHistory)
	gb.GET("/history_lend/:user_id", r.BooksContoller.LendBookHistory)
	gb.GET("/history_return/:user_id", r.BooksContoller.ReturnBookHistory)

}
