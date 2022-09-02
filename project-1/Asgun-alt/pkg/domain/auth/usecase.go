package auth

import "context"

type UseCase interface {
	Login(ctx context.Context, req *LoginUserRequest) (*Response, error)
}
