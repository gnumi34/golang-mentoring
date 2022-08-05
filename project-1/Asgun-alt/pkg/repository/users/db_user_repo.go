package dbusers

import (
	"context"

	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/helper/encrypt"
	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/helper/err"
	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/service/uservice"
	"gorm.io/gorm"
)

type DBUserRepository struct {
	db *gorm.DB
}

func NewDBUserRepository(gormDB *gorm.DB) uservice.UsersRepositoryInterface {
	return &DBUserRepository{db: gormDB}
}

func (repo *DBUserRepository) GetUser(ctx context.Context, userDomain uservice.UsersDomain) (uservice.UsersDomain, error) {
	var user Users
	result := repo.db.Where("username = ?", userDomain.Username).Find(&user)
	if result.Error != nil {
		return uservice.UsersDomain{}, result.Error
	}

	err := encrypt.CheckPassword(userDomain.Password, user.Password)
	if err != nil {
		return uservice.UsersDomain{}, result.Error
	}

	return user.ToUserDomain(), nil
}

func (repo *DBUserRepository) AddUsers(ctx context.Context, userDomain uservice.UsersDomain) (uservice.UsersDomain, error) {
	newUser := FromUserDomain(userDomain)

	hashedPassword, err := encrypt.HashPassword(userDomain.Password)
	if err != nil {
		return uservice.UsersDomain{}, err
	}

	newUser.Password = hashedPassword

	result := repo.db.Create(&newUser)
	if result.Error != nil {
		return uservice.UsersDomain{}, result.Error
	}

	return newUser.ToUserDomain(), nil
}

func (repo *DBUserRepository) UpdateUsers(ctx context.Context, userUpdateDomain uservice.UsersDomain) (uservice.UsersDomain, error) {
	var user Users
	updateUser := FromUserDomain(userUpdateDomain)

	hashedPassword, err := encrypt.HashPassword(userUpdateDomain.Password)
	if err != nil {
		return uservice.UsersDomain{}, err
	}

	updateUser.Password = hashedPassword

	resultUpdate := repo.db.Model(&user).Where("id = ?", updateUser.ID).Updates(updateUser)
	if resultUpdate.Error != nil {
		return uservice.UsersDomain{}, resultUpdate.Error
	}
	return updateUser.ToUserDomain(), nil
}

func (repo *DBUserRepository) DeleteUsers(ctx context.Context, id string) error {
	var table Users
	result := repo.db.Where("id = ?", id).Delete(&table)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return err.ErrNotFound
	}
	return nil
}
