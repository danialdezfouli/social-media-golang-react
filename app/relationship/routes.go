package relationship

import (
	"github.com/labstack/echo/v4"
	"jupiter/app/common/middleware"
)

func Routes(e *echo.Echo) {
	r := e.Group("/")
	r.Use(middleware.AuthMiddleware())

	r.POST("/follow/:id", follow)
	r.DELETE("/unfollow/:id", unfollow)
}
