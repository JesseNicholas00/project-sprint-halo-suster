package auth

import (
	"context"

	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/jmoiron/sqlx"
)

func (repo *authRepostioryImpl) FindUserByNip(
	ctx context.Context,
	nip int64,
) (res User, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	query := `
		SELECT
			*
		FROM
			users
		WHERE
			nip = :nip
	`
	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	rows, err := sqlx.NamedQueryContext(
		ctx,
		sess.Ext,
		query,
		map[string]interface{}{
			"nip": nip,
		},
	)

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
		err = ErrNipNotFound
		return
	}

	return
}
