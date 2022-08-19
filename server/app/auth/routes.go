package auth

import (
	"jupiter/app/common/middleware"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.POST("/auth/login", login)
	e.POST("/auth/register", register)
	e.POST("/auth/refresh", refreshToken, middleware.RefreshTokenValidationMiddleware())
	e.GET("/auth/me", me, middleware.AuthMiddleware(), middleware.UserMiddleware())
	e.POST("/auth/profile", updateProfile, middleware.AuthMiddleware(), middleware.UserMiddleware())
}
