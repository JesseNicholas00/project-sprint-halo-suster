package medicalrecord

import (
	"context"

	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const createPatientQuery = `
	INSERT INTO patients(
		identity_number,
		phone_number,
		birth_date,
		gender,
		image_url
	) VALUES (
		:identity_number,
		:phone_number,
		:birth_date,
		:gender,
		:image_url
	)
`

func (repo *medicalRecordRepositoryImpl) CreatePatient(
	ctx context.Context,
	patient Patient,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	if _, err = sqlx.NamedExecContext(
		ctx,
		sess.Ext,
		createPatientQuery,
		patient,
	); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return ErrDuplicateIdentityNumber
		} else {
			return errorutil.AddCurrentContext(err)
		}
	}

	return nil
}
