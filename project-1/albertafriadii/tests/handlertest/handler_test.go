package handlertest

import (
	"fmt"
	"golang-mentoring/project-1/albertafriadii/handlers"
	"golang-mentoring/project-1/albertafriadii/model"
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	server   = handlers.Handler{}
	userInit = model.User{}
)

func TestMain(m *testing.M) {
	err := godotenv.Load(os.ExpandEnv("../../cmd/.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	os.Exit(m.Run())
}

func Database() {
	var err error

	TestDbDriver := os.Getenv("TestDbDriver")

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbUser"), os.Getenv("TestDbName"), os.Getenv("TestDbPassword"))

	server.DB, err = gorm.Open(postgres.Open(DBURL))

	if err != nil {
		fmt.Printf("Cannot connect to %s database", TestDbDriver)
		log.Fatal("This is the error: ", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", TestDbDriver)
	}
}

func refreshUserTable() error {
	err := server.DB.Find(&model.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func oneUser() (model.User, error) {
	err := refreshUserTable()
	if err != nil {
		log.Fatal("Not Found")
	}

	user := model.User{
		UserId:     uuid.NewString(),
		Username:   "player",
		Email:      "player@gmail.com",
		Password:   "123Play456_",
		RePassword: "123Play456_",
	}

	err = server.DB.Model(&model.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	return user, nil
}
