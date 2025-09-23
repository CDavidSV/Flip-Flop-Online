package handler

import (
	"github.com/CDavidSV/Flip-Flop-Online/backend/api/util"
	"github.com/labstack/echo/v4"
)

func (h *Handler) PostCreateGame(c echo.Context) error {
	return util.ServerErrorResponse(c, "Not implemented")
}

func (h *Handler) PostInviteUser(c echo.Context) error {
	return util.ServerErrorResponse(c, "Not implemented")
}

func (h *Handler) PostAcceptInvite(c echo.Context) error {
	return util.ServerErrorResponse(c, "Not implemented")
}

func (h *Handler) GetGameState(c echo.Context) error {
	return util.ServerErrorResponse(c, "Not Implemented")
}

func (h *Handler) GetIsUserInGame(c echo.Context) error {
	return util.ServerErrorResponse(c, "Not Implemented")
}

func (h *Handler) PostRequestToJoin(c echo.Context) error {
	return util.ServerErrorResponse(c, "Not implemented")
}

func (h *Handler) PostAcceptJoinRequest(c echo.Context) error {
	return util.ServerErrorResponse(c, "Not implemented")
}
