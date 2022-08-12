package request

import (
	domain "github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/user/service/users"
)

type GetUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AddUser struct {
	Username   string `json:"username" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password_1 string `json:"password_1" validate:"required,gte=8,eqfield=Password_2"`
	Password_2 string `json:"password_2" validate:"required,gte=8"`
}

type UpdateUser struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password_1 string `json:"password_1"`
	Password_2 string `json:"password_2"`
}

func (user *GetUser) ToGetUserDomain() domain.UsersDomain {
	return domain.UsersDomain{
		Username: user.Username,
		Password: user.Password,
	}
}

func (user *AddUser) ToUserDomain() domain.UsersDomain {
	return domain.UsersDomain{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password_1,
	}
}

func (user *UpdateUser) ToUpdateUserDomain() domain.UsersDomain {
	return domain.UsersDomain{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password_1,
	}
}
