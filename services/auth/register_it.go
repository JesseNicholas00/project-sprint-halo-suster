package auth

import (
	"context"
	"errors"

	"github.com/JesseNicholas00/HaloSuster/repos/auth"
	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (svc *authServiceImpl) RegisterIt(
	ctx context.Context,
	req RegisterItReq,
	res *RegisterItRes,
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

	cryptedPw, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		svc.bcryptCost,
	)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	repoRes, err := svc.repo.CreateUser(ctx, auth.User{
		Id:       uuid.NewString(),
		Nip:      req.Nip,
		Name:     req.Name,
		Admin:    true,
		Active:   true,
		Password: string(cryptedPw),
	})
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	token, err := svc.generateToken(repoRes)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	*res = RegisterItRes{
		UserId:      repoRes.Id,
		Nip:         repoRes.Nip,
		Name:        repoRes.Name,
		AccessToken: token,
	}

	return nil
}
