package medicalrecord

import (
	"github.com/JesseNicholas00/HaloSuster/controllers"
	"github.com/JesseNicholas00/HaloSuster/middlewares"
	"github.com/JesseNicholas00/HaloSuster/services/medicalrecord"
	"github.com/labstack/echo/v4"
)

type medicalRecordController struct {
	service medicalrecord.MedicalRecordService
	authMw  middlewares.Middleware
}

func (ctrl *medicalRecordController) Register(server *echo.Echo) error {
	g := server.Group("/v1/medical", ctrl.authMw.Process)

	g.POST("/patient", ctrl.registerPatient)
	g.GET("/patient", ctrl.listPatients)
	g.POST("/record", ctrl.createRecord)
	g.GET("/record", ctrl.listRecord)

	return nil
}

func NewMedicalRecordController(
	service medicalrecord.MedicalRecordService,
	authMw middlewares.Middleware,
) controllers.Controller {
	return &medicalRecordController{
		service: service,
		authMw:  authMw,
	}
}
