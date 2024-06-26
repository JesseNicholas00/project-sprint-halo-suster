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

// CreateRecord mocks base method.
func (m *MockMedicalRecordService) CreateRecord(arg0 context.Context, arg1 medicalrecord.CreateRecordReq, arg2 *medicalrecord.CreateRecordRes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRecord", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRecord indicates an expected call of CreateRecord.
func (mr *MockMedicalRecordServiceMockRecorder) CreateRecord(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRecord", reflect.TypeOf((*MockMedicalRecordService)(nil).CreateRecord), arg0, arg1, arg2)
}

// ListPatients mocks base method.
func (m *MockMedicalRecordService) ListPatients(arg0 context.Context, arg1 medicalrecord.ListPatientsReq, arg2 *medicalrecord.ListPatientsRes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPatients", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListPatients indicates an expected call of ListPatients.
func (mr *MockMedicalRecordServiceMockRecorder) ListPatients(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPatients", reflect.TypeOf((*MockMedicalRecordService)(nil).ListPatients), arg0, arg1, arg2)
}

// ListRecord mocks base method.
func (m *MockMedicalRecordService) ListRecord(arg0 context.Context, arg1 medicalrecord.ListRecordReq, arg2 *medicalrecord.ListRecordRes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRecord", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListRecord indicates an expected call of ListRecord.
func (mr *MockMedicalRecordServiceMockRecorder) ListRecord(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRecord", reflect.TypeOf((*MockMedicalRecordService)(nil).ListRecord), arg0, arg1, arg2)
}

// RegisterPatient mocks base method.
func (m *MockMedicalRecordService) RegisterPatient(arg0 context.Context, arg1 medicalrecord.RegisterPatientReq, arg2 *medicalrecord.RegisterPatientRes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterPatient", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterPatient indicates an expected call of RegisterPatient.
func (mr *MockMedicalRecordServiceMockRecorder) RegisterPatient(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterPatient", reflect.TypeOf((*MockMedicalRecordService)(nil).RegisterPatient), arg0, arg1, arg2)
}
