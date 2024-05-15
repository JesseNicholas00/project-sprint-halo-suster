package medicalrecord

import (
	"testing"

	"github.com/JesseNicholas00/HaloSuster/services/medicalrecord/mocks"
	"github.com/golang/mock/gomock"
)

//go:generate mockgen -destination mocks/mock_repo.go -package mocks github.com/JesseNicholas00/HaloSuster/repos/medicalrecord MedicalRecordRepository

func NewWithMockedRepo(
	t *testing.T,
) (
	mockCtrl *gomock.Controller,
	service *medicalRecordServiceImpl,
	mockedRepo *mocks.MockMedicalRecordRepository,
) {
	mockCtrl = gomock.NewController(t)
	mockedRepo = mocks.NewMockMedicalRecordRepository(mockCtrl)
	service = NewMedicalRecordService(mockedRepo).(*medicalRecordServiceImpl)
	return
}
