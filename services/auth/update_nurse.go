package auth

import (
	"context"

	"github.com/JesseNicholas00/HaloSuster/repos/auth"
	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/JesseNicholas00/HaloSuster/utils/transaction"
)

func (svc *authServiceImpl) UpdateNurse(
	ctx context.Context,
	req UpdateNurseReq,
	res *UpdateNurseRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	ctx, sess, err := svc.dbRizz.AppendTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}
	return transaction.RunWithAutoCommit(&sess, func() error {
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
	})
}
