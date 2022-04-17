package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"jupiter/app/common/token"
	"jupiter/config"
)

func AuthMiddleware() echo.MiddlewareFunc {
	secret := config.GetConfig().Auth.AccessTokenSecret
	config := middleware.JWTConfig{
		Claims:     &token.JwtCustomClaims{},
		SigningKey: []byte(secret),
	}

	return middleware.JWTWithConfig(config)
}

func RefreshTokenValidationMiddleware() echo.MiddlewareFunc {
	secret := config.GetConfig().Auth.RefreshTokenSecret
	config := middleware.JWTConfig{
		Claims:      &token.JwtCustomClaims{},
		SigningKey:  []byte(secret),
		TokenLookup: "cookie:refresh_token",
	}

	return middleware.JWTWithConfig(config)
}
