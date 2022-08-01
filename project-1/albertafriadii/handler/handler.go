package handler

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	if DbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

		h.DB, err = gorm.Open(DbDriver, DBURL)

		if err != nil {
			fmt.Printf("Cannot connect to %s database", DbDriver)
			log.Fatal("This is the error: ", err)
			return
		} else {
			fmt.Printf("We are connected to the %s database\n", DbDriver)
		}
	}
}

// const Key = "secret"
