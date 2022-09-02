package db

import (
	"context"
	"golang-mentoring/project-1/Asgun-alt/pkg/domain/users/response"
	"golang-mentoring/project-1/Asgun-alt/pkg/helper/errcode"

	"gorm.io/gorm"
)

type DBAuthRepository struct {
	db *gorm.DB
}

func NewDBAuthRepository(gormDB *gorm.DB) *DBAuthRepository {
	return &DBAuthRepository{db: gormDB}
}

func (r *DBAuthRepository) Login(ctx context.Context, username string) (*response.User, error) {
	var res response.User

	err := r.db.WithContext(ctx).First(&res, "username = ?", username).Error
	if err != nil {
		return nil, errcode.ErrRecordNotFound
	}

	return &res, nil
}
