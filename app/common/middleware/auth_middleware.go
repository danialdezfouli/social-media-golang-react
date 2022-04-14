package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"jupiter/app/common/token"
	"jupiter/config"
)

func AuthMiddleware(e *echo.Echo) echo.MiddlewareFunc {
	secret := config.GetConfig().Auth.AccessTokenSecret
	config := middleware.JWTConfig{
		Claims:     &token.JwtCustomClaims{},
		SigningKey: []byte(secret),
		KeyFunc:    getKey,
	}

	return middleware.JWTWithConfig(config)
}

func getKey(token *jwt.Token) (interface{}, error) {

	fmt.Println(token)

	return token, nil
}
