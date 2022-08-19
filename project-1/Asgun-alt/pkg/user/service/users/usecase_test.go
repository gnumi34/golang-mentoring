package users_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/app/middlewares"
	_mockUserRepository "github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/user/mocks/service/users"
	"github.com/gnumi34/golang-mentoring/tree/main/project-1/Asgun-alt/pkg/user/service/users"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var (
	userRepository _mockUserRepository.UsersRepositoryInterface
	userService    users.UsersUsecaseInterface
	userDomain     users.UsersDomain
	usersDomain    []users.UsersDomain
	configJWT      middlewares.ConfigJWT
)

func setup() {
	configJWT = middlewares.ConfigJWT{
		SecretKey:       viper.GetString(`jwt.secret_key`),
		ExpiresDuration: viper.GetInt(`jwt.expire_duration`),
	}
	userService = users.NewUserUseCase(&userRepository, &configJWT)
	userDomain = users.UsersDomain{
		ID:         "1a",
		Username:   "test-username",
		Email:      "test-email@gmail.com",
		Password:   "Test-password1",
		Token:      "test-JWT-token",
		Created_At: time.Now(),
		Updated_At: time.Now(),
	}
	usersDomain = append(usersDomain, userDomain)
}

func Test_Login(t *testing.T) {
	setup()
	t.Run("test case 1 | valid login", func(t *testing.T) {
		ctx := context.Background()
		req := users.UsersDomain{
			Username: "test-username",
			Password: "Test-Password1",
		}

		userRepository.On("Login", ctx, req).Return(userDomain, nil).Once()
		user, err := userService.Login(context.Background(), users.UsersDomain{
			Username: "test-username",
			Password: "Test-Password1",
		})
		assert.Nil(t, err)
		assert.Equal(t, "test-username", user.Username)
	})

	t.Run("test case 2 | username is empty", func(t *testing.T) {
		ctx := context.Background()
		req := users.UsersDomain{
			Username: "",
			Password: "Test-Password1",
		}

		userRepository.On("Login", ctx, req).Return(userDomain, nil).Once()
		_, err := userService.Login(context.Background(), users.UsersDomain{
			Username: "",
			Password: "Test-Password1",
		})

		assert.NotNil(t, err)
	})
}

func Test_GetUser(t *testing.T) {
	setup()
	t.Run("test case 1 | valid user", func(t *testing.T) {
		ctx := context.Background()
		req := users.UsersDomain{
			Username: "test-username",
			Password: "Test-Password1",
		}

		userRepository.On("GetUser", ctx, req).Return(userDomain, nil).Once()
		user, err := userService.GetUser(context.Background(), users.UsersDomain{
			Username: "test-username",
			Password: "Test-Password1",
		})
		assert.Nil(t, err)
		assert.Equal(t, "test-username", user.Username)
	})

	t.Run("test case 2 | invalid username empty", func(t *testing.T) {
		ctx := context.Background()
		req := users.UsersDomain{
			Username: "",
			Password: "Test-Password1",
		}

		userRepository.On("GetUser", ctx, req).Return(userDomain, nil).Once()
		_, err := userService.GetUser(context.Background(), users.UsersDomain{
			Username: "",
			Password: "Test-Password1",
		})

		assert.NotNil(t, err)
	})
}

func Test_AddUsers(t *testing.T) {
	setup()
	t.Run("test case 1 | valid add user", func(t *testing.T) {
		ctx := context.Background()
		req := users.UsersDomain{
			ID:       "0632b87b-9771-4ee9-a18e-392d1208dd46",
			Username: "new-test-username",
			Email:    "newTest-email@gmail.com",
			Password: "New-testPassword1",
		}

		userRepository.On("AddUser", ctx, req).Return(userDomain, nil).Once()
		user, err := userService.AddUser(context.Background(), users.UsersDomain{
			ID:       "0632b87b-9771-4ee9-a18e-392d1208dd46",
			Username: "new-test-username",
			Email:    "newTest-email@gmail.com",
			Password: "New-testPassword1",
		})

		assert.Nil(t, err)
		assert.Equal(t, user.Username, userDomain.Username)
	})
}

func Test_UpdateUsers(t *testing.T) {
	setup()
	t.Run("test case 1", func(t *testing.T) {
		ctx := context.Background()
		req := users.UsersDomain{
			ID:       "test-id",
			Username: "test-new-username",
			Email:    "test-email@gmail.com",
			Password: "Test-password1",
		}

		userRepository.On("UpdateUser", ctx, req).Return(users.UsersDomain{
			ID:         userDomain.ID,
			Username:   "test-new-username",
			Email:      userDomain.Email,
			Password:   userDomain.Password,
			Created_At: userDomain.Created_At,
			Updated_At: userDomain.Updated_At,
		}, nil).Once()
		user, err := userService.UpdateUser(context.Background(), users.UsersDomain{
			ID:       "test-id",
			Username: "test-new-username",
			Email:    "test-email@gmail.com",
			Password: "Test-password1",
		})

		assert.NoError(t, err)
		assert.Equal(t, "test-new-username", user.Username)
	})

	t.Run("test case 2 | id empty", func(t *testing.T) {
		ctx := context.Background()
		req := users.UsersDomain{
			ID:       "",
			Username: "test-username",
			Email:    "test-email@gmail.com",
			Password: "Test-password1",
		}

		userRepository.On("UpdateUser", ctx, req).Return(users.UsersDomain{}, errors.New("Can't find user, id is empty"))
		_, err := userService.UpdateUser(context.Background(), users.UsersDomain{
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
	t.Run("test case 1 | success", func(t *testing.T) {
		ctx := context.Background()
		req := userDomain.ID

		userRepository.On("DeleteUser", ctx, req).Return(nil).Once()
		err := userService.DeleteUser(context.Background(), userDomain.ID)
		assert.NoError(t, err)
	})

	t.Run("test case 2 | error", func(t *testing.T) {
		ctx := context.Background()
		req := userDomain.ID

		userRepository.On("DeleteUser", ctx, req).Return(errors.New("Unexpected Error")).Once()
		err := userService.DeleteUser(context.Background(), userDomain.ID)
		assert.Error(t, err)
	})
}
