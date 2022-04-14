package token

import (
	"jupiter/config"
	"time"

	"github.com/golang-jwt/jwt"
)

type AccessToken struct {
	token string
	expAt time.Time
}

func NewAccessToken(claims *JwtCustomClaims) (*AccessToken, error) {
	//exp := config.GetDuration("ACCESS_TOKEN_DURATION", time.Duration(time.Now().Add(time.Minute*15).Unix()))
	//secret := config.GetString("ACCESS_TOKEN_SECRET", "")
	// exp := config.GetConfig().Auth.AccessTokenDuration
	secret := config.GetConfig().Auth.AccessTokenSecret

	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	}

	token, err := generateJWT(claims, secret)
	if err != nil {
		return nil, err
	}

	at := new(AccessToken)
	at.token = token
	at.expAt = time.Now().Add(config.GetDuration("ACCESS_TOKEN_DURATION",
		time.Duration(time.Now().Add(time.Minute*15).Unix())))

	return at, nil
}

func VerifyAccessToken(tokenStr string) (jwt.MapClaims, error) {
	token, claims, err := Verify(tokenStr, config.GetString("ACCESS_TOKEN_SECRET", ""))
	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	return claims, nil
}

func (t *AccessToken) String() string {
	return t.token
}

func (t *AccessToken) ExpiresAt() time.Time {
	return t.expAt
}
