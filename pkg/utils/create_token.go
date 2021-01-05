package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(key string, data interface{}) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		key:   data,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	token, err := at.SignedString([]byte("123456"))
	if err != nil {
		return "", err
	}
	return token, nil
}
