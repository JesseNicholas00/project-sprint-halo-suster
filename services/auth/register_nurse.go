package auth

import (
	"context"
	"errors"

	"github.com/JesseNicholas00/HaloSuster/repos/auth"
	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/google/uuid"
)

func (svc *authServiceImpl) RegisterNurse(
	ctx context.Context,
	req RegisterNurseReq,
	res *RegisterNurseRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	_, err := svc.repo.FindUserByNip(ctx, req.Nip)

	if err == nil {
		return ErrNipAlreadyExists
	}

	// unexpected kind of error
	if !errors.Is(err, auth.ErrNipNotFound) {
		return errorutil.AddCurrentContext(err)
	}

	repoRes, err := svc.repo.CreateUser(ctx, auth.User{
		Id:       uuid.NewString(),
		Nip:      req.Nip,
		Name:     req.Name,
		Admin:    false,
		Active:   false,
		ImageUrl: req.ImageUrl,
	})
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	*res = RegisterNurseRes{
		UserId: repoRes.Id,
		Nip:    repoRes.Nip,
		Name:   repoRes.Name,
	}

	return nil
}
