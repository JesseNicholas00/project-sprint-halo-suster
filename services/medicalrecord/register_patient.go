package medicalrecord

import (
	"context"
	"errors"

	"github.com/JesseNicholas00/HaloSuster/repos/medicalrecord"
	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/JesseNicholas00/HaloSuster/utils/helper"
)

func (svc *medicalRecordServiceImpl) RegisterPatient(
	ctx context.Context,
	req RegisterPatientReq,
	res *RegisterPatientRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	if err := svc.repo.CreatePatient(ctx, medicalrecord.Patient{
		IdentityNumber: req.IdentityNumber,
		PhoneNumber:    req.PhoneNumber,
		Name:           req.Name,
		BirthDate:      helper.MustParse(req.BirthDate),
		Gender:         req.Gender,
		ImageUrl:       req.IdentityCardImg,
	}); err != nil {
		switch {
		case errors.Is(err, medicalrecord.ErrDuplicateIdentityNumber):
			return ErrDuplicateIdentityNumber

		default:
			return errorutil.AddCurrentContext(err)
		}
	}

	return nil
}
