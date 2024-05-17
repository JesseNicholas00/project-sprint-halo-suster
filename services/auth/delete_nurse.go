package auth

import (
	"context"

	"github.com/JesseNicholas00/HaloSuster/repos/auth"
	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
)

func (svc *authServiceImpl) DeleteNurse(
	ctx context.Context,
	req DeleteNurseReq,
	res *DeleteNurseRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	err := svc.repo.DeleteNurse(ctx, req.UserId)
	if err != nil {
		if err == auth.ErrUserIdNotFound {
			return ErrUserNotFound
		}
		return errorutil.AddCurrentContext(err)
	}

	return nil
}
