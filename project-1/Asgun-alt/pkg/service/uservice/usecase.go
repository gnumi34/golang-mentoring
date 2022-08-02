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

func NewUserUsecCase(userRepo UsersRepositoryInterface, contextTimeout time.Duration) UsersRepositoryInterface {
	return &UserUseCase{repo: userRepo, ctx: contextTimeout}
}

func (usecase UserUseCase) AddUsers(ctx context.Context, userDomain UsersDomain) (UsersDomain, error) {
	if userDomain.ID == "" {
		userDomain.ID = uuid.NewV4().String()
	}

	users, result := usecase.repo.AddUsers(ctx, userDomain)
	if result != nil {
		return UsersDomain{}, result
	}

	return users, nil
}

func (usecase UserUseCase) UpdateUsers(ctx context.Context, updateDomain UsersDomain) (UsersDomain, error) {

	if updateDomain.ID == "" {
		return UsersDomain{}, err.ErrIDEmpty
	}
	if updateDomain.Username == "" {
		return UsersDomain{}, err.ErrUsernameEmpty
	}
	if updateDomain.Email == "" {
		return UsersDomain{}, err.ErrEmailEmpty
	}
	if updateDomain.Password == "" {
		return UsersDomain{}, err.ErrPasswordEmpty
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
