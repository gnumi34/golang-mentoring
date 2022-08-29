package uservice

import (
	"context"
	"time"

	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/helper/err"
	uuid "github.com/satori/go.uuid"
)

type UserUseCase struct {
	repo UsersRepositoryInterface
	ctx  time.Duration
}

func NewUserUseCase(userRepo UsersRepositoryInterface, contextTimeout time.Duration) UsersUsecaseInterface {
	return &UserUseCase{repo: userRepo, ctx: contextTimeout}
}

func (usecase *UserUseCase) GetUser(ctx context.Context, userDomain UsersDomain) (UsersDomain, error) {
	if userDomain.Username == "" {
		return UsersDomain{}, err.ErrUsernameEmpty
	}
	if userDomain.Password == "" {
		return UsersDomain{}, err.ErrPasswordEmpty
	}

	user, err := usecase.repo.GetUser(ctx, userDomain)
	if err != nil {
		return UsersDomain{}, err
	}
	return user, nil
}

func (usecase *UserUseCase) AddUsers(ctx context.Context, userDomain UsersDomain) (UsersDomain, error) {
	if userDomain.ID == "" {
		userDomain.ID = uuid.NewV4().String()
	}

	users, result := usecase.repo.AddUsers(ctx, userDomain)
	if result != nil {
		return UsersDomain{}, result
	}

	return users, nil
}

func (usecase *UserUseCase) UpdateUsers(ctx context.Context, updateDomain UsersDomain) (UsersDomain, error) {

	if updateDomain.ID == "" {
		return UsersDomain{}, err.ErrIDEmpty
	}
	if updateDomain.Username == "" {
		return UsersDomain{}, err.ErrUsernameEmpty
	}

	users, result := usecase.repo.UpdateUsers(ctx, updateDomain)
	if result != nil {
		return UsersDomain{}, result
	}
	if (users.Created_At == time.Time{} && users.Email == "") {
		return UsersDomain{}, err.ErrNotFound
	}
	return users, nil
}

func (usecase *UserUseCase) DeleteUsers(ctx context.Context, id string) error {
	result := usecase.repo.DeleteUsers(ctx, id)
	if result != nil {
		return result
	}
	return nil
}
