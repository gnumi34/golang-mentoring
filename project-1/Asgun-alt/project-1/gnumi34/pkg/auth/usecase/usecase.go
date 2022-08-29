package usecase

import (
	"context"
	"time"

	"github.com/gnumi34/golang-mentoring/project-1/gnumi34/pkg/domain/auth"
	"github.com/gnumi34/golang-mentoring/project-1/gnumi34/pkg/domain/common"
	"github.com/gnumi34/golang-mentoring/project-1/gnumi34/pkg/helpers"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type AuthUseCase struct {
	DBRepo auth.DBRepository
}

func NewAuthUseCase(dbRepo auth.DBRepository) *AuthUseCase {
	return &AuthUseCase{
		DBRepo: dbRepo,
	}
}

func (uc *AuthUseCase) ValidateUser(ctx context.Context, req *auth.ValidateUserRequest) (*auth.Response, error) {
	user, err := uc.DBRepo.FindByUserName(ctx, req.UserName)
	if err != nil {
		return nil, err
	}

	isValid := helpers.CheckPasswordHash(req.Password, user.Password)
	if !isValid {
		return nil, common.ErrPasswordNotMatch
	}

	expiredAt := time.Now().Add(time.Hour * 1)
	claims := &common.JWTCustomClaims{
		UserName: user.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(viper.GetString("app.app_secret")))
	if err != nil {
		return nil, err
	}

	return &auth.Response{
		Token:     t,
		ExpiredAt: expiredAt,
	}, nil
}
