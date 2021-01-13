package middleware

import (
	"gin-admin/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AdminJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if _, err := utils.ParseToken(ctx); err == nil {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(200, gin.H{"code": 4000, "msg": err.Error()})
		}
	}
}
