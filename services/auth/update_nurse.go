package auth

import (
	"context"
	"errors"

	"github.com/JesseNicholas00/HaloSuster/repos/auth"
	"github.com/JesseNicholas00/HaloSuster/types/nip"
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
	_, err := svc.repo.FindUserByNip(ctx, req.Nip)
	if err == nil {
		return ErrNipAlreadyExists
	}
	if nip.GetRole(req.Nip) != nip.RoleNurse {
		return ErrUserNotFound
	}
	err = svc.repo.UpdateNurse(ctx, auth.User{
		Id:   req.UserId,
		Nip:  req.Nip,
		Name: req.Name,
	})
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrUserIdNotFound):
			return ErrUserNotFound
		case errors.Is(err, auth.ErrDuplicateUser):
			return ErrNipAlreadyExists
		default:
			return errorutil.AddCurrentContext(err)
		}
	}

	return nil
}
