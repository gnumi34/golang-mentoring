package config

<<<<<<< HEAD
import (
	"github.com/labstack/echo/v4/middleware"
)
=======
import "github.com/labstack/echo/v4/middleware"
>>>>>>> 289c4e129d7f4946ed954ea9078f420bf430068c

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})
