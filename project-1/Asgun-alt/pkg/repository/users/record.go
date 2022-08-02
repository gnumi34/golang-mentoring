package dbusers

import (
	"time"

	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/service/uservice"
	"gorm.io/gorm"
)

type Users struct {
	ID         string `gorm:"primaryKey"`
	Username   string
	Email      string `gorm:"unique"`
	Password   string
	Created_At time.Time      `gorm:"autoCreateTime"`
	Updated_At time.Time      `gorm:"autoUpdateTime"`
	Deleted_At gorm.DeletedAt `gorm:"index"`
}

func (user Users) ToUserDomain() uservice.UsersDomain {
	return uservice.UsersDomain{
		ID:         user.ID,
		Username:   user.Username,
		Email:      user.Email,
		Password:   user.Password,
		Created_At: user.Created_At,
		Updated_At: user.Updated_At,
	}
}

func FromUserDomain(domain uservice.UsersDomain) Users {
	return Users{
		ID:         domain.ID,
		Username:   domain.Username,
		Email:      domain.Email,
		Password:   domain.Password,
		Created_At: domain.Created_At,
		Updated_At: domain.Updated_At,
	}
}
