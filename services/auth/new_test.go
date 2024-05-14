package auth

import (
	"testing"

	"github.com/JesseNicholas00/HaloSuster/services/auth/mocks"
	gomock "github.com/golang/mock/gomock"
)

//go:generate mockgen -destination mocks/mock_repo.go -package mocks github.com/JesseNicholas00/HaloSuster/repos/auth AuthRepository

func NewWithMockedRepo(
	t *testing.T,
) (
	mockCtrl *gomock.Controller,
	service *authServiceImpl,
	mockedRepo *mocks.MockAuthRepository,
) {
	mockCtrl = gomock.NewController(t)
	mockedRepo = mocks.NewMockAuthRepository(mockCtrl)
	service = NewAuthService(mockedRepo, "testKey", 8).(*authServiceImpl)
	return
}
