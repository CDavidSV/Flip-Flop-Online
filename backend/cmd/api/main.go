package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/CDavidSV/Flip-Flop-Online/backend/api/handler"
	"github.com/CDavidSV/Flip-Flop-Online/backend/api/middlewares"
	"github.com/CDavidSV/Flip-Flop-Online/backend/config"
	"github.com/CDavidSV/Flip-Flop-Online/backend/internal/data"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func loadRoutes(e *echo.Echo, db *pgxpool.Pool) {
	h := handler.NewHandler(db)

	e.GET("/health", h.HealthCheck)

	// Auth routes
	authGroup := e.Group("/auth")
	authGroup.POST("/signup", h.Signup)
}

func main() {
	e := echo.New()
	e.HideBanner = true

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	fmt.Println(config.Banner + "\nFlip-Flop API v" + config.Version)

	// Middleware
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{LogLevel: log.ERROR}))
	e.Use(middlewares.Logger)
	e.Use(middleware.CORSWithConfig(config.CorsConfig))

	// Connect to Database
	logger.Info("Connecting to database...")
	dbPool, err := data.NewPostgresPool(config.DSN)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer dbPool.Close()
	logger.Info("Database connection established")

	// Load Routes
	loadRoutes(e, dbPool)

	if err := e.Start(config.ServerAddress); err != nil {
		logger.Error(err.Error())
	}
}
