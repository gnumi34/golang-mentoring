package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4/middleware"
)

func NewJWTMiddlewareConfig() middleware.JWTConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	return middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
	}
}
