package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

type ConfigJWT struct {
	SecretKey      string
	ExpireDuration int
}

type JWTCustomClaims struct {
	username string
	jwt.StandardClaims
}

func (c *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JWTCustomClaims{},
		SigningKey: []byte(c.SecretKey),
	}
}

func (c *ConfigJWT) GenerateToken(username string) (string, error) {
	claims := JWTCustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(c.ExpireDuration))).Unix(),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(c.SecretKey))
	return token, err
}
