package dbusers

// import (
// 	"time"

// 	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/user/service/users"
// 	"gorm.io/gorm"
// )

// type Users struct {
// 	ID         string `gorm:"primaryKey"`
// 	Username   string
// 	Email      string `gorm:"unique"`
// 	Password   string
// 	Created_At time.Time      `gorm:"autoCreateTime"`
// 	Updated_At time.Time      `gorm:"autoUpdateTime"`
// 	Deleted_At gorm.DeletedAt `gorm:"index"`
// }

// func (user *Users) ToUserDomain() users.UsersDomain {
// 	return users.UsersDomain{
// 		ID:         user.ID,
// 		Username:   user.Username,
// 		Email:      user.Email,
// 		Password:   user.Password,
// 		Created_At: user.Created_At,
// 		Updated_At: user.Updated_At,
// 	}
// }

// func FromUserDomain(domain users.UsersDomain) Users {
// 	return Users{
// 		ID:         domain.ID,
// 		Username:   domain.Username,
// 		Email:      domain.Email,
// 		Password:   domain.Password,
// 		Created_At: domain.Created_At,
// 		Updated_At: domain.Updated_At,
// 	}
// }
