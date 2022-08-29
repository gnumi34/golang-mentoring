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

// func ExtractJWT(ctx echo.Context) (payload *JWTCustomClaims, err error) {
// 	header := ctx.Request().Header
// 	bearerToken := header.Get("Authorization")

// 	// extract bearer token if exist
// 	tokenPayload := &JWTCustomClaims{}
// 	if len(bearerToken) > 0 {
// 		token := strings.Split(bearerToken, " ")[1]
// 		token = strings.Split(token, ".")[1]
// 		tokenByte, err := jwt.DecodeSegment(token)
// 		if err != nil {
// 			return nil, err
// 		}
// 		if err := json.Unmarshal(tokenByte, tokenPayload); err != nil {
// 			return nil, err
// 		}
// 	}
// 	return tokenPayload, nil
// }

// func ValidateAuthorization(ctx echo.Context) (payload *JWTCustomClaims, err error) {
// 	tokenPayload, err := ExtractJWT(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if tokenPayload.username == "" {
// 		return nil, errcode.ErrUnauthorized
// 	}

// 	return tokenPayload, nil
// }
