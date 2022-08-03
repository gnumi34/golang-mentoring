package request

import "github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/service/uservice"

type GetUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AddUsers struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password_1 string `json:"password_1"`
	Password_2 string `json:"password_2"`
}

type UpdateUsers struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password_1 string `json:"password_1"`
	Password_2 string `json:"password_2"`
}

func (user *GetUser) ToGetUserDomain() uservice.UsersDomain {
	return uservice.UsersDomain{
		Username: user.Username,
		Password: user.Password,
	}
}

func (user *AddUsers) ToUserDomain() uservice.UsersDomain {
	return uservice.UsersDomain{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password_1,
	}
}

func (user *UpdateUsers) ToUpdateUserDomain() uservice.UsersDomain {
	return uservice.UsersDomain{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password_1,
	}
}
