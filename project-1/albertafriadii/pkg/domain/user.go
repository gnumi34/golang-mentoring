package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type InputCreateUser struct {
<<<<<<< HEAD
	UserId     string `json:"user_id"`
=======
>>>>>>> 289c4e129d7f4946ed954ea9078f420bf430068c
	Username   string `json:"username" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password_1" validate:"required,min=8"`
	Repassword string `json:"password_2" validate:"required,min=8"`
}
type InputUpdateUser struct {
	UserId     string `json:"user_id"`
	Username   string `json:"username" validate:"required"`
<<<<<<< HEAD
	Email      string `json:"email"`
=======
	Email      string `json:"email" validate:"required,email"`
>>>>>>> 289c4e129d7f4946ed954ea9078f420bf430068c
	Password   string `json:"password_1"`
	Repassword string `json:"password_2"`
}

type GetUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
	GetUser(ctx context.Context, u Users) (Users, error)
	LoginUser(ctx context.Context, u Users) (string, error)
	CreateUser(ctx context.Context, u Users) (Users, error)
<<<<<<< HEAD
	UpdateUser(ctx context.Context, u Users, UserID string) (Users, error)
	DeleteUser(ctx context.Context, UserID string) error
=======
	UpdateUser(ctx context.Context, u Users, user_id string) (Users, error)
	DeleteUser(ctx context.Context, user_id string) error
>>>>>>> 289c4e129d7f4946ed954ea9078f420bf430068c
}

type UserRepositoryInterface interface {
	GetUser(ctx context.Context, u Users) (Users, error)
	LoginUser(ctx context.Context, u Users) (bool, error)
	CreateUser(ctx context.Context, u Users) (Users, error)
<<<<<<< HEAD
	UpdateUser(ctx context.Context, u Users, UserID string) (Users, error)
	DeleteUser(ctx context.Context, UserID string) error
=======
	UpdateUser(ctx context.Context, u Users, user_id string) (Users, error)
	DeleteUser(ctx context.Context, user_id string) error
>>>>>>> 289c4e129d7f4946ed954ea9078f420bf430068c
}

func FromUserDomain(u Users) Users {
	return Users{
		UserId:    u.UserId,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *GetUser) ToGetUserDomain() Users {
	return Users{
		Username: u.Username,
		Password: u.Password,
	}
}

func (u *InputCreateUser) ToCreateUserDomain() Users {
	return Users{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}

func (u *InputUpdateUser) ToUpdateUserDomain() Users {
	return Users{
		UserId:   u.UserId,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}
