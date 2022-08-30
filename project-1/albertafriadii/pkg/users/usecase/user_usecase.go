package usecase

import (
	"context"
	"fmt"
	"time"

<<<<<<< HEAD
	"github.com/albertafriadii/tree/featured/albert-jwt-auth/pkg/config"
	"github.com/albertafriadii/tree/featured/albert-jwt-auth/pkg/domain"
=======
	"github.com/albertafriadii/tree/fix/albertafriadii/pkg/config"
	"github.com/albertafriadii/tree/fix/albertafriadii/pkg/domain"
>>>>>>> 289c4e129d7f4946ed954ea9078f420bf430068c
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
)

type userUsecase struct {
<<<<<<< HEAD
	userRepo domain.UserRepositoryInterface
}

func NewUserUsecase(u domain.UserRepositoryInterface) domain.UserUsecaseInterface {
	return &userUsecase{
		userRepo: u,
=======
	userRepo       domain.UserRepositoryInterface
	contextTimeout time.Duration
}

func NewUserUsecase(u domain.UserRepositoryInterface, t time.Duration) domain.UserUsecaseInterface {
	return &userUsecase{
		userRepo:       u,
		contextTimeout: t,
>>>>>>> 289c4e129d7f4946ed954ea9078f420bf430068c
	}
}

func (uc *userUsecase) GetUser(ctx context.Context, u domain.Users) (domain.Users, error) {
<<<<<<< HEAD
=======
	if u.Username == "" {
		return domain.Users{}, config.ErrUsernameEmpty
	}
	if u.Password == "" {
		return domain.Users{}, config.ErrPasswordEmpty
	}
>>>>>>> 289c4e129d7f4946ed954ea9078f420bf430068c

	user, err := uc.userRepo.GetUser(ctx, u)
	if err != nil {
		// fmt.Println(err)
		return domain.Users{}, err
	}
	return user, nil
}

func (uc *userUsecase) LoginUser(ctx context.Context, u domain.Users) (string, error) {
<<<<<<< HEAD
=======
	if u.Username == "" {
		return "", config.ErrUsernameEmpty
	}
	if u.Password == "" {
		return "", config.ErrPasswordEmpty
	}
>>>>>>> 289c4e129d7f4946ed954ea9078f420bf430068c

	// generate token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = u.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	_, err = uc.userRepo.LoginUser(ctx, u)
	if err != nil {
		return "", err
	}

	return t, nil
}

func (uc *userUsecase) CreateUser(ctx context.Context, u domain.Users) (domain.Users, error) {
<<<<<<< HEAD

	var err error

=======
>>>>>>> 289c4e129d7f4946ed954ea9078f420bf430068c
	if u.UserId == "" {
		u.UserId = uuid.NewV4().String()
	}

<<<<<<< HEAD
	u.Password, err = config.HashPassword(u.Password)
	if err != nil {
		fmt.Println(err)
		return domain.Users{}, nil
	}

	u, err = uc.userRepo.CreateUser(ctx, u)
=======
	user, err := uc.userRepo.CreateUser(ctx, u)
>>>>>>> 289c4e129d7f4946ed954ea9078f420bf430068c
	if err != nil {
		fmt.Println(err)
		return domain.Users{}, err
	}
<<<<<<< HEAD
	return u, nil
}

func (uc *userUsecase) UpdateUser(ctx context.Context, u domain.Users, UserID string) (domain.Users, error) {

	user, err := uc.userRepo.UpdateUser(ctx, u, UserID)
	if err != nil {
		fmt.Println(err)
		return domain.Users{}, err
	}

	return user, nil
}

func (uc *userUsecase) DeleteUser(ctx context.Context, UserID string) error {
	err := uc.userRepo.DeleteUser(ctx, UserID)
=======
	return user, nil
}

func (uc *userUsecase) UpdateUser(ctx context.Context, u domain.Users, user_id string) (domain.Users, error) {
	if u.Username == "" {
		return domain.Users{}, config.ErrUsernameEmpty
	}

	user, err := uc.userRepo.UpdateUser(ctx, u, user_id)
	if err != nil {
		fmt.Println(err)
		return domain.Users{}, err
	}

	if (user.CreatedAt == time.Time{} && user.Email == "") {
		return domain.Users{}, config.ErrNotFound
	}
	return user, nil
}

func (uc *userUsecase) DeleteUser(ctx context.Context, user_id string) error {
	err := uc.userRepo.DeleteUser(ctx, user_id)
>>>>>>> 289c4e129d7f4946ed954ea9078f420bf430068c
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
