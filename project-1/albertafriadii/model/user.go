package model

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserId     uuid.UUID `gorm:"column:user_id" json:"user_id,omitempty"`
	Username   string    `gorm:"column:username" json:"username,omitempty"`
	Email      string    `gorm:"column:email" json:"email"`
	Password   string    `gorm:"column:password" json:"password_1,omitempty"`
	RePassword string    `json:"password_2,omitempty"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
	// DeletedAt  time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (u *User) BeforeSave(db *gorm.DB) error {
	hash, err := HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Username == "" {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Required username"}
		}

		if len(u.Password) < 8 {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Must be at least 8 letters, alphanumeric + symbol, has at least 1 uppercase letter, has at least 1 number, and has at least 1 symbol."}
		}
		return nil

	default:
		if u.Username == "" || u.Password == "" || u.RePassword == "" {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Required username or password or repassword"}
		}

		if u.Password != u.RePassword {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Password not match"}
		}

		if len(u.Password) < 8 {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Must be at least 8 letters, alphanumeric + symbol, has at least 1 uppercase letter, has at least 1 number, and has at least 1 symbol."}
		}
		return nil
	}
}

func (u *User) SaveAUser(db *gorm.DB) (*User, error) {

	err := db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) UpdateAUser(db *gorm.DB, user_id string) (*User, error) {

	err := u.BeforeSave(db)
	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug().Model(&User{}).Where("user_id = ?", u.UserId).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"username":   u.Username,
			"email":      u.Email,
			"password_1": u.Password,
			"password_2": u.RePassword,
			"update_at":  time.Now(),
		},
	)

	if db.Error != nil {
		return &User{}, db.Error
	}

	err = db.Debug().Model(&User{}).Where("user_id = ?", user_id).Take(&u).Error
	if err != nil {
		return &User{}, nil
	}
	return u, nil
}

func (u *User) DeleteAUser(db *gorm.DB, user_id string) (int64, error) {
	db = db.Debug().Model(&User{}).Where("user_id = ?", user_id).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
