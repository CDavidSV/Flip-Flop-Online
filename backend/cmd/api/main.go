package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{LogLevel: log.ERROR}))
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.Logger.Fatal(e.Start(":8000"))
}
