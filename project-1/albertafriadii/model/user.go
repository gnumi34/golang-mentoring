package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserId     uuid.UUID `gorm:"column:user_id" json:"user_id,omitempty"`
	Username   string    `gorm:"column:username" json:"username,omitempty"`
	Email      string    `gorm:"column:email" json:"email"`
	Password   string    `gorm:"column:password" json:"password_1,omitempty"`
	RePassword string    `json:"password_2,omitempty"`
	Token      string    `json:"token,omitempty" bson:"-"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
