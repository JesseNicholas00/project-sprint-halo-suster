package auth

import (
	"context"

	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
)

func (repo *authRepositoryImpl) ActivateUserByUserId(
	ctx context.Context,
	userId string,
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
		Stmt(ctx, repo.statements.activateUserByUserId).
		QueryxContext(ctx, userId)

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
