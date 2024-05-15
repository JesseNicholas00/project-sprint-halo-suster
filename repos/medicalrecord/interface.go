package medicalrecord

import "context"

type MedicalRecordRepository interface {
	CreatePatient(ctx context.Context, patient Patient) error
}
