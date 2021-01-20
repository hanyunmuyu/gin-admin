package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func getKey() string {
	var sKey string
	v := Config()
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		sKey = v.GetString("jwt.signingKey")
	})
	return sKey
}
func CreateToken(key string, data interface{}) (string, error) {
	sKey := getKey()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		key:   data,
		"exp": time.Now().Add(time.Hour * time.Duration(24*Config().GetInt("jwt.expiresAt"))).Unix(),
	})
	token, err := at.SignedString([]byte(sKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(ctx *gin.Context) (uint, error) {
	sKey := getKey()
	authorization := ctx.GetHeader("Authorization")
	claim, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
		return []byte(sKey), nil
	})
	if err != nil {
		return 0, err
	}
	if admin, ok := claim.Claims.(jwt.MapClaims); ok {
		adminMap, _ := admin["admin"].(map[string]interface{})
		id, _ := adminMap["id"]
		adminId := 0
		if i, err := strconv.Atoi(fmt.Sprintf("%v", id)); err == nil {
			adminId = i
			return uint(i), nil
		} else {
			return uint(adminId), err

		}
	}
	return 0, nil
}
