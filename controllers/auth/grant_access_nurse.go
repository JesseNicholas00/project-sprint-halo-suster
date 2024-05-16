package auth

import (
	"errors"
	"net/http"

	"github.com/JesseNicholas00/HaloSuster/services/auth"
	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/JesseNicholas00/HaloSuster/utils/request"
	"github.com/labstack/echo/v4"
)

func (ctrl *authController) grantAccessNurse(c echo.Context) error {
	var req auth.GrantAccessNurseReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	var res auth.GrantAccessNurseRes
	err := ctrl.service.GrantAccessNurse(c.Request().Context(), req, &res)
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrUserNotFound):
			return echo.NewHTTPError(http.StatusNotFound, echo.Map{
				"message": "user not found",
			})

		default:
			return errorutil.AddCurrentContext(err)
		}
	}

	return c.NoContent(http.StatusOK)
}
