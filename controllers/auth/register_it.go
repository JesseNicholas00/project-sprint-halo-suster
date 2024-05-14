package auth

import (
	"errors"
	"net/http"

	"github.com/JesseNicholas00/HaloSuster/services/auth"
	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/JesseNicholas00/HaloSuster/utils/request"
	"github.com/labstack/echo/v4"
)

func (ctrl *authController) registerIt(c echo.Context) error {
	var req auth.RegisterItReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	var res auth.RegisterItRes
	err := ctrl.service.RegisterIt(c.Request().Context(), req, &res)
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrNipAlreadyExists):
			return echo.NewHTTPError(http.StatusConflict)

		default:
			return errorutil.AddCurrentContext(err)
		}
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "User registered successfully",
		"data":    res,
	})
}
