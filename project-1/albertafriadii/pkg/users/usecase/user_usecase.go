package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/albertafriadii/tree/featured/albert-jwt-auth/pkg/config"
	"github.com/albertafriadii/tree/featured/albert-jwt-auth/pkg/domain"
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
)

type userUsecase struct {
	userRepo domain.UserRepositoryInterface
}

func NewUserUsecase(u domain.UserRepositoryInterface) domain.UserUsecaseInterface {
	return &userUsecase{
		userRepo: u,
	}
}

func (uc *userUsecase) GetUser(ctx context.Context, u domain.Users) (domain.Users, error) {

	user, err := uc.userRepo.GetUser(ctx, u)
	if err != nil {
		// fmt.Println(err)
		return domain.Users{}, err
	}
	return user, nil
}

func (uc *userUsecase) LoginUser(ctx context.Context, u domain.Users) (string, error) {

	// generate token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = u.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	_, err = uc.userRepo.LoginUser(ctx, u)
	if err != nil {
		return "", err
	}

	return t, nil
}

func (uc *userUsecase) CreateUser(ctx context.Context, u domain.Users) (domain.Users, error) {

	var err error

	if u.UserId == "" {
		u.UserId = uuid.NewV4().String()
	}

	u.Password, err = config.HashPassword(u.Password)
	if err != nil {
		fmt.Println(err)
		return domain.Users{}, nil
	}

	u, err = uc.userRepo.CreateUser(ctx, u)
	if err != nil {
		fmt.Println(err)
		return domain.Users{}, err
	}
	return u, nil
}

func (uc *userUsecase) UpdateUser(ctx context.Context, u domain.Users, UserID string) (domain.Users, error) {

	user, err := uc.userRepo.UpdateUser(ctx, u, UserID)
	if err != nil {
		fmt.Println(err)
		return domain.Users{}, err
	}

	return user, nil
}

func (uc *userUsecase) DeleteUser(ctx context.Context, UserID string) error {
	err := uc.userRepo.DeleteUser(ctx, UserID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
