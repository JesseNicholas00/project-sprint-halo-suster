package medicalrecord

import (
	"context"
	"time"

	"github.com/JesseNicholas00/HaloSuster/repos/medicalrecord"
	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
)

func (svc *medicalRecordServiceImpl) ListPatients(
	ctx context.Context,
	req ListPatientsReq,
	res *ListPatientsRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	repoRes, err := svc.repo.ListPatients(ctx, medicalrecord.PatientFilter{
		IdentityNumber: req.IdentityNumber,
		Limit:          *req.Limit,
		Offset:         *req.Offset,
		Name:           req.Name,
		PhoneNumber:    req.PhoneNumber,
		CreatedAtSort:  req.CreatedAtSort,
	})
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	for _, patient := range repoRes {
		res.Data = append(res.Data, ListPatientsResData{
			IdentityNumber: patient.IdentityNumber,
			PhoneNumber:    patient.PhoneNumber,
			Name:           patient.Name,
			BirthDate:      patient.BirthDate.Format(time.DateOnly),
			Gender:         patient.Gender,
			CreatedAt:      patient.CreatedAt.Format(time.RFC3339Nano),
		})
	}

	return nil
}
