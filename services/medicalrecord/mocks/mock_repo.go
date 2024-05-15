// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/JesseNicholas00/HaloSuster/repos/medicalrecord (interfaces: MedicalRecordRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	medicalrecord "github.com/JesseNicholas00/HaloSuster/repos/medicalrecord"
	gomock "github.com/golang/mock/gomock"
)

// MockMedicalRecordRepository is a mock of MedicalRecordRepository interface.
type MockMedicalRecordRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMedicalRecordRepositoryMockRecorder
}

// MockMedicalRecordRepositoryMockRecorder is the mock recorder for MockMedicalRecordRepository.
type MockMedicalRecordRepositoryMockRecorder struct {
	mock *MockMedicalRecordRepository
}

// NewMockMedicalRecordRepository creates a new mock instance.
func NewMockMedicalRecordRepository(ctrl *gomock.Controller) *MockMedicalRecordRepository {
	mock := &MockMedicalRecordRepository{ctrl: ctrl}
	mock.recorder = &MockMedicalRecordRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMedicalRecordRepository) EXPECT() *MockMedicalRecordRepositoryMockRecorder {
	return m.recorder
}

// CreatePatient mocks base method.
func (m *MockMedicalRecordRepository) CreatePatient(arg0 context.Context, arg1 medicalrecord.Patient) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePatient", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePatient indicates an expected call of CreatePatient.
func (mr *MockMedicalRecordRepositoryMockRecorder) CreatePatient(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePatient", reflect.TypeOf((*MockMedicalRecordRepository)(nil).CreatePatient), arg0, arg1)
}
