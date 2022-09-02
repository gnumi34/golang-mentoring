package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type InputCreateUser struct {
	UserId     string `json:"user_id"`
	Username   string `json:"username" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password_1" validate:"required,min=8"`
	Repassword string `json:"password_2" validate:"required,min=8"`
}
type InputUpdateUser struct {
	UserId           string `json:"user_id"`
	Username         string `json:"username" validate:"required"`
	Email            string `json:"email"`
	ExistingPassword string `json:"existing_password"`
	Password         string `json:"password_1"`
	Repassword       string `json:"password_2"`
}

type GetUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Result struct {
	Token string `json:"token"`
}

type Users struct {
	UserId    string         `gorm:"column:user_id" json:"user_id"`
	Username  string         `gorm:"column:username" json:"username" validate:"required"`
	Email     string         `gorm:"column:email" json:"email" validate:"required,email"`
	Password  string         `gorm:"column:password" json:"password" validate:"required,min=8"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type UserUsecaseInterface interface {
	LoginUser(ctx context.Context, req *GetUser) (*Result, error)
	CreateUser(ctx context.Context, u *Users) (*Users, error)
	UpdateUser(ctx context.Context, existPassword string, u *Users, UserID string) error
	DeleteUser(ctx context.Context, UserID string) error
}

type UserRepositoryInterface interface {
	GetUser(ctx context.Context, UserID string) (*Users, error)
	LoginUser(ctx context.Context, Username string) (*Users, error)
	CreateUser(ctx context.Context, u *Users) (*Users, error)
	UpdateUser(ctx context.Context, u *Users, UserID string) error
	DeleteUser(ctx context.Context, UserID string) error
}

func FromUserDomain(u *Users) *Users {
	return &Users{
		UserId:    u.UserId,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *GetUser) ToGetUserDomain() *Users {
	return &Users{
		Username: u.Username,
		Password: u.Password,
	}
}

func (u *InputCreateUser) ToCreateUserDomain() *Users {
	return &Users{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}

func (u *InputUpdateUser) ToUpdateUserDomain() *Users {
	return &Users{
		UserId:   u.UserId,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}
