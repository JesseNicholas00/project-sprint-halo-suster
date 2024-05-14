package auth

import (
	"github.com/JesseNicholas00/HaloSuster/types/nip"
	"github.com/labstack/echo/v4"
)

func (ctrl *authController) loginIt(c echo.Context) error {
	return ctrl.loginShared(c, nip.RoleIt)
}
