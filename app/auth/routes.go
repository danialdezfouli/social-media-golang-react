package auth

import (
	"github.com/labstack/echo/v4"
	"jupiter/app/common/middleware"
)

func Routes(e *echo.Echo) {
	restricted := e.Group("/auth")
	restricted.Use(middleware.AuthMiddleware(e))
	restricted.GET("/me", me)

	e.POST("/auth/login", login)
	e.POST("/auth/register", register)
}
