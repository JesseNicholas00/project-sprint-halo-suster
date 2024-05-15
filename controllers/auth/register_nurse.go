package auth

import (
	"errors"
	"net/http"

	"github.com/JesseNicholas00/HaloSuster/services/auth"
	"github.com/JesseNicholas00/HaloSuster/types/nip"
	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/JesseNicholas00/HaloSuster/utils/request"
	"github.com/labstack/echo/v4"
)

func (ctrl *authController) registerNurse(c echo.Context) error {
	var req auth.RegisterNurseReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	if nip.GetRole(req.Nip) != nip.RoleNurse {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "wrong nip role",
		})
	}

	var res auth.RegisterNurseRes
	err := ctrl.service.RegisterNurse(c.Request().Context(), req, &res)
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrNipAlreadyExists):
			return echo.NewHTTPError(http.StatusConflict, echo.Map{
				"message": "nip already registered",
			})

		default:
			return errorutil.AddCurrentContext(err)
		}
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "User registered successfully",
		"data":    res,
	})
}
