package feeds

import (
	"github.com/labstack/echo/v4"
	"jupiter/app/common/middleware"
)

func Routes(e *echo.Echo) {
	r := e.Group("/")
	r.Use(middleware.AuthMiddleware())

	r.GET("feeds", feeds)
}
