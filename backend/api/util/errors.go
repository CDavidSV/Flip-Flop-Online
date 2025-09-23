package util

import (
	"net/http"

	"github.com/CDavidSV/Flip-Flop-Online/backend/internal/types"
	"github.com/CDavidSV/Flip-Flop-Online/backend/internal/validator"
	"github.com/labstack/echo/v4"
)

type APIErrorCode string

const (
	ValidationError APIErrorCode = "validation_error"
	InvalidRequest  APIErrorCode = "invalid_request_payload"
	ServerError     APIErrorCode = "sinternal_erver_error"
)

func buildErrorJSON(errorCode APIErrorCode) types.JSONMap {
	return types.JSONMap{
		"error_code": errorCode,
	}
}

func ClientErrorResponse(c echo.Context, status int, errorCode APIErrorCode) error {
	return c.JSON(status, buildErrorJSON(errorCode))
}

func ServerErrorResponse(c echo.Context, message string) error {
	resJSON := buildErrorJSON(ServerError)
	resJSON["msg"] = message

	return c.JSON(http.StatusInternalServerError, resJSON)
}

func ValidationErrorResponse(c echo.Context, errors []validator.FieldError) error {
	return c.JSON(http.StatusBadRequest, types.JSONMap{"msg": "validation failed", "errors": errors})
}
