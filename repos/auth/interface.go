package auth

import "context"

type AuthRepository interface {
	CreateUser(ctx context.Context, user User) (User, error)
	ListUsers(ctx context.Context, filter UserFilter) ([]User, error)
	ListAllUsers(ctx context.Context, filter AllUsersFilter) ([]User, error)
	FindUserByNip(ctx context.Context, nip int64) (User, error)
	ActivateNurseByUserId(ctx context.Context, req ActivateUserReq) (User, error)
	UpdateNurse(ctx context.Context, user User) error
	DeleteNurse(ctx context.Context, userId string) error
}
