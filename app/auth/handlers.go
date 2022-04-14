package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"jupiter/app/auth/dto"
	"jupiter/app/auth/service"
	"jupiter/app/common/token"
	"net/http"
)

func me(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*token.JwtCustomClaims)
	name := claims.Username

	fmt.Println(claims)

	return c.String(http.StatusOK, "Welcome "+name+"!")

}

func login(c echo.Context) error {
	loginService := service.LoginService{}
	authService := service.AuthService{}
	input := new(dto.LoginInput)

	if err := c.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(input); err != nil {
		return err
	}

	user, attemptError := loginService.Attempt(input)
	if attemptError != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, attemptError.Error())
	}

	accessToken, atErr := authService.GenerateAccessToken(user)
	if atErr != nil {
		return echo.ErrInternalServerError
	}

	refreshToken, rtErr := authService.GenerateRefreshToken(user)
	if rtErr != nil {
		return echo.ErrInternalServerError
	}

	// TODO: save refreshToken to cookie

	return c.JSON(http.StatusOK, loginResponse{
		AccessToken:  accessToken.String(),
		RefreshToken: refreshToken.String(),
	})
}

func register(c echo.Context) error {
	return c.String(http.StatusOK, "Register!")
}
