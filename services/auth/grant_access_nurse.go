package auth

import (
	"context"
	"errors"

	"github.com/JesseNicholas00/HaloSuster/repos/auth"
	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
)

func (svc *authServiceImpl) GrantAccessNurse(
	ctx context.Context,
	req GrantAccessNurseReq,
	res *GrantAccessNurseRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	_, err := svc.repo.ActivateUserByUserId(ctx, req.UserId)

	if err == auth.ErrUserIdNotFound {
		return ErrUserNotFound
	}

	if !errors.Is(err, auth.ErrNipNotFound) {
		return errorutil.AddCurrentContext(err)
	}

	return nil
}
