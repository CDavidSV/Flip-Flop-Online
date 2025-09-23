package middleware

import (
	"strings"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Authorize(jwks keyfunc.Keyfunc) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			prefix := "Bearer "
			headers := c.Request().Header.Get("Authorization")

			// If the header is empty or does not start with "Bearer "
			if headers == "" || !strings.HasPrefix(headers, prefix) {
				return echo.ErrUnauthorized
			}

			token := strings.TrimPrefix(headers, prefix)
			if token == "" {
				return echo.ErrUnauthorized
			}

			// Parse the token
			parsedToken, err := jwt.Parse(token, jwks.Keyfunc)
			if err != nil || !parsedToken.Valid {
				return echo.ErrUnauthorized
			}

			claims := parsedToken.Claims.(jwt.MapClaims)

			// Store user information from token into context
			c.Set("user_id", claims["sub"])
			c.Set("role", claims["role"])

			return next(c)
		}
	}
}
