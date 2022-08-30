package repository

import (
	"context"
	"fmt"

<<<<<<< HEAD
	"github.com/albertafriadii/tree/featured/albert-jwt-auth/pkg/config"
	"github.com/albertafriadii/tree/featured/albert-jwt-auth/pkg/domain"
=======
	"github.com/albertafriadii/tree/fix/albertafriadii/pkg/config"
	"github.com/albertafriadii/tree/fix/albertafriadii/pkg/domain"
>>>>>>> 289c4e129d7f4946ed954ea9078f420bf430068c
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

<<<<<<< HEAD
=======
	hashedPassword, err := config.HashPassword(u.Password)
	if err != nil {
		fmt.Println(err)
		return domain.Users{}, nil
	}

	user.Password = hashedPassword

>>>>>>> 289c4e129d7f4946ed954ea9078f420bf430068c
	result := d.db.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		return domain.Users{}, result.Error
	}

	return domain.Users(user), nil
}

<<<<<<< HEAD
func (d *DBUserRepository) UpdateUser(ctx context.Context, u domain.Users, UserID string) (domain.Users, error) {
=======
func (d *DBUserRepository) UpdateUser(ctx context.Context, u domain.Users, user_id string) (domain.Users, error) {
>>>>>>> 289c4e129d7f4946ed954ea9078f420bf430068c
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

<<<<<<< HEAD
	result := d.db.Model(&user).Where("user_id = ?", UserID).Updates(updateUser)
=======
	result := d.db.Model(&user).Where("user_id = ?", user_id).Updates(updateUser)
>>>>>>> 289c4e129d7f4946ed954ea9078f420bf430068c
	if result.Error != nil {
		fmt.Println(result.Error)
		return domain.Users{}, result.Error
	}

	return domain.Users(updateUser), nil
}

<<<<<<< HEAD
func (d *DBUserRepository) DeleteUser(ctx context.Context, UserID string) error {
	var user domain.Users

	result := d.db.Where("user_id = ?", UserID).Delete(&user)
=======
func (d *DBUserRepository) DeleteUser(ctx context.Context, user_id string) error {
	var user domain.Users

	result := d.db.Where("user_id = ?", user_id).Delete(&user)
>>>>>>> 289c4e129d7f4946ed954ea9078f420bf430068c
	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		return config.ErrNotFound
	}
	return nil
}
