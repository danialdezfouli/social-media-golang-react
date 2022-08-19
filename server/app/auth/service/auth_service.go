package service

import (
	"errors"
	"fmt"
	"jupiter/app"
	"jupiter/app/auth/dto"
	"jupiter/app/common/token"
	"jupiter/app/model"
	"jupiter/config"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AuthService struct {
}

func (s AuthService) GenerateAccessToken(user *model.User) (*token.JwtCustomToken, error) {
	claims := &token.JwtCustomClaims{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	t, err := token.Generate(token.AccessToken, claims)

	if err != nil {
		return nil, errors.New("service.authService.GenerateAccessToken")
	}

	return t, nil
}

func (s AuthService) GenerateRefreshToken(user *model.User) (*token.JwtCustomToken, error) {
	claims := &token.JwtCustomClaims{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	t, err := token.Generate(token.RefreshToken, claims)

	if err != nil {
		return nil, errors.New("service.authService.GenerateAccessToken")
	}

	return t, nil
}

func (s AuthService) Response(c echo.Context, user *model.User) error {
	accessToken, atErr := s.GenerateAccessToken(user)
	if atErr != nil {
		return echo.ErrInternalServerError
	}

	refreshToken, rtErr := s.GenerateRefreshToken(user)
	if rtErr != nil {
		return echo.ErrInternalServerError
	}

	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken.String(),
		Path:     "/",
		Expires:  refreshToken.ExpiresAt(),
		Secure:   config.GetConfig().App.Production,
		HttpOnly: true,
		//SameSite: http.SameSiteNoneMode,
		SameSite: http.SameSiteLaxMode,
	})

	return c.JSON(http.StatusOK, echo.Map{
		"access_token": accessToken.String(),
	})
}

func (s AuthService) User(c echo.Context) (*model.User, error) {
	jwtUser := c.Get("user").(*jwt.Token)
	claims := jwtUser.Claims.(*token.JwtCustomClaims)

	user := &model.User{
		ID: claims.UserID,
	}

	result := app.GetDB().First(user)

	if result.Error != nil {
		return nil, echo.ErrUnauthorized
	}

	if user.Suspended {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "your account is suspended")
	}
	return user, nil
}

func (s AuthService) UpdateProfile(user *model.User, input *dto.UpdateProfileInput) {
	db := app.GetDB()

	result := db.Model(&user).Limit(1).Updates(map[string]interface{}{
		"name": input.Name,
		"bio":  input.Bio,
	})

	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}

}
