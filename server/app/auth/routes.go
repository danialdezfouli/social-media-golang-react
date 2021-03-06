package auth

import (
	"github.com/labstack/echo/v4"
	"jupiter/app/common/middleware"
)

func Routes(e *echo.Echo) {
	e.POST("/auth/login", login)
	e.POST("/auth/register", register)
	e.GET("/auth/me", me, middleware.AuthMiddleware(), middleware.UserMiddleware())
	e.POST("/auth/refresh", refreshToken, middleware.RefreshTokenValidationMiddleware())
}
