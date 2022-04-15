package auth

import (
	"github.com/labstack/echo/v4"
	"jupiter/app/auth/dto"
	"jupiter/app/auth/service"
	"net/http"
)

func me(c echo.Context) error {
	user, err := service.AuthService{}.User(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, meResponse{
		Name:      user.Name,
		Username:  user.Username,
		Suspended: user.Suspended,
		Image:     user.Image,
	})
}

func login(c echo.Context) error {
	authService := service.AuthService{}
	loginService := service.LoginService{}
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

	return authService.Response(c, user)

}

func register(c echo.Context) error {
	authService := service.AuthService{}
	registerService := service.RegisterService{}
	input := new(dto.RegisterInput)

	if err := c.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(input); err != nil {
		return err
	}

	user, err := registerService.Register(c, input)
	if err != nil {
		return err
	}

	return authService.Response(c, user)
}

func refreshToken(c echo.Context) error {
	authService := service.AuthService{}
	user, err := authService.User(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return authService.Response(c, user)
}
