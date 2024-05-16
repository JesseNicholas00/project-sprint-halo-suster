package auth

import (
	"context"

	"github.com/JesseNicholas00/HaloSuster/repos/auth"
	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"golang.org/x/crypto/bcrypt"
)

func (svc *authServiceImpl) GrantAccessNurse(
	ctx context.Context,
	req GrantAccessNurseReq,
	res *GrantAccessNurseRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	cryptedPw, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		svc.bcryptCost,
	)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	_, err = svc.repo.ActivateNurseByUserId(ctx, auth.ActivateUserReq{
		Id:       req.UserId,
		Password: string(cryptedPw),
	})

	if err != nil {
		if err == auth.ErrUserIdNotFound {
			return ErrUserNotFound
		}
		return errorutil.AddCurrentContext(err)
	}

	return nil
}
