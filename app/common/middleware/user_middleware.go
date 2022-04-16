package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"jupiter/app"
	"jupiter/app/common/token"
	"jupiter/app/model"
	"net/http"
)

func UserMiddleware() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			jwtUser := c.Get("user").(*jwt.Token)
			claims := jwtUser.Claims.(*token.JwtCustomClaims)

			user := &model.User{
				ID: claims.UserID,
			}

			result := app.GetDB().First(user)

			if result.Error != nil {
				return echo.ErrUnauthorized
			}

			if user.Suspended {
				return echo.NewHTTPError(http.StatusUnauthorized, "your account is suspended")
			}

			c.Set("user", user)

			return next(c)

		}
	}
}
