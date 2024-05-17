package medicalrecord

import (
	"context"

	"github.com/JesseNicholas00/HaloSuster/repos/medicalrecord"
	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/JesseNicholas00/HaloSuster/utils/transaction"
)

func (svc *medicalRecordServiceImpl) CreateRecord(
	ctx context.Context,
	req CreateRecordReq,
	res *CreateRecordRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := svc.dbRizzer.GetOrAppendTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return transaction.RunWithAutoCommit(&sess, func() error {
		matchingPatients, err := svc.repo.ListPatients(
			ctx,
			medicalrecord.PatientFilter{
				IdentityNumber: &req.IdentityNumber,
				Limit:          1,
				Offset:         0,
			},
		)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		if len(matchingPatients) != 1 {
			return ErrIdentityNumberNotFound
		}

		patient := matchingPatients[0]

		if err := svc.repo.CreateRecord(ctx, medicalrecord.Record{
			PatientId:          patient.IdentityNumber,
			PatientPhoneNumber: patient.PhoneNumber,
			PatientName:        patient.Name,
			PatientBirthDate:   patient.BirthDate,
			PatientGender:      patient.Gender,
			Symptoms:           req.Symptoms,
			Medications:        req.Medications,
			CreatedByUserId:    req.CreatedById,
		}); err != nil {
			return errorutil.AddCurrentContext(err)
		}

		return nil
	})
}
