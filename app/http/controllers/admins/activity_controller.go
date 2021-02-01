package admins

import (
	"gin-admin/app/http"
	"gin-admin/app/services/admins"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	activityService = admins.NewActivityService()
)

type ActivityController struct {
	http.BaseController
}

func (a ActivityController) GetActivityList(ctx *gin.Context) {
	page := 1
	if p, err := strconv.Atoi(ctx.DefaultQuery("page", "1")); err == nil {
		page = p
	}
	activityList := activityService.GetActivityList(page, 15)
	a.Success(ctx, activityList)
}
