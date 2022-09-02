package dbusers

import (
	"context"

	"golang-mentoring/project-1/Asgun-alt/pkg/domain/users"
	"golang-mentoring/project-1/Asgun-alt/pkg/domain/users/response"
	"golang-mentoring/project-1/Asgun-alt/pkg/helper/encrypt"
	"golang-mentoring/project-1/Asgun-alt/pkg/helper/errcode"

	"gorm.io/gorm"
)

type DBUserRepository struct {
	db *gorm.DB
}

func NewDBUserRepository(gormDB *gorm.DB) users.UsersRepositoryInterface {
	return &DBUserRepository{db: gormDB}
}

func (repo *DBUserRepository) FindUserByID(ctx context.Context, id uint) (*users.UsersDomain, error) {
	var res response.User

	err := repo.db.WithContext(ctx).First(&res, "id = ?", id).Error
	if err != nil {
		return nil, errcode.ErrUserNotFound
	}

	return (*users.UsersDomain)(&res), nil
}

func (repo *DBUserRepository) GetUser(ctx context.Context, userDomain *users.UsersDomain) (*users.UsersDomain, error) {
	var user response.User

	result := repo.db.Where("username = ?", userDomain.Username).Find(&user)
	if result.Error != nil {
		return nil, errcode.ErrRecordNotFound
	}

	return (*users.UsersDomain)(&user), nil
}

func (repo *DBUserRepository) AddUser(ctx context.Context, userDomain *users.UsersDomain) (*users.UsersDomain, error) {
	newUser := response.FromUserDomain(userDomain)

	hashedPassword, err := encrypt.HashPassword(userDomain.Password)
	if err != nil {
		return nil, err
	}

	newUser.Password = hashedPassword

	result := repo.db.Create(&newUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return (*users.UsersDomain)(newUser), nil
}

func (repo *DBUserRepository) UpdateUser(ctx context.Context, userUpdateDomain *users.UsersDomain) (*users.UsersDomain, error) {
	var user response.User

	resultUpdate := repo.db.Model(&user).Where("id = ?", userUpdateDomain.ID).Updates(&userUpdateDomain)
	if resultUpdate.Error != nil {
		return nil, resultUpdate.Error
	}

	return (*users.UsersDomain)(userUpdateDomain), nil
}

func (repo *DBUserRepository) DeleteUser(ctx context.Context, id uint) error {
	var table response.User
	result := repo.db.Where("id = ?", id).Delete(&table)
	if result.Error != nil {
		return errcode.ErrUserNotFound
	}
	if result.RowsAffected == 0 {
		return errcode.ErrNotFound
	}
	return nil
}
