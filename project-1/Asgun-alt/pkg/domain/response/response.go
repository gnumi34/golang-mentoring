package response

import (
	"time"

	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/service/uservice"
)

type UserResponse struct {
	ID         string
	Username   string
	Email      string
	Password   string
	Created_At time.Time
	Updated_At time.Time
	Deleted_At time.Time
}

type GetUsersResponse struct {
	ID         string    `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

func FromUserDomain(userDomain uservice.UsersDomain) UserResponse {
	return UserResponse{
		ID:         userDomain.ID,
		Username:   userDomain.Username,
		Email:      userDomain.Email,
		Created_At: userDomain.Created_At,
		Updated_At: userDomain.Updated_At,
		Deleted_At: userDomain.Deleted_At,
	}
}

func FromGetUserDomain(userDomain uservice.UsersDomain) GetUsersResponse {
	return GetUsersResponse{
		ID:         userDomain.ID,
		Username:   userDomain.Username,
		Email:      userDomain.Email,
		Created_At: userDomain.Created_At,
		Updated_At: userDomain.Updated_At,
	}
}
