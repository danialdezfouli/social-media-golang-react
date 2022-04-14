package service

import (
	"errors"
	"jupiter/app/common/token"
	"jupiter/app/model"
)

type AuthService struct {
	User *model.User
}

func (s AuthService) GenerateToken(user *model.User) (*token.AccessToken, *token.RefreshToken, error) {
	accessToken, err := token.NewAccessToken(&token.JwtCustomClaims{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
	})

	if err != nil {
		return nil, nil, errors.New("service.authService.GenerateToken")
	}

	refreshToken, err := token.NewRefreshToken(&token.JwtCustomClaims{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
	})

	return accessToken, refreshToken, nil

}

func (s AuthService) RefreshToken(user *model.User, accessToken string) (string, string, error) {

	token := "at"

	return token, "---", nil

}
