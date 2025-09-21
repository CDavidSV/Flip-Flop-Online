package handler

import (
	"github.com/CDavidSV/Flip-Flop-Online/backend/config"
	"github.com/CDavidSV/Flip-Flop-Online/backend/internal/data"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	queries *data.Queries
}

func NewHandler(db *pgxpool.Pool) *Handler {
	return &Handler{queries: data.NewQueries(db)}
}

func (h *Handler) HealthCheck(c echo.Context) error {
	return c.JSON(200, map[string]string{"status": "ok", "version": config.Version})
}
