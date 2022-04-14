package config

import (
	"time"
)

type AuthConfig struct {
	AccessTokenSecret   string
	AccessTokenDuration time.Duration

	RefreshTokenSecret   string
	RefreshTokenDuration time.Duration
}

func newAuthConfig() *AuthConfig {
	return &AuthConfig{
		AccessTokenDuration: time.Minute * 15,
		AccessTokenSecret:   GetEnv("ACCESS_TOKEN_SECRET"),

		RefreshTokenDuration: time.Hour * 24 * 30,
		RefreshTokenSecret:   GetEnv("REFRESH_TOKEN_SECRET"),
	}
}
