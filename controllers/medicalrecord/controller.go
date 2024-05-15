package medicalrecord

import (
	"github.com/JesseNicholas00/HaloSuster/controllers"
	"github.com/JesseNicholas00/HaloSuster/services/medicalrecord"
	"github.com/labstack/echo/v4"
)

type medicalRecordController struct {
	service medicalrecord.MedicalRecordService
}

func (ctrl *medicalRecordController) Register(server *echo.Echo) error {
	g := server.Group("/v1/medical")
	g.POST("/patient", ctrl.registerPatient)
	return nil
}

func NewMedicalRecordController(
	service medicalrecord.MedicalRecordService,
) controllers.Controller {
	return &medicalRecordController{
		service: service,
	}
}
