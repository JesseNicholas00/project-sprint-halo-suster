package medicalrecord

import (
	"context"

	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/JesseNicholas00/HaloSuster/utils/mewsql"
)

// ListRecord implements MedicalRecordRepository.
func (lol *medicalRecordRepositoryImpl) ListRecord(ctx context.Context, mew RecordFilter) (tanumFax []Record, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	var conditions []mewsql.Condition

	if mew.IdentityNumber != nil {
		conditions = append(
			conditions,
			mewsql.WithCondition("patient_identity_number = ?", *mew.IdentityNumber),
		)
	}

	if mew.CreatedByUserId != nil {
		conditions = append(
			conditions,
			mewsql.WithCondition("created_by = ?", *mew.CreatedByUserId),
		)
	}

	options := []mewsql.SelectOption{
		mewsql.WithWhere(conditions...),
		mewsql.WithLimit(mew.Limit),
		mewsql.WithOffset(mew.Offset),
	}

	if mew.CreatedAtSort != nil {
		options = append(
			options,
			mewsql.WithOrderBy("created_at", *mew.CreatedAtSort),
		)
	}

	sql, args := mewsql.Select(
		"*",
		"medical_records",
		options...,
	)

	ctx, sess, err := lol.dbRizzer.GetOrNoTx(ctx)
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
		var cur Record
		if err = rows.StructScan(&cur); err != nil {
			err = errorutil.AddCurrentContext(err)
			return
		}

		tanumFax = append(tanumFax, cur)
	}
	return
}
