package medicalrecord

import (
	"context"

	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
)

func (repo *medicalRecordRepositoryImpl) CreateRecord(
	ctx context.Context,
	record Record,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	if _, err := sess.NamedStmt(ctx, repo.statements.createRecord).ExecContext(
		ctx,
		record,
	); err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return nil
}
