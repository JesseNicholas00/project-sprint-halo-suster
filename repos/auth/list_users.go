package auth

import (
	"context"

	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/JesseNicholas00/HaloSuster/utils/mewsql"
)

func (repo *authRepositoryImpl) ListUsers(
	ctx context.Context,
	filter UserFilter,
) (res []User, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	var conditions []mewsql.Condition

	if filter.UserId != nil {
		conditions = append(
			conditions,
			mewsql.WithCondition("user_id= ?", *filter.UserId),
		)
	}

	if filter.Name != nil {
		conditions = append(
			conditions,
			mewsql.WithCondition("name ILIKE ?", "%"+*filter.Name+"%"),
		)
	}

	if filter.Nip != nil {
		conditions = append(
			conditions,
			mewsql.WithCondition("nip >= ?", *filter.Nip),
		)
	}

	if filter.Role != nil {
		conditions = append(
			conditions,
			mewsql.WithCondition("is_admin IS ?", *filter.Role == "it"),
		)
	}

	options := []mewsql.SelectOption{
		mewsql.WithWhere(conditions...),
		mewsql.WithLimit(filter.Limit),
		mewsql.WithOffset(filter.Offset),
	}

	if filter.CreatedAtSort != nil {
		options = append(
			options,
			mewsql.WithOrderBy("created_at", *filter.CreatedAtSort),
		)
	} else {
		options = append(
			options,
			mewsql.WithOrderBy("created_at", "desc"),
		)
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

		res = append(res, cur)
	}
	return
}
