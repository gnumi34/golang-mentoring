package repository

import (
	"context"
	"fmt"

	"golang-mentoring/project-1/albertafriadii/pkg/config"
	"golang-mentoring/project-1/albertafriadii/pkg/domain"

	"gorm.io/gorm"
)

type DBUserRepository struct {
	db *gorm.DB
}

func NewUserRepositroy(DB *gorm.DB) domain.UserRepositoryInterface {
	return &DBUserRepository{
		db: DB,
	}
}

func (d *DBUserRepository) GetUser(ctx context.Context, UserID string) (*domain.Users, error) {
	var user domain.Users

	err := d.db.WithContext(ctx).First(&user, "user_id = ?", UserID).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

func (d *DBUserRepository) LoginUser(ctx context.Context, Username string) (*domain.Users, error) {

	var user domain.Users

	err := d.db.WithContext(ctx).First(&user, "username = ?", Username).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

func (d *DBUserRepository) CreateUser(ctx context.Context, u *domain.Users) (*domain.Users, error) {
	user := domain.FromUserDomain(u)

	result := d.db.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	return user, nil
}

func (d *DBUserRepository) UpdateUser(ctx context.Context, u *domain.Users, UserID string) error {

	err := d.db.WithContext(ctx).Where("user_id = ?", UserID).Updates(u).Error
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (d *DBUserRepository) DeleteUser(ctx context.Context, UserID string) error {
	var user domain.Users

	result := d.db.Where("user_id = ?", UserID).Delete(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		return config.ErrNotFound
	}
	return nil
}
