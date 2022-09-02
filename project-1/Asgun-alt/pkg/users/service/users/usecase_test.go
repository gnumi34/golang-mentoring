package users_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"golang-mentoring/project-1/Asgun-alt/pkg/domain/users"
	userUseCase "golang-mentoring/project-1/Asgun-alt/pkg/users/service/users"
	_mockUserRepository "golang-mentoring/project-1/Asgun-alt/pkg/users/service/users/mocks/users"

	"github.com/stretchr/testify/assert"
)

var (
	userRepository _mockUserRepository.UsersRepositoryInterface
	userService    users.UsersUsecaseInterface
	userDomain     *users.UsersDomain
	usersDomain    []users.UsersDomain
)

func setup() {
	userService = userUseCase.NewUserUseCase(&userRepository)
	userDomain = &users.UsersDomain{
		ID:        uint(1),
		Username:  "test-username",
		Email:     "test-email@gmail.com",
		Password:  "Test-password1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	usersDomain = append(usersDomain, *userDomain)
}

func Test_GetUser(t *testing.T) {
	setup()
	t.Run("test case 1 | valid user", func(t *testing.T) {
		ctx := context.Background()
		req := &users.UsersDomain{
			Username: "test-username",
			Password: "Test-password1",
		}

		userRepository.On("GetUser", ctx, req).Return(userDomain, nil).Once()
		_, err := userService.GetUser(ctx, &users.UsersDomain{
			Username: "test-username",
			Password: "Test-password1",
		})
		assert.Nil(t, err)
	})

	t.Run("test case 2 | invalid username empty", func(t *testing.T) {
		ctx := context.Background()
		req := &users.UsersDomain{
			Username: "",
			Password: "Test-Password1",
		}

		userRepository.On("GetUser", ctx, req).Return(userDomain, nil).Once()
		_, err := userService.GetUser(ctx, &users.UsersDomain{
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
		req := &users.UsersDomain{
			ID:       uint(1),
			Username: "new-test-username",
			Email:    "newTest-email@gmail.com",
			Password: "New-testPassword1",
		}

		userRepository.On("AddUser", ctx, req).Return(userDomain, nil).Once()
		user, err := userService.AddUser(ctx, &users.UsersDomain{
			ID:       uint(1),
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
		req := &users.UsersDomain{
			ID:       uint(1),
			Username: "test-new-username",
			Email:    "test-email@gmail.com",
			Password: "Test-password1",
		}

		userRepository.On("UpdateUser", ctx, req).Return(&users.UsersDomain{
			ID:        userDomain.ID,
			Username:  "test-new-username",
			Email:     userDomain.Email,
			Password:  userDomain.Password,
			CreatedAt: userDomain.CreatedAt,
			UpdatedAt: userDomain.UpdatedAt,
		}, nil).Once()

		user, err := userService.UpdateUser(ctx, userDomain.Password, &users.UsersDomain{
			ID:       uint(1),
			Username: "test-new-username",
			Email:    "test-email@gmail.com",
			Password: "Test-password1",
		})

		assert.NoError(t, err)
		assert.Equal(t, "test-new-username", user.Username)
	})

	t.Run("test case 2 | id empty", func(t *testing.T) {
		ctx := context.Background()
		req := &users.UsersDomain{
			ID:       uint(1),
			Username: "test-username",
			Email:    "test-email@gmail.com",
			Password: "Test-password1",
		}

		userRepository.On("UpdateUser", ctx, req).Return(nil, errors.New("Can't find user, id is empty"))
		_, err := userService.UpdateUser(ctx, userDomain.Password, &users.UsersDomain{
			ID:       uint(1),
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
		err := userService.DeleteUser(ctx, userDomain.ID)
		assert.NoError(t, err)
	})

	t.Run("test case 2 | error", func(t *testing.T) {
		ctx := context.Background()
		req := userDomain.ID

		userRepository.On("DeleteUser", ctx, req).Return(errors.New("Unexpected Error")).Once()
		err := userService.DeleteUser(ctx, userDomain.ID)
		assert.Error(t, err)
	})
}
