package response

import (
	"golang-mentoring/project-1/Asgun-alt/pkg/domain/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Username  string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type GetUsersResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromUserDomain(userDomain *users.UsersDomain) *User {
	return &User{
		ID:        userDomain.ID,
		Username:  userDomain.Username,
		Email:     userDomain.Email,
		Password:  userDomain.Password,
		CreatedAt: userDomain.CreatedAt,
		UpdatedAt: userDomain.UpdatedAt,
		DeletedAt: userDomain.DeletedAt,
	}
}
