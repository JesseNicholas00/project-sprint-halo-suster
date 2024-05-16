package medicalrecord

import (
	"net/http"

	"github.com/JesseNicholas00/HaloSuster/services/medicalrecord"
	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/JesseNicholas00/HaloSuster/utils/helper"
	"github.com/JesseNicholas00/HaloSuster/utils/request"
	"github.com/labstack/echo/v4"
)

func (ctrl *medicalRecordController) listPatients(c echo.Context) error {
	var req medicalrecord.ListPatientsReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	if req.CreatedAt != nil {
		if *req.CreatedAt != "asc" && *req.CreatedAt != "desc" {
			req.CreatedAt = nil
		}
	}
	if req.Limit == nil {
		req.Limit = helper.ToPointer(5)
	}
	if req.Offset == nil {
		req.Offset = helper.ToPointer(0)
	}

	var res medicalrecord.ListPatientsRes
	if err := ctrl.service.ListPatients(
		c.Request().Context(),
		req,
		&res,
	); err != nil {
		return errorutil.AddCurrentContext(err)
	}

	// nil to empty array
	if res.Data == nil {
		res.Data = make([]medicalrecord.ListPatientsResData, 0)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    res.Data,
	})
}
