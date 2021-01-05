package admins

import (
	"gin-admin/app/http"
	"gin-admin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	http.BaseController
}

// @Summary 登录
// @Description | 参数 | 说明 |备注|
// @Description | :-----: | :----: | :----: |
// @Description |token|用户的token||
// @Tags  admin
// @Accept mpfd
// @Produce json
// @version 1.0
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @success 200 {object} utils.JSONResult{data=map[string]string} "成功"
// @Router /admin/v1/login [POST]
func (login *LoginController) Login(ctx *gin.Context) {
	loginStruct := struct {
		Name     string `json:"name" form:"name" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}{}
	if err := ctx.ShouldBind(&loginStruct); err != nil {
		lang := make(map[string]string)
		lang["Name.required"] = "用户名"
		lang["Password.required"] = "密码"
		err := login.Translate(err, lang)
		if err != nil {
			login.Error(ctx, err.Error())
		} else {
			login.Error(ctx, "")
		}
		return
	}
	admin := adminService.GetAdminByName(loginStruct.Name)
	if admin.ID == 0 {
		login.Error(ctx, "用户不存在")
		return
	}
	if admin.Password != utils.EncodeMD5(loginStruct.Password) {
		login.Error(ctx, "密码错误")
		return
	}
	token, _ := utils.CreateToken("admin", admin)
	login.Success(ctx, gin.H{"token": token})
}
