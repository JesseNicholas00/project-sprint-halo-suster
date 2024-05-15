package medicalrecord

import "context"

type MedicalRecordService interface {
	CreatePatient(
		ctx context.Context,
		req CreatePatientReq,
		res *CreatePatientRes,
	) error
}
