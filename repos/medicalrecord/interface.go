package medicalrecord

import "context"

type MedicalRecordRepository interface {
	CreatePatient(ctx context.Context, patient Patient) error
	ListPatients(ctx context.Context, filter PatientFilter) ([]Patient, error)
	CreateRecord(ctx context.Context, record Record) error
	ListRecord(ctx context.Context, filter RecordFilter) ([]Record, error)
}
