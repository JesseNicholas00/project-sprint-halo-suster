package auth

import (
	"github.com/JesseNicholas00/HaloSuster/types/nip"
	"github.com/labstack/echo/v4"
)

func (ctrl *authController) loginNurse(c echo.Context) error {
	return ctrl.loginShared(c, nip.RoleNurse)
}
