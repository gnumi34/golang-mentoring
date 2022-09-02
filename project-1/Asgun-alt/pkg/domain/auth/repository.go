package auth

import (
	"context"
	"golang-mentoring/project-1/Asgun-alt/pkg/domain/users/response"
)

type Repository interface {
	Login(ctx context.Context, username string) (*response.User, error)
}
