package medicalrecord

import (
	"github.com/JesseNicholas00/HaloSuster/repos/medicalrecord"
	"github.com/JesseNicholas00/HaloSuster/utils/ctxrizz"
)

type medicalRecordServiceImpl struct {
	repo     medicalrecord.MedicalRecordRepository
	dbRizzer ctxrizz.DbContextRizzer
}

func NewMedicalRecordService(
	repo medicalrecord.MedicalRecordRepository,
	dbRizzer ctxrizz.DbContextRizzer,
) MedicalRecordService {
	return &medicalRecordServiceImpl{
		repo:     repo,
		dbRizzer: dbRizzer,
	}
}
