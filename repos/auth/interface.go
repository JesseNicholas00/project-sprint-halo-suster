package auth

import "context"

type AuthRepository interface {
	CreateUser(ctx context.Context, user User) (User, error)
	FindUserByNip(ctx context.Context, nip int64) (User, error)
}
