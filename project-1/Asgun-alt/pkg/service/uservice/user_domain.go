package uservice

import (
	"context"
	"time"
)

type UsersDomain struct {
	ID         string
	Username   string
	Email      string
	Password   string
	Created_At time.Time
	Updated_At time.Time
	Deleted_At time.Time
}

type UsersUsecaseInterface interface {
	GetUser(ctx context.Context, userDomain UsersDomain) (UsersDomain, error)
	AddUsers(ctx context.Context, userDomain UsersDomain) (UsersDomain, error)
	UpdateUsers(ctx context.Context, userUpdateDomain UsersDomain) (UsersDomain, error)
	DeleteUsers(CTX context.Context, id string) error
}

type UsersRepositoryInterface interface {
	GetUser(ctx context.Context, userDomain UsersDomain) (UsersDomain, error)
	AddUsers(ctx context.Context, userDomain UsersDomain) (UsersDomain, error)
	UpdateUsers(ctx context.Context, userUpdateDomain UsersDomain) (UsersDomain, error)
	DeleteUsers(CTX context.Context, id string) error
}
