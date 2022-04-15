package token

import (
	"github.com/golang-jwt/jwt"
	"jupiter/config"
	"time"
)

type JwtTokenType int

const (
	AccessToken JwtTokenType = iota + 1
	RefreshToken
)

type JwtCustomClaims struct {
	UserID   uint   `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

type JwtCustomToken struct {
	token     string
	expAt     time.Time
	tokenType JwtTokenType
}

func (t *JwtCustomToken) String() string {
	return t.token
}

func (t *JwtCustomToken) ExpiresAt() time.Time {
	return t.expAt
}

func getTokenConfigs(tokenType JwtTokenType) (expire time.Time, secret string) {
	isAccessToken := tokenType == AccessToken

	if isAccessToken {
		expire = time.Now().Add(config.GetConfig().Auth.AccessTokenDuration)
		secret = config.GetConfig().Auth.AccessTokenSecret
	} else {
		expire = time.Now().Add(config.GetConfig().Auth.RefreshTokenDuration)
		secret = config.GetConfig().Auth.RefreshTokenSecret
	}

	return
}

//
//func Verify(token string, secretKey string) (*jwt.Token, jwt.MapClaims, error) {
//	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
//		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
//		}
//		return []byte(secretKey), nil
//	})
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return t, t.Claims.(jwt.MapClaims), nil
//}
