package auth

import (
	"net/http"

	"github.com/JesseNicholas00/HaloSuster/services/auth"
	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/JesseNicholas00/HaloSuster/utils/helper"
	"github.com/JesseNicholas00/HaloSuster/utils/request"
	"github.com/labstack/echo/v4"
)

func (ctrl *authController) listUsers(c echo.Context) error {
	var req auth.ListUsersReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	if req.CreatedAtSort != nil {
		if *req.CreatedAtSort != "asc" && *req.CreatedAtSort != "desc" {
			req.CreatedAtSort = nil
		}
	}
	if req.Limit == nil {
		req.Limit = helper.ToPointer(5)
	}
	if req.Offset == nil {
		req.Offset = helper.ToPointer(0)
	}

	var res auth.ListUsersRes
	if err := ctrl.service.ListUsers(
		c.Request().Context(),
		req,
		&res,
	); err != nil {
		return errorutil.AddCurrentContext(err)
	}

	// nil to empty array
	if res.Data == nil {
		res.Data = make([]auth.ListUsersResData, 0)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    res.Data,
	})
}
