package modeltest

import (
	"golang-mentoring/project-1/albertafriadii/model"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSaveAUser(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	newUser := model.User{
		UserId:     uuid.NewString(),
		Username:   "test",
		Email:      "test@gmail.com",
		Password:   "Password_1",
		RePassword: "Password_1",
	}
	savedUser, err := newUser.SaveAUser(server.DB)
	if err != nil {
		t.Errorf("this is error getting the users: %v\n", err)
		return
	}

	assert.Equal(t, newUser.UserId, savedUser.UserId)
	assert.Equal(t, newUser.Username, savedUser.Username)
	assert.Equal(t, newUser.Email, savedUser.Email)
}

func TestGetUser(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := oneUser()
	if err != nil {
		log.Fatalf("Cannot get users table: %v\n", err)
	}

	getUser, err := userInit.GetUser(server.DB, user.UserId)
	if err != nil {
		t.Errorf("this is error getting one user: %v\n", err)
		return
	}
	assert.Equal(t, getUser.UserId, user.UserId)
	assert.Equal(t, getUser.Username, user.Username)
}

func TestUpdateAUser(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := oneUser()
	if err != nil {
		log.Fatalf("Cannot get user: %v\n", err)
	}

	userUpdate := model.User{
		UserId:     "5c2913b6-ca4e-4dfb-9a9e-d1678f7b7d7d",
		Username:   "updatetest",
		Email:      "updatetest@gmail.com",
		Password:   "Password_123",
		RePassword: "Password_123",
	}
	updatedUser, err := userUpdate.UpdateAUser(server.DB, user.UserId)
	if err != nil {
		t.Errorf("this is the error updating the user:%v\n", err)
		return
	}
	assert.Equal(t, updatedUser.UserId, userUpdate.UserId)
	assert.Equal(t, updatedUser.Username, userUpdate.Username)
	assert.Equal(t, updatedUser.Email, userUpdate.Email)
}

func TestDeleteAUser(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := oneUser()

	if err != nil {
		log.Fatalf("Cannot get user: %v\n", err)
	}

	isDeleted, err := userInit.DeleteAUser(server.DB, user.UserId)
	if err != nil {
		t.Errorf("this is error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, isDeleted, &userInit)
}
