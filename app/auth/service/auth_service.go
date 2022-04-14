package service

import (
	"errors"
	"jupiter/app/common/token"
	"jupiter/app/model"
)

type AuthService struct {
	User *model.User
}

func (s AuthService) GenerateAccessToken(user *model.User) (*token.JwtCustomToken, error) {
	claims := &token.JwtCustomClaims{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	token, err := token.Generate(token.AccessToken, claims)

	if err != nil {
		return nil, errors.New("service.authService.GenerateAccessToken")
	}

	return token, nil
}

func (s AuthService) GenerateRefreshToken(user *model.User) (*token.JwtCustomToken, error) {
	claims := &token.JwtCustomClaims{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	token, err := token.Generate(token.RefreshToken, claims)

	if err != nil {
		return nil, errors.New("service.authService.GenerateAccessToken")
	}

	return token, nil
}
