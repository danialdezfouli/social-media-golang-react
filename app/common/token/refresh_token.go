package token

import (
	"jupiter/config"
	"time"

	"github.com/golang-jwt/jwt"
)

type RefreshToken struct {
	token string
	expAt time.Time
}

func NewRefreshToken(claims *JwtCustomClaims) (*RefreshToken, error) {
	//exp := config.GetDuration("REFRESH_TOKEN_DURATION", time.Duration(time.Now().Add(time.Hour*24*7).Unix()))
	secret := config.GetString("REFRESH_TOKEN_SECRET", "")

	token, err := generateJWT(claims, secret)
	if err != nil {
		return nil, err
	}

	rt := new(RefreshToken)
	rt.token = token
	rt.expAt = time.Now().Add(config.GetDuration("REFRESH_TOKEN_DURATION",
		time.Duration(time.Now().Add(time.Hour*24*7).Unix())))

	return rt, nil
}

func VerifyRefreshToken(tokenStr string) (jwt.MapClaims, error) {
	token, claims, err := Verify(tokenStr, config.GetString("REFRESH_TOKEN_SECRET", ""))
	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	return claims, nil
}

func (t *RefreshToken) String() string {
	return t.token
}

func (t *RefreshToken) ExpiresAt() time.Time {
	return t.expAt
}
