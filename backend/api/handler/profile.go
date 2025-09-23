package handler

import (
	"github.com/CDavidSV/Flip-Flop-Online/backend/api/util"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetProfile(c echo.Context) error {
	return util.ServerErrorResponse(c, "Not implemented")
}

func (h *Handler) PutUpdateProfile(c echo.Context) error {
	return util.ServerErrorResponse(c, "Not implemented")
}

func (h *Handler) PostUploadAvatar(c echo.Context) error {
	return util.ServerErrorResponse(c, "Not implemented")
}

func (h *Handler) GetGameHistory(c echo.Context) error {
	return util.ServerErrorResponse(c, "Not implemented")
}

func (h *Handler) GetStatistics(c echo.Context) error {
	return util.ServerErrorResponse(c, "Not implemented")
}
