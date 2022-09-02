package auth

import "github.com/golang-jwt/jwt"

type JWTCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
