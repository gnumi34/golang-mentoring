package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DB_Host     string
	DB_User     string
	DB_Password string
	DB_Port     string
	DB_Name     string
	DB_SSL      string
}

func (c *Config) IntializeDB() *gorm.DB {
	dsn := fmt.Sprintf("host = %s user = %s password = %s dbname = %s port = %s sslmode = %s", c.DB_Host, c.DB_User, c.DB_Password, c.DB_Name, c.DB_Port, c.DB_SSL)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Cannot connect to database")
		log.Fatal("This is the error: ", err)
	} else {
		fmt.Println("We are connected to the database")
	}

	return db
}
