package feeds

import "github.com/labstack/echo/v4"

func Routes(e *echo.Echo) {
	e.GET("/feeds", home)
}
