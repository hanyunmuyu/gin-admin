package middleware

import (
	"gin-admin/pkg/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func AdminJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader("Authorization")
		if !strings.HasPrefix(authorization, "Bearer ") {
			ctx.AbortWithStatusJSON(200, gin.H{"code": 4000, "msg": "JWT认证应该以'Bearer '开头"})
		} else {
			token := strings.Split(authorization, " ")
			if _, err := utils.ParseToken(token[len(token)-1], "123456"); err == nil {
				ctx.Next()
			} else {
				ctx.AbortWithStatusJSON(200, gin.H{"code": 4000, "msg": err.Error()})
			}
		}
	}
}
