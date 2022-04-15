package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"jupiter/app/common/token"
	"jupiter/config"
)

func AuthMiddleware(e *echo.Echo) echo.MiddlewareFunc {
	secret := config.GetConfig().Auth.AccessTokenSecret
	config := middleware.JWTConfig{
		Claims:     &token.JwtCustomClaims{},
		SigningKey: []byte(secret),
	}

	return middleware.JWTWithConfig(config)
}

// RefreshTokenValidationMiddleware TODO: fix it
func RefreshTokenValidationMiddleware(e *echo.Echo) echo.MiddlewareFunc {
	secret := config.GetConfig().Auth.RefreshTokenSecret
	config := middleware.JWTConfig{
		Claims:     &token.JwtCustomClaims{},
		SigningKey: []byte(secret),
	}

	return middleware.JWTWithConfig(config)
}
