package request

import (
	domain "golang-mentoring/project-1/Asgun-alt/pkg/domain/users"
)

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AddUser struct {
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password1 string `json:"password_1" validate:"required,gte=8,eqfield=Password2"`
	Password2 string `json:"password_2" validate:"required,gte=8"`
}

type UpdateUser struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	OldPassword string `json:"old_password"`
	Password1   string `json:"password_1"`
	Password2   string `json:"password_2"`
}

func (user *GetUser) ToGetUserDomain() *domain.UsersDomain {
	return &domain.UsersDomain{
		Username: user.Username,
		Password: user.Password,
	}
}

func (user *AddUser) ToUserDomain() *domain.UsersDomain {
	return &domain.UsersDomain{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password1,
	}
}

func (user *UpdateUser) ToUpdateUserDomain() *domain.UsersDomain {
	return &domain.UsersDomain{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password1,
	}
}
