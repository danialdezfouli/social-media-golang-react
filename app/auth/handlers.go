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

type loginResponse struct {
	AccessToken string `json:"accessToken"`
}
type meResponse struct {
	User struct {
		Name string
	}
}

func hello(c echo.Context) error {

	return c.String(http.StatusOK, "Hello")

}
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

	accessToken, _, tokenError := authService.GenerateToken(user)
	if tokenError != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, tokenError.Error())
	}

	// save refreshToken to cookie

	return c.JSON(http.StatusOK, loginResponse{
		AccessToken: accessToken.String(),
	})
}

func register(c echo.Context) error {
	return c.String(http.StatusOK, "Register!")
}
