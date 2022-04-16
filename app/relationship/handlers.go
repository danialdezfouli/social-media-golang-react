package relationship

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func follow(c echo.Context) error {
	return c.String(http.StatusOK, "Follow")
}

func unfollow(c echo.Context) error {
	return c.String(http.StatusOK, "Unfollow")
}
