package medicalrecord

import (
	"errors"
	"net/http"

	"github.com/JesseNicholas00/HaloSuster/services/auth"
	"github.com/JesseNicholas00/HaloSuster/services/medicalrecord"
	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/JesseNicholas00/HaloSuster/utils/request"
	"github.com/labstack/echo/v4"
)

func (ctrl *medicalRecordController) createRecord(c echo.Context) error {
	var req medicalrecord.CreateRecordReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	req.CreatedById = c.
		Get("session").(auth.GetSessionFromTokenRes).
		UserId

	if err := ctrl.service.CreateRecord(
		c.Request().Context(),
		req,
		&medicalrecord.CreateRecordRes{},
	); err != nil {
		if errors.Is(err, medicalrecord.ErrIdentityNumberNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, echo.Map{
				"message": "patient identity number not found",
			})
		}

		return errorutil.AddCurrentContext(err)
	}

	return c.NoContent(http.StatusCreated)
}
