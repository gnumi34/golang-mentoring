package auth

import (
	"context"
)

type UseCase interface {
	ValidateUser(ctx context.Context, req *ValidateUserRequest) (*Response, error)
}
