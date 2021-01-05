package utils

import "github.com/dgrijalva/jwt-go"

func ParseToken(token string, secret string) (jwt.Claims, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return claim.Claims, nil
}
