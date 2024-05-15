package medicalrecord

import "context"

type MedicalRecordService interface {
	RegisterPatient(
		ctx context.Context,
		req RegisterPatientReq,
		res *RegisterPatientRes,
	) error
}
