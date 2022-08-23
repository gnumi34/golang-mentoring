package auth

import (
	"context"

	"github.com/gnumi34/golang-mentoring/project-1/gnumi34/pkg/domain/users"
)

type DBRepository interface {
	FindByUserName(ctx context.Context, username string) (*users.User, error)
}
