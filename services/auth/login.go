package auth

import (
	"context"

	"github.com/JesseNicholas00/HaloSuster/repos/auth"
	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func (svc *authServiceImpl) Login(
	ctx context.Context,
	req LoginReq,
	res *LoginRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	user, err := svc.repo.FindUserByNip(ctx, req.Nip)

	if err != nil {
		switch {
		case errors.Is(err, auth.ErrNipNotFound):
			return ErrUserNotFound

		default:
			return errorutil.AddCurrentContext(err)
		}
	}

	if !user.Active {
		return ErrUserHasNoAccess
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)
	if err != nil {
		return ErrInvalidCredentials
	}

	token, err := svc.generateToken(user)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	*res = LoginRes{
		UserId:      user.Id,
		Nip:         user.Nip,
		Name:        user.Name,
		AccessToken: token,
	}

	return nil
}
