package usecase

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"golang-mentoring/project-1/albertafriadii/pkg/config"
	"golang-mentoring/project-1/albertafriadii/pkg/domain"

	"github.com/joho/godotenv"

	"github.com/golang-jwt/jwt"
)

type userUsecase struct {
	userRepo domain.UserRepositoryInterface
}

func NewUserUsecase(u domain.UserRepositoryInterface) domain.UserUsecaseInterface {
	return &userUsecase{
		userRepo: u,
	}
}

func (uc *userUsecase) FindAll(ctx context.Context) ([]domain.Users, error) {
	users, err := uc.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc *userUsecase) LoginUser(ctx context.Context, req *domain.GetUser) (*domain.Result, error) {

	user, err := uc.userRepo.LoginUser(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	isValid := config.CheckPassword(req.Password, user.Password)
	if !isValid {
		return nil, config.ErrPasswordNotMatch
	}

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	// generate token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	return &domain.Result{
		Token: t,
	}, nil
}

func (uc *userUsecase) CreateUser(ctx context.Context, u *domain.Users) (*domain.Users, error) {

	var err error

	u.Password, err = config.HashPassword(u.Password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	users, err := uc.userRepo.CreateUser(ctx, u)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return users, nil
}

func (uc *userUsecase) UpdateUser(ctx context.Context, existPassword string, u *domain.Users) error {
	user, err := uc.userRepo.GetUser(ctx, u.UserId)
	if err != nil {
		return err
	}

	if existPassword != "" {
		if !config.CheckPassword(existPassword, user.Password) {
			return config.ErrPasswordNotMatch
		}

		u.Password, err = config.HashPassword(u.Password)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	err = uc.userRepo.UpdateUser(ctx, u)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (uc *userUsecase) DeleteUser(ctx context.Context, UserID uint) error {
	err := uc.userRepo.DeleteUser(ctx, UserID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
