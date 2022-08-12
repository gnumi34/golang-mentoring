package users

import (
	"context"

	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/helper/errcode"
	uuid "github.com/satori/go.uuid"
)

type UserUseCase struct {
	repo UsersRepositoryInterface
}

func NewUserUseCase(userRepo UsersRepositoryInterface) UsersUsecaseInterface {
	return &UserUseCase{repo: userRepo}
}

func (usecase *UserUseCase) GetUser(ctx context.Context, userDomain UsersDomain) (UsersDomain, error) {
	if userDomain.Username == "" {
		return UsersDomain{}, errcode.ErrUsernameEmpty
	}
	if userDomain.Password == "" {
		return UsersDomain{}, errcode.ErrPasswordEmpty
	}

	user, err := usecase.repo.GetUser(ctx, userDomain)
	if err != nil {
		return UsersDomain{}, err
	}
	return user, nil
}

func (usecase *UserUseCase) AddUser(ctx context.Context, userDomain UsersDomain) (UsersDomain, error) {
	if userDomain.ID == "" {
		userDomain.ID = uuid.NewV4().String()
	}

	users, err := usecase.repo.AddUser(ctx, userDomain)
	if err != nil {
		return UsersDomain{}, err
	}
	return users, nil
}

func (usecase *UserUseCase) UpdateUser(ctx context.Context, updateDomain UsersDomain) (UsersDomain, error) {
	users, err := usecase.repo.UpdateUser(ctx, updateDomain)
	if err != nil {
		return UsersDomain{}, err
	}
	return users, nil
}

func (usecase *UserUseCase) DeleteUser(ctx context.Context, id string) error {
	err := usecase.repo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
