package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/albertafriadii/tree/albert/buildAPI/config"
	"github.com/albertafriadii/tree/albert/buildAPI/pkg/users/domain"
	uuid "github.com/satori/go.uuid"
)

type userUsecase struct {
	userRepo       domain.UserRepositoryInterface
	contextTimeout time.Duration
}

func NewUserUsecase(u domain.UserRepositoryInterface, t time.Duration) domain.UserUsecaseInterface {
	return &userUsecase{
		userRepo:       u,
		contextTimeout: t,
	}
}

func (uc *userUsecase) GetUser(ctx context.Context, u domain.Users) (domain.Users, error) {
	if u.Username == "" {
		return domain.Users{}, config.ErrUsernameEmpty
	}
	if u.Password == "" {
		return domain.Users{}, config.ErrPasswordEmpty
	}

	user, err := uc.userRepo.GetUser(ctx, u)
	if err != nil {
		fmt.Println(err)
		return domain.Users{}, err
	}
	return user, nil
}

func (uc *userUsecase) CreateUser(ctx context.Context, u domain.Users) (domain.Users, error) {
	if u.UserId == "" {
		u.UserId = uuid.NewV4().String()
	}

	user, err := uc.userRepo.CreateUser(ctx, u)
	if err != nil {
		fmt.Println(err)
		return domain.Users{}, err
	}
	return user, nil
}

func (uc *userUsecase) UpdateUser(ctx context.Context, u domain.Users, user_id string) (domain.Users, error) {
	if u.Username == "" {
		return domain.Users{}, config.ErrUsernameEmpty
	}

	user, err := uc.userRepo.UpdateUser(ctx, u, user_id)
	if err != nil {
		fmt.Println(err)
		return domain.Users{}, err
	}

	if (user.CreatedAt == time.Time{} && user.Email == "") {
		return domain.Users{}, config.ErrNotFound
	}
	return user, nil
}

func (uc *userUsecase) DeleteUser(ctx context.Context, user_id string) error {
	err := uc.userRepo.DeleteUser(ctx, user_id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
