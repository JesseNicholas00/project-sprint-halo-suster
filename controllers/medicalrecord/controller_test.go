package medicalrecord

import (
	"testing"

	"github.com/JesseNicholas00/HaloSuster/controllers/medicalrecord/mocks"
	"github.com/JesseNicholas00/HaloSuster/middlewares"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	. "github.com/smartystreets/goconvey/convey"
)

//go:generate mockgen -destination mocks/mock_service.go -package mocks github.com/JesseNicholas00/HaloSuster/services/medicalrecord MedicalRecordService

func NewControllerWithMockedService(
	t *testing.T,
) (
	mockCtrl *gomock.Controller,
	controller *medicalRecordController,
	mockedService *mocks.MockMedicalRecordService,
) {
	mockCtrl = gomock.NewController(t)
	mockedService = mocks.NewMockMedicalRecordService(mockCtrl)
	controller = NewMedicalRecordController(
		mockedService,
		middlewares.NewNoopMiddleware(),
	).(*medicalRecordController)
	return
}

func TestRegister(t *testing.T) {
	mockCtrl, controller, _ := NewControllerWithMockedService(t)
	defer mockCtrl.Finish()

	Convey("When registering methods with an echo instance", t, func() {
		e := echo.New()
		err := controller.Register(e)
		Convey("Should not return error", func() {
			So(err, ShouldBeNil)
		})
	})
}
