package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParseToken(ctx *gin.Context) (uint, error) {
	authorization := ctx.GetHeader("Authorization")
	claim, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
		return []byte("123456"), nil
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
