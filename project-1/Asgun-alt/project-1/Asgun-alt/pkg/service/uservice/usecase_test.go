package uservice_test

import (
	"context"
	"errors"
	"testing"
	"time"

	_mockUserRepository "github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/domain/mocks"
	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/service/uservice"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userRepository _mockUserRepository.UsersRepositoryInterface
	userService    uservice.UsersUsecaseInterface
	userDomain     uservice.UsersDomain
	usersDomain    []uservice.UsersDomain
)

func setup() {
	userService = uservice.NewUserUseCase(&userRepository, time.Hour*1)
	userDomain = uservice.UsersDomain{
		ID:         "test-id",
		Username:   "test-username",
		Email:      "test-email@gmail.com",
		Password:   "Test-password1",
		Created_At: time.Now(),
		Updated_At: time.Now(),
	}
	usersDomain = append(usersDomain, userDomain)
}

func Test_GetUser(t *testing.T) {
	setup()
	userRepository.On("GetUser",
		mock.Anything,
		mock.AnythingOfType("uservice.UsersDomain")).Return(userDomain, nil).Once()

	t.Run("test case 1 | valid user", func(t *testing.T) {
		user, err := userService.GetUser(context.Background(), uservice.UsersDomain{
			Username: "test-username",
			Password: "Test-password1",
		})

		assert.Nil(t, err)
		assert.Equal(t, "test-username", user.Username)
	})

	t.Run("test case 2 | invalid username empty", func(t *testing.T) {
		_, err := userService.GetUser(context.Background(), uservice.UsersDomain{
			Username: "",
			Password: "Test-password1",
		})

		assert.NotNil(t, err)
	})
}

func Test_AddUsers(t *testing.T) {
	setup()
	t.Run("test case 1 | valid add user", func(t *testing.T) {
		userRepository.On("AddUsers",
			mock.Anything,
			mock.AnythingOfType("uservice.UsersDomain")).Return(userDomain, nil).Once()
		user, err := userService.AddUsers(context.Background(), uservice.UsersDomain{
			Username: "test-username",
			Email:    "test-email@gmail.com",
			Password: "Test-password1",
		})

		assert.Nil(t, err)
		assert.Equal(t, "test-username", user.Username)
	})

	t.Run("test case 2 | error add user", func(t *testing.T) {
		userRepository.On("AddUsers",
			mock.Anything,
			mock.AnythingOfType("uservice.UsersDomain")).Return(uservice.UsersDomain{}, errors.New("Unexpected Error")).Once()
		user, err := userService.AddUsers(context.Background(), uservice.UsersDomain{})

		assert.Error(t, err)
		assert.Equal(t, user, uservice.UsersDomain{})
	})
}

func Test_UpdateUsers(t *testing.T) {
	setup()
	t.Run("test case 1", func(t *testing.T) {
		userRepository.On("UpdateUsers",
			mock.Anything,
			mock.AnythingOfType("uservice.UsersDomain")).Return(uservice.UsersDomain{
			ID:         userDomain.ID,
			Username:   "test-new-username",
			Email:      userDomain.Email,
			Password:   userDomain.Password,
			Created_At: userDomain.Created_At,
			Updated_At: userDomain.Updated_At,
		}, nil).Once()

		user, err := userService.UpdateUsers(context.Background(), uservice.UsersDomain{
			ID:       "test-id",
			Username: "test-new-username",
			Email:    "test-email@gmail.com",
			Password: "Test-password1",
		})

		assert.NoError(t, err)
		assert.Equal(t, "test-new-username", user.Username)
	})

	t.Run("test case 2 | id empty", func(t *testing.T) {
		_, err := userService.UpdateUsers(context.Background(), uservice.UsersDomain{
			ID:       "",
			Username: "test-username",
			Email:    "test-email@gmail.com",
			Password: "Test-password1",
		})

		assert.NotNil(t, err)
	})
}

func Test_DeleteUsers(t *testing.T) {
	setup()
	t.Run("test case 1", func(t *testing.T) {
		userRepository.On("DeleteUsers",
			mock.Anything,
			mock.AnythingOfType("string")).Return(nil).Once()

		err := userService.DeleteUsers(context.Background(), userDomain.ID)
		assert.NoError(t, err)
	})

	t.Run("test case 2 | error", func(t *testing.T) {
		userRepository.On("DeleteUsers",
			mock.Anything,
			mock.AnythingOfType("string")).Return(errors.New("Unexpected Error")).Once()

		err := userService.DeleteUsers(context.Background(), userDomain.ID)
		assert.Error(t, err)
	})
}
