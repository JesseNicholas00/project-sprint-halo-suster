package auth

import (
	"context"

	"github.com/JesseNicholas00/HaloSuster/repos/auth"
	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
)

func (svc *authServiceImpl) UpdateNurse(
	ctx context.Context,
	req UpdateNurseReq,
	res *UpdateNurseRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	err := svc.repo.UpdateNurse(ctx, auth.User{
		Id:   req.UserId,
		Nip:  req.Nip,
		Name: req.Name,
	})
	if err != nil {
		if err == auth.ErrUserIdNotFound {
			return ErrUserNotFound
		}
		return errorutil.AddCurrentContext(err)
	}

	return nil
}
