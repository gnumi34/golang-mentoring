package users

import (
	"context"

	"golang-mentoring/project-1/Asgun-alt/pkg/domain/users"
	"golang-mentoring/project-1/Asgun-alt/pkg/helper/encrypt"
	"golang-mentoring/project-1/Asgun-alt/pkg/helper/errcode"
)

type UserUseCase struct {
	repo users.UsersRepositoryInterface
}

func NewUserUseCase(userRepo users.UsersRepositoryInterface) *UserUseCase {
	return &UserUseCase{
		repo: userRepo,
	}
}

func (usecase *UserUseCase) GetUser(ctx context.Context, userDomain *users.UsersDomain) (*users.UsersDomain, error) {
	user, err := usecase.repo.GetUser(ctx, userDomain)
	if err != nil {
		return nil, errcode.ErrUserNotFound
	}

	match := encrypt.CheckPassword(userDomain.Password, user.Password)
	if !match {
		return nil, errcode.ErrWrongPassword
	}

	return user, nil
}

func (usecase *UserUseCase) AddUser(ctx context.Context, userDomain *users.UsersDomain) (*users.UsersDomain, error) {
	var err error

	userDomain.Password, err = encrypt.HashPassword(userDomain.Password)
	if err != nil {
		return nil, err
	}

	users, err := usecase.repo.AddUser(ctx, userDomain)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (usecase *UserUseCase) UpdateUser(ctx context.Context, oldPassword string, req *users.UsersDomain) (*users.UsersDomain, error) {
	user, err := usecase.repo.FindUserByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	if req.Password != "" {
		match := encrypt.CheckPassword(oldPassword, user.Password)
		if !match {
			return nil, errcode.ErrWrongPassword
		}

		req.Password, err = encrypt.HashPassword(req.Password)
		if err != nil {
			return nil, err
		}
	}

	users, err := usecase.repo.UpdateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (usecase *UserUseCase) DeleteUser(ctx context.Context, id uint) error {
	err := usecase.repo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
