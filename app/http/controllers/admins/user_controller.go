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

func (userController *UserController) GetUserList(ctx *gin.Context) {
	page := 1
	if p, err := strconv.Atoi(ctx.DefaultQuery("page", "1")); err != nil {
		page = p
	}
	userList := userService.GetUserList(page, 15)
	userController.Success(ctx, userList)
}
