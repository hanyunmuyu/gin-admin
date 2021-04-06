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
// @Description |name|用户名||
// @Description |mobile|手机号||
// @Description |email|邮箱||
// @Description |page|页码||
// @Tags  admin
// @version 1.0
// @Param page query int false "页码" default(1)
// @success 200 {object} utils.JSONResult{data=utils.Paginate}
// @Router /admin/user/list [GET]
func (userController *UserController) GetUserList(ctx *gin.Context) {
	page := 1
	if p, err := strconv.Atoi(ctx.DefaultQuery("page", "1")); err == nil {
		page = p
	}
	keyword := ctx.DefaultQuery("keyword", "")
	startDate := ctx.DefaultQuery("startDate", "")
	endDate := ctx.DefaultQuery("endDate", "")
	userList := userService.GetUserList(page, keyword, startDate, endDate, 15)
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
// @Param password formData string false "密码" minlength(5)
// @Param mobile formData string false "手机号"
// @Param email formData string false "邮箱"
// @Param avatar formData string false "头像"
// @Param name formData string false "用户名"
// @success 200 {object} utils.JSONResult{} "更新成功"
// @Router /admin/user/{userId} [PUT]
func (userController *UserController) UpdateUser(ctx *gin.Context) {
	form := struct {
		UserId uint `json:"userId" uri:"userId" binding:"required"`
	}{}
	if err := ctx.ShouldBindUri(&form); err != nil {
		lang := make(map[string]string)
		lang["UserId"] = "用户id"
		err := userController.Translate(err, lang)
		if err != nil {
			userController.Error(ctx, err.Error())
			return
		} else {
			userController.Error(ctx, "")
			return
		}
	}
	userForm := struct {
		Password string `json:"password" form:"password"`
		Mobile   string `json:"mobile" form:"mobile" binding:"omitempty,mobile"`
		Email    string `json:"email" form:"email" binding:"omitempty,email"`
		Avatar   string `json:"avatar" form:"avatar"`
		Name     string `json:"name" form:"name"`
	}{}

	if err := ctx.ShouldBind(&userForm); err != nil {
		lang := make(map[string]string)
		lang["Password"] = "密码"
		lang["Mobile"] = "手机号"
		lang["Email"] = "邮箱"
		lang["Name"] = "用户名"
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
	if userForm.Email != "" {
		if u := userService.GetUserByEmail(userForm.Email); u.ID != user.ID {
			userController.Error(ctx, "邮箱已经存在")
			return
		}
		user.Email = userForm.Email
	}
	if userForm.Mobile != "" {
		if u, row := userService.GetUserByMobile(userForm.Mobile); row > 0 && u.ID != user.ID {
			userController.Error(ctx, "手机号已经存在")
			return
		}
		user.Mobile = userForm.Mobile
	}
	if userForm.Name != "" {
		if u := userService.GetUserByName(userForm.Name); u.ID != user.ID {
			userController.Error(ctx, "用户名已经存在")
			return
		}
		user.Name = userForm.Name
	}
	if userForm.Avatar != "" {
		user.Avatar = userForm.Avatar
	}
	if userForm.Password != "" {
		user.Password = utils.EncodeMD5(userForm.Password)
	}
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
// @Router /admin/user/{userId} [delete]
func (userController *UserController) DeleteUser(ctx *gin.Context) {
	form := struct {
		UserId uint `json:"userId" uri:"userId" binding:"required"`
	}{}
	if err := ctx.ShouldBindUri(&form); err != nil {
		lang := make(map[string]string)
		lang["UserId"] = "用户id"
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
func (u UserController) AddUser(ctx *gin.Context) {
	form := struct {
		Name   string `form:"name"`
		Mobile string `form:"mobile"`
	}{}
	if err := ctx.ShouldBind(&form); err != nil {
		u.Error(ctx, err.Error())
		return
	}
	user := userService.GetUserByName(form.Name)
	if user.ID > 0 {
		u.Error(ctx, "用户已经存在")
		return
	}
	user.Name = form.Name
	user.Mobile = form.Mobile
	row := userService.AddUser(user)
	if row <= 0 {
		u.Error(ctx, "添加失败")
		return
	}
	u.Success(ctx, gin.H{})
}
