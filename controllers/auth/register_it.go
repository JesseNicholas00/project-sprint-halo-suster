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

func (ctrl *authController) registerIt(c echo.Context) error {
	var req auth.RegisterItReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	if nip.GetRole(req.Nip) != nip.RoleIt {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "wrong nip role",
		})
	}

	var res auth.RegisterItRes
	err := ctrl.service.RegisterIt(c.Request().Context(), req, &res)
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
