package auth

import (
	"github.com/JesseNicholas00/HaloSuster/controllers"
	"github.com/JesseNicholas00/HaloSuster/services/auth"
	"github.com/labstack/echo/v4"
)

type authController struct {
	service auth.AuthService
}

func (s *authController) Register(server *echo.Echo) error {
	g := server.Group("/v1/user")

	g.POST("/it/register", s.registerIt)
	g.POST("/it/login", s.loginIt)

	g.POST("/nurse/login", s.loginNurse)

	return nil
}

func NewAuthController(
	service auth.AuthService,
) controllers.Controller {
	return &authController{
		service: service,
	}
}
