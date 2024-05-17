package medicalrecord

import (
	"testing"

	authmocks "github.com/JesseNicholas00/HaloSuster/services/auth/mocks"
	"github.com/JesseNicholas00/HaloSuster/services/medicalrecord/mocks"
	"github.com/JesseNicholas00/HaloSuster/utils/ctxrizz"
	"github.com/golang/mock/gomock"
)

//go:generate mockgen -destination mocks/mock_repo.go -package mocks github.com/JesseNicholas00/HaloSuster/repos/medicalrecord MedicalRecordRepository

func NewWithMockedRepo(
	t *testing.T,
) (
	mockCtrl *gomock.Controller,
	service *medicalRecordServiceImpl,
	mockedRepo *mocks.MockMedicalRecordRepository,
	mockedAuthRepo *authmocks.MockAuthRepository,
) {
	mockCtrl = gomock.NewController(t)
	mockedRepo = mocks.NewMockMedicalRecordRepository(mockCtrl)
	mockedAuthRepo = authmocks.NewMockAuthRepository(mockCtrl)
	service = NewMedicalRecordService(
		mockedRepo,
		mockedAuthRepo,
		ctxrizz.NewDbContextNoopRizzer(),
	).(*medicalRecordServiceImpl)
	return
}
