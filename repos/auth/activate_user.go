package auth

import (
	"context"

	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
)

func (repo *authRepositoryImpl) ActivateNurseByUserId(
	ctx context.Context,
	req ActivateUserReq,
) (res User, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	rows, err := sess.
		NamedStmt(ctx, repo.statements.activateNurseByUserId).
		QueryxContext(ctx, map[string]interface{}{
			"user_id":  req.Id,
			"password": req.Password,
			"active":   true,
		})

	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(&res)
		if err != nil {
			err = errorutil.AddCurrentContext(err)
			return
		}
	}

	if res.Id == "" {
		err = ErrUserIdNotFound
		return
	}

	return
}
