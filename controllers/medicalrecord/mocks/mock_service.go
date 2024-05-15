// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/JesseNicholas00/HaloSuster/services/medicalrecord (interfaces: MedicalRecordService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	medicalrecord "github.com/JesseNicholas00/HaloSuster/services/medicalrecord"
	gomock "github.com/golang/mock/gomock"
)

// MockMedicalRecordService is a mock of MedicalRecordService interface.
type MockMedicalRecordService struct {
	ctrl     *gomock.Controller
	recorder *MockMedicalRecordServiceMockRecorder
}

// MockMedicalRecordServiceMockRecorder is the mock recorder for MockMedicalRecordService.
type MockMedicalRecordServiceMockRecorder struct {
	mock *MockMedicalRecordService
}

// NewMockMedicalRecordService creates a new mock instance.
func NewMockMedicalRecordService(ctrl *gomock.Controller) *MockMedicalRecordService {
	mock := &MockMedicalRecordService{ctrl: ctrl}
	mock.recorder = &MockMedicalRecordServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMedicalRecordService) EXPECT() *MockMedicalRecordServiceMockRecorder {
	return m.recorder
}

// CreatePatient mocks base method.
func (m *MockMedicalRecordService) CreatePatient(arg0 context.Context, arg1 medicalrecord.CreatePatientReq, arg2 *medicalrecord.CreatePatientRes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePatient", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePatient indicates an expected call of CreatePatient.
func (mr *MockMedicalRecordServiceMockRecorder) CreatePatient(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePatient", reflect.TypeOf((*MockMedicalRecordService)(nil).CreatePatient), arg0, arg1, arg2)
}
