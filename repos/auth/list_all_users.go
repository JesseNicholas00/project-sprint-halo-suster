package auth

import (
	"context"

	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/JesseNicholas00/HaloSuster/utils/mewsql"
	"github.com/lib/pq"
)

// ListAllUsers implements AuthRepository.
func (repo *authRepositoryImpl) ListAllUsers(
	ctx context.Context,
	filter AllUsersFilter,
) (mewers []User, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	var conditions []mewsql.Condition

	if len(filter.UserIds) > 0 {
		conditions = append(
			conditions,
			mewsql.WithCondition("user_id = ANY (?)", pq.Array(filter.UserIds)),
		)
	}

	options := []mewsql.SelectOption{
		mewsql.WithWhere(conditions...),
	}
	sql, args := mewsql.Select(
		"*",
		"users",
		options...,
	)

	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	rows, err := sess.Ext.QueryxContext(ctx, sql, args...)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var cur User
		if err = rows.StructScan(&cur); err != nil {
			err = errorutil.AddCurrentContext(err)
			return
		}

		mewers = append(mewers, cur)
	}
	return
}
