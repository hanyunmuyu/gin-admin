package admins

import (
	"gin-admin/app/http"
	"gin-admin/app/services/admins"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	userService = admins.UserService{}
)

type UserController struct {
	http.BaseController
}

// @Summary 用户列表
// @Security ApiKeyAuth
// @Description | 参数 | 说明 |备注|
// @Description | :-----: | :----: | :----: |
// @Description |page|页码||
// @Tags  admin
// @version 1.0
// @Param page query int false "页码" default(1)
// @success 200 {object} utils.JSONResult{data=[]models.User} "desc"
// @Router /admin/v1/user/list [GET]
func (userController *UserController) GetUserList(ctx *gin.Context) {
	page := 1
	if p, err := strconv.Atoi(ctx.DefaultQuery("page", "1")); err != nil {
		page = p
	}
	userList := userService.GetUserList(page, 15)
	userController.Success(ctx, userList)
}
