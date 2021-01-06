package admins

import (
	"gin-admin/app/http"
	"gin-admin/app/services/admins"
	"gin-admin/pkg/utils"
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

// @Summary 更新用户信息
// @Security ApiKeyAuth
// @Description | 参数 | 说明 |备注|
// @Description | :-----: | :----: | :----: |
// @Accept mpfd
// @Tags  admin
// @version 1.0
// @Param userId path int true "用户id" min(1)
// @Param password formData string true "密码" minlength(5)
// @success 200 {object} utils.JSONResult{} "更新成功"
// @Router /admin/v1/user/{userId} [PUT]
func (userController *UserController) UpdateUser(ctx *gin.Context) {
	form := struct {
		UserId uint `json:"userId" uri:"userId" binding:"required"`
	}{}
	if err := ctx.ShouldBindUri(&form); err != nil {
		lang := make(map[string]string)
		lang["UserId.required"] = "用户id"
		err := userController.Translate(err, lang)
		if err != nil {
			userController.Error(ctx, err.Error())
			return
		} else {
			userController.Error(ctx, "")
			return
		}
	}
	password := struct {
		Password string `json:"password" form:"password" binding:"required"`
	}{}

	if err := ctx.ShouldBind(&password); err != nil {
		lang := make(map[string]string)
		lang["Password.required"] = "密码"
		err := userController.Translate(err, lang)
		if err != nil {
			userController.Error(ctx, err.Error())
			return
		} else {
			userController.Error(ctx, "")
			return
		}
	}
	user := userService.GetUserById(form.UserId)
	if user.ID == 0 {
		userController.Error(ctx, "用户不存在")
		return
	}
	user.Password = utils.EncodeMD5(password.Password)
	err := userService.UpdateUser(user)
	if err == nil {
		userController.Success(ctx, gin.H{})
	} else {
		userController.Error(ctx, err.Error())
	}
}

// @Summary 删除用户
// @Security ApiKeyAuth
// @Description | 参数 | 说明 |备注|
// @Description | :-----: | :----: | :----: |
// @Accept mpfd
// @Tags  admin
// @version 1.0
// @Param userId path int true "用户id" min(1)
// @success 200 {object} utils.JSONResult{} "删除成功"
// @Router /admin/v1/user/{userId} [delete]
func (userController *UserController) DeleteUser(ctx *gin.Context) {
	form := struct {
		UserId uint `json:"userId" uri:"userId" binding:"required"`
	}{}
	if err := ctx.ShouldBindUri(&form); err != nil {
		lang := make(map[string]string)
		lang["UserId.required"] = "用户id"
		err := userController.Translate(err, lang)
		if err != nil {
			userController.Error(ctx, err.Error())
			return
		} else {
			userController.Error(ctx, "")
			return
		}
	}
	user := userService.GetUserById(form.UserId)
	if user.ID == 0 {
		userController.Error(ctx, "用户不存在")
		return
	}
	if row := userService.DeleteUser(form.UserId); row > 0 {
		userController.Success(ctx, gin.H{})
	} else {
		userController.Error(ctx, "删除失败")
	}
}
