package middleware

import (
	"github.com/gnumi34/golang-mentoring/project-1/gnumi34/pkg/domain/common"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func NewJWTMiddlewareConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &common.JWTCustomClaims{},
		SigningKey: []byte(viper.GetString("app.app_secret")),
	}
}
