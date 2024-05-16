package auth

import (
	"context"

	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/lib/pq"
)

func (repo *authRepositoryImpl) UpdateNurse(
	ctx context.Context,
	user User,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	stmt := sess.NamedStmt(ctx, repo.statements.updateNurseByNurseId)
	result, err := stmt.Exec(user)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return ErrDuplicateUser
		}
		return errorutil.AddCurrentContext(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}
	if rowsAffected == 0 {
		return ErrUserIdNotFound
	}
	return nil
}
