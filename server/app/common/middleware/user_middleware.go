package middleware

import (
	"errors"
	"github.com/labstack/echo/v4"
	"jupiter/app/auth/service"
)

func UserMiddleware() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := service.AuthService{}
			user, err := auth.User(c)

			if err != nil {
				if errors.Is(err, &echo.HTTPError{}) {
					return err
				}

				return echo.ErrUnauthorized
			}

			c.Set("user", user)
			return next(c)
		}
	}
}
