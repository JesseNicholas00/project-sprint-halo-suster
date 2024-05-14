package auth

import (
	"context"

	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func (repo *authRepostioryImpl) CreateUser(
	ctx context.Context,
	user User,
) (res User, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	query := `
		INSERT INTO users(
			user_id,
			nip,
			name,
			password,
			active,
			image_url
		) VALUES (
			:user_id,
			:nip,
			:name,
			:password,
			:active,
			:image_url
		) RETURNING
			user_id,
			nip,
			name,
			password,
			active,
			image_url,
			created_at
	`
	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	rows, err := sqlx.NamedQueryContext(ctx, sess.Ext, query, user)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			err = ErrDuplicateUser
		} else {
			err = errorutil.AddCurrentContext(err)
		}
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

	return
}
