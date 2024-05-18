package auth

import (
	"context"

	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
)

func (repo *authRepositoryImpl) DeleteNurse(
	ctx context.Context,
	userId string,
) (err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	results, err := sess.
		Stmt(ctx, repo.statements.deleteNurseByUserId).
		ExecContext(ctx, userId)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}
	if rowsAffected == 0 {
		return ErrUserIdNotFound
	}

	return nil
}
