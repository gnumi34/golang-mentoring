package repository

import (
	"golang-mentoring/project-1/albertafriadii/model"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	SaveAUser(u model.User) (string, error)
	UpdateAUser(u model.User, user_id string) (*model.User, error)
	DeleteAUser(u model.User, user_id string) error
	GetUser(user_id string) (*model.User, error)
}

type Database struct {
	connect *gorm.DB
}

func NewUserRepository() UserRepository {
	if DB == nil {
		_, err = Initialize()
		if err != nil {
			log.Fatal(err)
		}
	}
	return &Database{
		connect: DB,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (db Database) BeforeSave(u model.User) error {
	hash, err := HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

func (db Database) SaveAUser(u model.User) (string, error) {

	result := db.connect.Create(&u)
	if result.Error != nil {
		return "0", result.Error
	}
	return u.UserId, nil
}

func (db Database) GetUser(user_id string) (*model.User, error) {
	var u model.User
	result := db.connect.Preload(clause.Associations).Find(&u, "user_id = ?", user_id)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected > 0 {
		return &u, nil
	}

	return nil, nil
}

func (db Database) UpdateAUser(u model.User, user_id string) (*model.User, error) {

	err := db.BeforeSave(u)
	if err != nil {
		log.Fatal(err)
	}

	err = db.connect.Model(&model.User{}).Where("user_id = ?", user_id).UpdateColumns(
		map[string]interface{}{
			"username":   u.Username,
			"email":      u.Email,
			"password_1": u.Password,
			"password_2": u.RePassword,
			"update_at":  time.Now(),
		},
	).Error

	if db.connect.Error != nil {
		return &model.User{}, db.connect.Error
	}

	return &u, nil
}

func (db Database) DeleteAUser(u model.User, user_id string) error {
	u.UserId = user_id
	result := db.connect.Delete(&u)
	return result.Error
}
