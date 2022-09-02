package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type UsersDomain struct {
	ID        uint
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type UsersUsecaseInterface interface {
	GetUser(ctx context.Context, req *UsersDomain) (*UsersDomain, error)
	AddUser(ctx context.Context, req *UsersDomain) (*UsersDomain, error)
	UpdateUser(ctx context.Context, oldPassword string, req *UsersDomain) (*UsersDomain, error)
	DeleteUser(ctx context.Context, id uint) error
}

type UsersRepositoryInterface interface {
	GetUser(ctx context.Context, req *UsersDomain) (*UsersDomain, error)
	AddUser(ctx context.Context, req *UsersDomain) (*UsersDomain, error)
	UpdateUser(ctx context.Context, req *UsersDomain) (*UsersDomain, error)
	DeleteUser(ctx context.Context, id uint) error
	FindUserByID(ctx context.Context, id uint) (*UsersDomain, error)
}
