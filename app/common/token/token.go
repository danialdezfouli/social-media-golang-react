package token

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	UserID   uint   `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func generateJWT(claims *JwtCustomClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func Verify(token string, secretKey string) (*jwt.Token, jwt.MapClaims, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, nil, err
	}

	return t, t.Claims.(jwt.MapClaims), nil
}
