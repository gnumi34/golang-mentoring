package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type UsersDomain struct {
	ID         string
	Username   string
	Email      string
	Password   string
	Created_At time.Time
	Updated_At time.Time
	Deleted_At gorm.DeletedAt
}

type UsersUsecaseInterface interface {
	GetUser(ctx context.Context, userDomain UsersDomain) (UsersDomain, error)
	AddUser(ctx context.Context, userDomain UsersDomain) (UsersDomain, error)
	UpdateUser(ctx context.Context, userUpdateDomain UsersDomain) (UsersDomain, error)
	DeleteUser(ctx context.Context, id string) error
}

type UsersRepositoryInterface interface {
	GetUser(ctx context.Context, userDomain UsersDomain) (UsersDomain, error)
	AddUser(ctx context.Context, userDomain UsersDomain) (UsersDomain, error)
	UpdateUser(ctx context.Context, userUpdateDomain UsersDomain) (UsersDomain, error)
	DeleteUser(ctx context.Context, id string) error
}
