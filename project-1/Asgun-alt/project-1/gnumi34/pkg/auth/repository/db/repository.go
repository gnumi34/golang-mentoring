package repository

import (
	"context"
	"fmt"

	"github.com/gnumi34/golang-mentoring/project-1/gnumi34/pkg/domain/users"
	"gorm.io/gorm"
)

type AuthDBRepository struct {
	DB *gorm.DB
}

func NewUsersDBRepository(db *gorm.DB) *AuthDBRepository {
	return &AuthDBRepository{
		DB: db,
	}
}

func (r *AuthDBRepository) FindByUserName(ctx context.Context, username string) (*users.User, error) {
	var res users.User

	err := r.DB.WithContext(ctx).First(&res, "username = ?", username).Error
	if err != nil {
		return nil, fmt.Errorf("UsersDBRepository.FindByUserName: %w", err)
	}

	return &res, nil
}
