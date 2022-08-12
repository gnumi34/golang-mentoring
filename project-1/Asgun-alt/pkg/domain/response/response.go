package response

import (
	"time"

	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/user/service/users"
	"gorm.io/gorm"
)

// type UserResponse struct {
// 	ID         string
// 	Username   string
// 	Email      string
// 	Password   string
// 	Created_At time.Time
// 	Updated_At time.Time
// 	Deleted_At time.Time
// }

type User struct {
	ID         string `gorm:"primaryKey"`
	Username   string
	Email      string `gorm:"unique"`
	Password   string
	Created_At time.Time      `gorm:"autoCreateTime"`
	Updated_At time.Time      `gorm:"autoUpdateTime"`
	Deleted_At gorm.DeletedAt `gorm:"index"`
}

type GetUsersResponse struct {
	ID         string    `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

func FromUserDomain(userDomain users.UsersDomain) User {
	return User{
		ID:         userDomain.ID,
		Username:   userDomain.Username,
		Email:      userDomain.Email,
		Created_At: userDomain.Created_At,
		Updated_At: userDomain.Updated_At,
		Deleted_At: userDomain.Deleted_At,
	}
}

func FromGetUserDomain(userDomain users.UsersDomain) GetUsersResponse {
	return GetUsersResponse{
		ID:         userDomain.ID,
		Username:   userDomain.Username,
		Email:      userDomain.Email,
		Created_At: userDomain.Created_At,
		Updated_At: userDomain.Updated_At,
	}
}
