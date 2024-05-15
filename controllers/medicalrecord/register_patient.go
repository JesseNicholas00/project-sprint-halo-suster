package medicalrecord

import (
	"errors"
	"net/http"

	"github.com/JesseNicholas00/HaloSuster/services/medicalrecord"
	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/JesseNicholas00/HaloSuster/utils/request"
	"github.com/labstack/echo/v4"
)

func (s *medicalRecordController) registerPatient(c echo.Context) error {
	var req medicalrecord.RegisterPatientReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	var res medicalrecord.RegisterPatientRes
	err := s.service.RegisterPatient(c.Request().Context(), req, &res)
	if err != nil {
		switch {
		case errors.Is(err, medicalrecord.ErrDuplicateIdentityNumber):
			return echo.NewHTTPError(http.StatusConflict)

		default:
			return errorutil.AddCurrentContext(err)
		}
	}

	return c.NoContent(http.StatusCreated)
}
