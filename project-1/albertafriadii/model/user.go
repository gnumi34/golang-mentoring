package model

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserId     string         `gorm:"column:user_id" json:"user_id,omitempty"`
	Username   string         `gorm:"column:username" json:"username,omitempty"`
	Email      string         `gorm:"column:email" json:"email"`
	Password   string         `gorm:"column:password" json:"password_1,omitempty"`
	RePassword string         `gorm:"-" json:"password_2,omitempty"`
	CreatedAt  time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 7)
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

func (u *User) SaveAUser(db *gorm.DB) (*User, error) {

	err := db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) GetUser(db *gorm.DB, user_id string) (*User, error) {
	result := db.Debug().Where("user_id = ?", user_id).First(&u)
	return u, result.Error
}

func (u *User) UpdateAUser(db *gorm.DB, user_id string) (*User, error) {

	err := u.BeforeSave(db)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().Model(&u).Where("user_id = ?", user_id).Updates(
		map[string]interface{}{
			"username":    u.Username,
			"email":       u.Email,
			"password":    u.Password,
			"re_password": u.RePassword,
		},
	).Error

	if db.Error != nil {
		return &User{}, db.Error
	}

	return u, nil
}

func (u *User) DeleteAUser(db *gorm.DB, user_id string) (*User, error) {
	if err := db.Debug().Where("user_id = ? ", user_id).Delete(&u).Error; err != nil {
		return &User{}, err
	}
	return u, nil
}
