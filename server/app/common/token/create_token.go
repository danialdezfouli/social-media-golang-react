package token

import (
	"github.com/golang-jwt/jwt"
)

func Generate(tokenType JwtTokenType, claims *JwtCustomClaims) (*JwtCustomToken, error) {
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
