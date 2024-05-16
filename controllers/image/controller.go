package image

import (
	"github.com/JesseNicholas00/HaloSuster/controllers"
	"github.com/JesseNicholas00/HaloSuster/middlewares"
	"github.com/JesseNicholas00/HaloSuster/services/image"
	"github.com/labstack/echo/v4"
)

type imageController struct {
	service image.ImageService
	authMw  middlewares.Middleware
}

func (ctrl *imageController) Register(server *echo.Echo) error {
	g := server.Group("/v1/image")

	g.POST("", ctrl.uploadImage, ctrl.authMw.Process)

	return nil
}

func NewImageController(
	service image.ImageService,
	authMw middlewares.Middleware,
) controllers.Controller {
	return &imageController{
		service: service,
		authMw:  authMw,
	}
}
