package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %s", err)
	}

	DSN = os.Getenv("DATABASE_URI")
	AccessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
	RefreshTokenSecret = os.Getenv("REFRESH_TOKEN_SECRET")
	ServerAddress = os.Getenv("SERVER_ADDRESS")
	SupabaseURL = os.Getenv("SUPABASE_URL")
	SupabaseKey = os.Getenv("SUPABASE_KEY")
}

var (
	Banner string = `
    _________             ________
   / ____/ (_)___        / ____/ /___  ____
  / /_  / / / __ \______/ /_  / / __ \/ __ \
 / __/ / / / /_/ /_____/ __/ / / /_/ / /_/ /
/_/   /_/_/ .___/     /_/   /_/\____/ .___/
         /_/                       /_/
`

	Version            string = "0.0.1"
	DSN                string
	AccessTokenSecret  string
	RefreshTokenSecret string
	ServerAddress      string
	SupabaseURL        string
	SupabaseKey        string
	CorsConfig         = middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "HEAD", "OPTION", "PUT"},
		AllowHeaders:     []string{"User-Agent", "Content-Type", "Accept", "Accept-Encoding", "Accept-Language", "Cache-Control", "Connection", "DNT", "Host", "Origin", "Pragma", "Referer", "Cookie"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}
)
