package dbusers

import (
	"context"

	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/domain/response"
	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/helper/encrypt"
	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/helper/errcode"
	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/user/service/users"
	"gorm.io/gorm"
)

type DBUserRepository struct {
	db *gorm.DB
}

func NewDBUserRepository(gormDB *gorm.DB) users.UsersRepositoryInterface {
	return &DBUserRepository{db: gormDB}
}

func (repo *DBUserRepository) GetUser(ctx context.Context, userDomain users.UsersDomain) (users.UsersDomain, error) {
	var user response.User
	result := repo.db.Where("username = ?", userDomain.Username).Find(&user)
	if result.Error != nil {
		return users.UsersDomain{}, result.Error
	}

	err := encrypt.CheckPassword(userDomain.Password, user.Password)
	if err != nil {
		return users.UsersDomain{}, result.Error
	}

	return users.UsersDomain(user), nil
}

func (repo *DBUserRepository) AddUser(ctx context.Context, userDomain users.UsersDomain) (users.UsersDomain, error) {
	newUser := response.FromUserDomain(userDomain)

	hashedPassword, err := encrypt.HashPassword(userDomain.Password)
	if err != nil {
		return users.UsersDomain{}, err
	}

	newUser.Password = hashedPassword

	result := repo.db.Create(&newUser)
	if result.Error != nil {
		return users.UsersDomain{}, result.Error
	}

	return users.UsersDomain(newUser), nil
}

func (repo *DBUserRepository) UpdateUser(ctx context.Context, userUpdateDomain users.UsersDomain) (users.UsersDomain, error) {
	var user response.User
	updateUser := response.FromUserDomain(userUpdateDomain)

	if updateUser.Password != "" {
		hashedPassword, err := encrypt.HashPassword(userUpdateDomain.Password)
		if err != nil {
			return users.UsersDomain{}, err
		}

		updateUser.Password = hashedPassword
	}

	resultUpdate := repo.db.Model(&user).Where("id = ?", updateUser.ID).Updates(updateUser)
	if resultUpdate.Error != nil {
		return users.UsersDomain{}, resultUpdate.Error
	}
	return users.UsersDomain(updateUser), nil
}

func (repo *DBUserRepository) DeleteUser(ctx context.Context, id string) error {
	var table response.User
	result := repo.db.Where("id = ?", id).Delete(&table)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errcode.ErrNotFound
	}
	return nil
}
