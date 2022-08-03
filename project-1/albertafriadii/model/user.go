package model

import (
	"time"

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
