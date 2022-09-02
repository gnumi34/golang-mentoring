package middleware

import (
	common "golang-mentoring/project-1/Asgun-alt/pkg/common/auth"

	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func NewJWTMiddlewareConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &common.JWTCustomClaims{},
		SigningKey: []byte(viper.GetString("jwt.secretKey")),
	}
}
