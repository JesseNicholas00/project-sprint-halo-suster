package auth

import "context"

type AuthRepository interface {
	CreateUser(ctx context.Context, user User) (User, error)
	FindUserByNip(ctx context.Context, nip int64) (User, error)
	ActivateUserByUserId(ctx context.Context, req ActivateUserReq) (User, error)
	DeleteNurse(ctx context.Context, userId string) error
}
