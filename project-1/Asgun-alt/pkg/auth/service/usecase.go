package service

import (
	"context"
	common "golang-mentoring/project-1/Asgun-alt/pkg/common/auth"
	"golang-mentoring/project-1/Asgun-alt/pkg/domain/auth"
	"golang-mentoring/project-1/Asgun-alt/pkg/helper/encrypt"
	"golang-mentoring/project-1/Asgun-alt/pkg/helper/errcode"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type AuthUseCase struct {
	DBRepo auth.Repository
}

func NewAuthUseCase(dbRepo auth.Repository) *AuthUseCase {
	return &AuthUseCase{
		DBRepo: dbRepo,
	}
}

func (uc *AuthUseCase) Login(ctx context.Context, req *auth.LoginUserRequest) (*auth.Response, error) {
	user, err := uc.DBRepo.Login(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	match := encrypt.CheckPassword(req.Password, user.Password)
	if !match {
		return nil, errcode.ErrWrongPassword
	}

	expiredAt := time.Now().Add(time.Hour * 1)
	claims := &common.JWTCustomClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt.Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(viper.GetString("jwt.secretKey")))
	if err != nil {
		return nil, errcode.ErrTokenClaims
	}

	return &auth.Response{
		Token:     token,
		ExpiredAt: expiredAt,
	}, nil
}
