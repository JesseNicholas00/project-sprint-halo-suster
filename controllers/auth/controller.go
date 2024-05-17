package auth

import (
	"github.com/JesseNicholas00/HaloSuster/controllers"
	"github.com/JesseNicholas00/HaloSuster/middlewares"
	"github.com/JesseNicholas00/HaloSuster/services/auth"
	"github.com/labstack/echo/v4"
)

type authController struct {
	service auth.AuthService
	authMw  middlewares.Middleware
}

func (ctrl *authController) Register(server *echo.Echo) error {
	g := server.Group("/v1/user")

	g.GET("", ctrl.listUsers, ctrl.authMw.Process)

	g.POST("/it/register", ctrl.registerIt)
	g.POST("/it/login", ctrl.loginIt)

	g.POST("/nurse/login", ctrl.loginNurse)
	g.POST("/nurse/register", ctrl.registerNurse, ctrl.authMw.Process)
	g.POST("/nurse/:userId/access", ctrl.grantAccessNurse, ctrl.authMw.Process)

	g.PUT("/nurse/:userId", ctrl.updateNurse, ctrl.authMw.Process)

	g.DELETE("/nurse/:userId", ctrl.deleteNurse, ctrl.authMw.Process)

	return nil
}

func NewAuthController(
	service auth.AuthService,
	authMw middlewares.Middleware,
) controllers.Controller {
	return &authController{
		service: service,
		authMw:  authMw,
	}
}
