package medicalrecord

import (
	"context"

	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/JesseNicholas00/HaloSuster/utils/helper"
	"github.com/JesseNicholas00/HaloSuster/utils/mewsql"
)

func (repo *medicalRecordRepositoryImpl) ListPatients(
	ctx context.Context,
	filter PatientFilter,
) (res []Patient, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	var conditions []mewsql.Condition

	if filter.IdentityNumber != nil {
		curLen := helper.GetLen(*filter.IdentityNumber)
		lowerBound := *filter.IdentityNumber
		upperBound := *filter.IdentityNumber

		for len := curLen + 1; len <= 16; len++ {
			lowerBound = lowerBound * 10   // xxx0
			upperBound = upperBound*10 + 9 // xxx9
		}
		conditions = append(conditions,
			mewsql.WithCondition("identity_number >= ?", lowerBound),
			mewsql.WithCondition("identity_number <= ?", upperBound),
		)
	}

	if filter.Name != nil {
		conditions = append(
			conditions,
			mewsql.WithCondition("name ILIKE ?", "%"+*filter.Name+"%"),
		)
	}

	if filter.PhoneNumber != nil {
		conditions = append(
			conditions,
			mewsql.WithCondition(
				"phone_number ILIKE ?",
				"%"+*filter.PhoneNumber+"%",
			),
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
	}

	sql, args := mewsql.Select(
		"*",
		"patients",
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
		var cur Patient
		if err = rows.StructScan(&cur); err != nil {
			err = errorutil.AddCurrentContext(err)
			return
		}

		res = append(res, cur)
	}
	return
}
