package usecase

import (
	"context"

	"github.com/gnumi34/golang-mentoring/project-1/gnumi34/pkg/domain/common"
	"github.com/gnumi34/golang-mentoring/project-1/gnumi34/pkg/domain/users"
	"github.com/gnumi34/golang-mentoring/project-1/gnumi34/pkg/helpers"
)

type UsersUseCase struct {
	DBRepo users.DBRepository
}

func NewUsersUseCase(dbRepo users.DBRepository) *UsersUseCase {
	return &UsersUseCase{
		DBRepo: dbRepo,
	}
}

func (uc *UsersUseCase) FindAll(ctx context.Context) ([]users.User, error) {
	users, err := uc.DBRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uc *UsersUseCase) ValidateUser(ctx context.Context, req *users.ValidateUserRequest) (bool, error) {
	user, err := uc.DBRepo.FindByUserName(ctx, req.UserName)
	if err != nil {
		return false, err
	}

	return helpers.CheckPasswordHash(req.Password, user.Password), nil
}

func (uc *UsersUseCase) CreateUser(ctx context.Context, req *users.User) (*users.User, error) {
	var err error

	req.Password, err = helpers.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	req, err = uc.DBRepo.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (uc *UsersUseCase) UpdateUser(ctx context.Context, existingPassword string, req *users.User) error {
	user, err := uc.DBRepo.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}

	if existingPassword != "" {
		if !helpers.CheckPasswordHash(existingPassword, user.Password) {
			return common.ErrPasswordNotMatch
		}

		req.Password, err = helpers.HashPassword(req.Password)
		if err != nil {
			return err
		}
	}

	err = uc.DBRepo.UpdateByID(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UsersUseCase) DeleteUser(ctx context.Context, req *users.User) error {
	err := uc.DBRepo.DeleteByID(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
