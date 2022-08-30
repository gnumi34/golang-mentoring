package repository

import (
	"context"
	"fmt"

	"github.com/albertafriadii/tree/featured/albert-jwt-auth/pkg/config"
	"github.com/albertafriadii/tree/featured/albert-jwt-auth/pkg/domain"
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

func (d *DBUserRepository) GetUser(ctx context.Context, u domain.Users) (domain.Users, error) {
	var user domain.Users

	result := d.db.Where("username = ?", u.Username).Find(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		return domain.Users{}, result.Error
	}

	_, err := config.CheckPassword(u.Password, user.Password)

	if err != nil {
		// fmt.Println(err)
		return domain.Users{}, err
	}

	return domain.Users(user), nil
}

func (d *DBUserRepository) LoginUser(ctx context.Context, u domain.Users) (bool, error) {

	var user domain.Users

	result := d.db.Where("username = ?", u.Username).Find(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		return false, result.Error
	}

	_, err := config.CheckPassword(u.Password, user.Password)

	if err != nil {
		// fmt.Println(err)
		return false, err
	}

	return true, nil
}

func (d *DBUserRepository) CreateUser(ctx context.Context, u domain.Users) (domain.Users, error) {
	user := domain.FromUserDomain(u)

	result := d.db.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		return domain.Users{}, result.Error
	}

	return domain.Users(user), nil
}

func (d *DBUserRepository) UpdateUser(ctx context.Context, u domain.Users, UserID string) (domain.Users, error) {
	var user domain.Users
	updateUser := domain.FromUserDomain(u)

	if updateUser.Password != "" {
		hashedPassword, err := config.HashPassword(updateUser.Password)
		if err != nil {
			fmt.Println(err)
			return domain.Users{}, err
		}
		updateUser.Password = hashedPassword
	}

	result := d.db.Model(&user).Where("user_id = ?", UserID).Updates(updateUser)
	if result.Error != nil {
		fmt.Println(result.Error)
		return domain.Users{}, result.Error
	}

	return domain.Users(updateUser), nil
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
