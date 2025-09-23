package main

import (
	"fmt"
	"os"

	"github.com/CDavidSV/Flip-Flop-Online/backend/api/handler"
	m "github.com/CDavidSV/Flip-Flop-Online/backend/api/middleware"
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

	apiV1 := e.Group("/api/v1")

	// Social
	socialGroup := apiV1.Group("/social")
	socialGroup.POST("/friendship", h.PostRequestFriendship)
	socialGroup.PUT("/friendship/:user_id", h.PutAcceptFriendship)
	socialGroup.DELETE("/friendship/:user_id", h.DeleteRemoveFriend)
	socialGroup.GET("/friends", h.GetListFriends)
	socialGroup.GET("/friend-requests", h.GetListFriendRequests)
	socialGroup.GET("/friends/search", h.GetSearchFriends)
	socialGroup.POST("/block/:user_id", h.PostBlockUser)
	socialGroup.DELETE("/block/:user_id", h.DeleteUnblockUser)
	socialGroup.GET("/blocks", h.GetListBlockedUsers)
	socialGroup.GET("/user/:user_id/profile", h.GetUserProfile)
	socialGroup.GET("/user/:user_id/statistics", h.GetUserStatistics)
	socialGroup.GET("/user/:user_id/game-history", h.GetUserGameHistory)

	// Profile
	profileGroup := apiV1.Group("/profile")
	profileGroup.GET("", h.GetProfile)
	profileGroup.PUT("", h.PutUpdateProfile)
	profileGroup.POST("/avatar", h.PostUploadAvatar)
	profileGroup.GET("/game-history", h.GetGameHistory)
	profileGroup.GET("/statistics", h.GetStatistics)

	// Game
	gameGroup := apiV1.Group("/game")
	gameGroup.POST("", h.PostCreateGame)
	gameGroup.POST("/invite", h.PostInviteUser)
	gameGroup.PUT("/invite/:invite_id", h.PostAcceptInvite)
	gameGroup.GET("/state/:game_id", h.GetGameState)
	gameGroup.GET("/is-in-game", h.GetIsUserInGame)
	gameGroup.POST("/join-request", h.PostRequestToJoin)
	gameGroup.PUT("/join-request/:user_id", h.PostAcceptJoinRequest)
}

func main() {
	e := echo.New()
	e.HideBanner = true

	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level}")
	}
	e.Logger.SetLevel(log.INFO)

	fmt.Println(config.Banner + "\nFlip-Flop API v" + config.Version)

	// Middleware
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{LogLevel: log.ERROR}))
	e.Use(m.Logger)
	e.Use(middleware.CORSWithConfig(config.CorsConfig))

	// Connect to Database
	e.Logger.Info("Connecting to database...")
	dbPool, err := data.NewPostgresPool(config.DSN)
	if err != nil {
		e.Logger.Error("Unnable to connect to database", "error", err)
		os.Exit(1)
	}
	defer dbPool.Close()
	e.Logger.Info("Database connection established")

	// Run migrations
	// migrations.RunMigrations()

	// Load Routes
	loadRoutes(e, dbPool)

	if err := e.Start(config.ServerAddress); err != nil {
		e.Logger.Error(err.Error())
	}
}
