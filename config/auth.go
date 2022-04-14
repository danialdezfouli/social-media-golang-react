package config

import "time"

type AuthConfig struct {
	AccessTokenSecret   string
	AccessTokenDuration time.Duration

	RefreshTokenSecret   string
	RefreshTokenDuration time.Duration
}

func newAuthConfig() *AuthConfig {
	return &AuthConfig{
		AccessTokenDuration: GetDuration("ACCESS_TOKEN_DURATION", time.Duration(time.Now().Add(time.Minute*15).Unix())),
		AccessTokenSecret:   GetString("ACCESS_TOKEN_SECRET", ""),

		RefreshTokenDuration: GetDuration("REFRESH_TOKEN_DURATION", time.Duration(time.Now().Add(time.Minute*15).Unix())),
		RefreshTokenSecret:   GetString("REFRESH_TOKEN_SECRET", ""),
	}
}
