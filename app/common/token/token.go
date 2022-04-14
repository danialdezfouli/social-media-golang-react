package token

import (
	"github.com/golang-jwt/jwt"
	"jupiter/config"
	"time"
)

const (
	AccessToken  = 0
	RefreshToken = 1
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
	tokenType int
}

func Generate(tokenType int, claims *JwtCustomClaims) (*JwtCustomToken, error) {
	expireAt, secret := getTokenConfigs(tokenType)

	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expireAt.Unix(),
	}

	tokenStr, err := generateJWT(claims, secret)
	if err != nil {
		return nil, err
	}

	t := &JwtCustomToken{
		token:     tokenStr,
		tokenType: tokenType,
		expAt:     expireAt,
	}

	return t, nil
}

func generateJWT(claims *JwtCustomClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
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

func getTokenConfigs(tokenType int) (expire time.Time, secret string) {
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

func (t *JwtCustomToken) String() string {
	return t.token
}

func (t *JwtCustomToken) ExpiresAt() time.Time {
	return t.expAt
}
