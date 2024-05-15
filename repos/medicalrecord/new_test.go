package medicalrecord_test

import (
	"testing"

	"github.com/JesseNicholas00/HaloSuster/repos/medicalrecord"
	"github.com/JesseNicholas00/HaloSuster/utils/ctxrizz"
	"github.com/JesseNicholas00/HaloSuster/utils/unittesting"
)

func NewWithTestDatabase(t *testing.T) medicalrecord.MedicalRecordRepository {
	db := unittesting.SetupTestDatabase("../../migrations", t)
	return medicalrecord.NewMedicalRecordRepository(
		ctxrizz.NewDbContextRizzer(db),
	)
}
