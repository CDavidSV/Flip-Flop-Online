package handler

import (
	"github.com/CDavidSV/Flip-Flop-Online/backend/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	db *pgxpool.Pool
}

func NewHandler(db *pgxpool.Pool) *Handler {
	return &Handler{db: db}
}

func (h *Handler) HealthCheck(c echo.Context) error {
	return c.JSON(200, map[string]string{"status": "ok", "version": config.Version})
}
