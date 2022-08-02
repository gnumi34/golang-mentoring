package dbusers

import (
	"context"

	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/helper/encrypt"
	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/service/uservice"
	"gorm.io/gorm"
)

type DBUserRepository struct {
	db *gorm.DB
}

func NewDBUserRepository(gormDB *gorm.DB) uservice.UsersRepositoryInterface {
	return &DBUserRepository{db: gormDB}
}

func (repo DBUserRepository) AddUsers(ctx context.Context, userDomain uservice.UsersDomain) (uservice.UsersDomain, error) {
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

func (repo DBUserRepository) UpdateUsers(ctx context.Context, userUpdateDomain uservice.UsersDomain) (uservice.UsersDomain, error) {
	var user Users
	updateUser := FromUserDomain(userUpdateDomain)

	resultUpdate := repo.db.Model(&user).Where("id = ?", updateUser.ID).Updates(updateUser)
	if resultUpdate.Error != nil {
		return uservice.UsersDomain{}, resultUpdate.Error
	}
	return updateUser.ToUserDomain(), nil
}
