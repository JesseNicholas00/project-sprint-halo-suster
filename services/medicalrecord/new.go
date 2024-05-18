package medicalrecord

import (
	"github.com/JesseNicholas00/HaloSuster/repos/auth"
	"github.com/JesseNicholas00/HaloSuster/repos/medicalrecord"
	"github.com/JesseNicholas00/HaloSuster/utils/ctxrizz"
)

type medicalRecordServiceImpl struct {
	repo     medicalrecord.MedicalRecordRepository
	repoAuth auth.AuthRepository
	dbRizzer ctxrizz.DbContextRizzer
}

func NewMedicalRecordService(
	repo medicalrecord.MedicalRecordRepository,
	repoAuth auth.AuthRepository,
	dbRizzer ctxrizz.DbContextRizzer,
) MedicalRecordService {
	return &medicalRecordServiceImpl{
		repo:     repo,
		repoAuth: repoAuth,
		dbRizzer: dbRizzer,
	}
}
