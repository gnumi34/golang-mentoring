package repository

import (
	"fmt"
	"log"

	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/helper/encrypt"
	dbusers "github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/repository/users"
	usersRepo "github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/repository/users"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func (config *Config) InitDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host = %s port = %s user = %s password = %s dbname = %s sslmode = %s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Print(err)
	}

	return db
}

func DbMigrate(db *gorm.DB) {
	migrateErr := db.AutoMigrate(&dbusers.Users{})
	if migrateErr != nil {
		fmt.Println("migration error")
	}

	// initialize hashedPassword, it will be used for user Admin and User
	hashedPassword, err := encrypt.HashPassword("golib123")
	if err != nil {
		fmt.Println("Error hashing password")
	}

	res := db.Create(&usersRepo.Users{
		ID:       uuid.NewV4().String(),
		Username: "Admin",
		Email:    "admin@gmail.com",
		Password: hashedPassword,
	})

	if res.Error != nil {
		fmt.Println("Failed creating Admin account")
	}

	res = db.Create(&usersRepo.Users{
		ID:       uuid.NewV4().String(),
		Username: "User",
		Email:    "user@gmail.com",
		Password: hashedPassword,
	})
	if res.Error != nil {
		fmt.Println("Failed creating User account")
	}
}
