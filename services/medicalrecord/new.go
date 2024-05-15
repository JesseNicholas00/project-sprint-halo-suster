package medicalrecord

import (
	"github.com/JesseNicholas00/HaloSuster/repos/medicalrecord"
)

type medicalRecordServiceImpl struct {
	repo medicalrecord.MedicalRecordRepository
}

func NewMedicalRecordService(
	repo medicalrecord.MedicalRecordRepository,
) MedicalRecordService {
	return &medicalRecordServiceImpl{
		repo: repo,
	}
}
