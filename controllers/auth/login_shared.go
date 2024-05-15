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

func (ctrl *authController) loginShared(
	c echo.Context,
	expectedRole nip.NipRole,
) error {
	var req auth.LoginReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	if nip.GetRole(req.Nip) != expectedRole {
		return echo.NewHTTPError(http.StatusNotFound, echo.Map{
			"message": "wrong nip role",
		})
	}

	var res auth.LoginRes
	err := ctrl.service.Login(c.Request().Context(), req, &res)
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrUserNotFound):
			return echo.NewHTTPError(http.StatusNotFound, echo.Map{
				"message": "no such user found",
			})

		case errors.Is(err, auth.ErrInvalidCredentials):
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
				"message": "invalid credentials",
			})

		default:
			return errorutil.AddCurrentContext(err)
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "User logged in successfully",
		"data":    res,
	})
}
