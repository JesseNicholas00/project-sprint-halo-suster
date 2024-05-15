package medicalrecord

import (
	"github.com/JesseNicholas00/HaloSuster/utils/ctxrizz"
)

type medicalRecordRepositoryImpl struct {
	dbRizzer ctxrizz.DbContextRizzer
}

func NewMedicalRecordRepository(
	dbRizzer ctxrizz.DbContextRizzer,
) MedicalRecordRepository {
	return &medicalRecordRepositoryImpl{
		dbRizzer: dbRizzer,
	}
}
