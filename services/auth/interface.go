package auth

import "context"

type AuthService interface {
	RegisterIt(
		ctx context.Context,
		req RegisterItReq,
		res *RegisterItRes,
	) error
	Login(
		ctx context.Context,
		req LoginReq,
		res *LoginRes,
	) error
	GetSessionFromToken(
		ctx context.Context,
		req GetSessionFromTokenReq,
		res *GetSessionFromTokenRes,
	) error
	RegisterNurse(
		ctx context.Context,
		req RegisterNurseReq,
		res *RegisterNurseRes,
	) error
	GrantAccessNurse(
		ctx context.Context,
		req GrantAccessNurseReq,
		res *GrantAccessNurseRes,
	) error
}
