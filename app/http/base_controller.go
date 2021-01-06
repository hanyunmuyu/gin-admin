package http

import (
	"errors"
	"gin-admin/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

type BaseController struct {
}

func (base *BaseController) Success(ctx *gin.Context, data interface{}) {

	ctx.JSON(http.StatusOK, gin.H{"msg": "success", "code": 200, "data": data})
}
func (base *BaseController) Error(ctx *gin.Context, msg string) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"msg": msg, "code": 0})
}
func (base *BaseController) Translate(err error, lang map[string]string) error {
	errs := err.(validator.ValidationErrors)
	for _, e := range errs {
		key := e.Field()
		if _, ok := lang[key]; ok {
			return errors.New(strings.ReplaceAll(e.Translate(utils.Trans), key, lang[key]))
		} else {
			return errors.New(e.Translate(utils.Trans))
		}
	}
	return nil
}
