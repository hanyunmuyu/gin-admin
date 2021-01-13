package admins

import (
	"gin-admin/app/http"
	"gin-admin/app/models"
	"gin-admin/app/services/admins"
	"gin-admin/pkg/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	adminService = admins.AdminService{}
)

type AdminController struct {
	http.BaseController
}

// @Summary 管理员列表
// @Security ApiKeyAuth
// @Description | 参数 | 说明 |备注|
// @Description | :-----: | :----: | :----: |
// @Description |page|页码||
// @Tags  admin
// @version 1.0
// @Param page query int false "页码" default(1)
// @success 200 {object} utils.JSONResult{data=utils.Paginate}
// @Router /admin/v1/admin/list [GET]
func (ac AdminController) GetAdminList(ctx *gin.Context) {
	page := 1
	if p, err := strconv.Atoi(ctx.DefaultQuery("page", "1")); err != nil {
		page = p
	}

	adminList := adminService.GetAdminList(page, 15)
	ac.Success(ctx, adminList)
}
func (ac *AdminController) UpdateAdmin(ctx *gin.Context) {
	adminForm := struct {
		AdminId uint `uri:"adminId" binding:"required"`
	}{}
	if err := ctx.ShouldBindUri(&adminForm); err != nil {
		lang := make(map[string]string)
		lang["AdminId"] = "管理员id"
		err := ac.Translate(err, lang)
		if err != nil {
			ac.Error(ctx, err.Error())
			return
		} else {
			ac.Error(ctx, "")
			return
		}
	}

	form := struct {
		Name     string `form:"name" binding:"required"`
		Password string `form:"password" binding:"omitempty"`
		RoleId   uint   `form:"roleId" binding:"required,gt=0"`
	}{}
	if err := ctx.ShouldBind(&form); err != nil {
		lang := make(map[string]string)
		lang["Name"] = "管理员名称"
		lang["Password"] = "密码"
		lang["RoleId"] = "角色id"
		err := ac.Translate(err, lang)
		if err != nil {
			ac.Error(ctx, err.Error())
			return
		} else {
			ac.Error(ctx, "")
			return
		}
	}
	adminOld := adminService.GetAdminByName(form.Name)
	if adminOld.ID != adminForm.AdminId {
		ac.Error(ctx, "管理员已经存在！换个名字试试！")
		return
	}
	admin := adminService.GetAdminById(adminForm.AdminId)
	if form.Password != "" {
		admin.Password = utils.EncodeMD5(form.Password)
	}
	admin.RoleId = form.RoleId
	adminService.UpdateAdmin(admin)
	ac.Success(ctx, gin.H{})
}
func (ac *AdminController) GetAdminInfo(ctx *gin.Context) {
	adminId, err := utils.ParseToken(ctx)
	if err != nil {
		ac.Error(ctx, err.Error())
		return
	}
	adminInfo := adminService.GetAdminById(adminId)
	ac.Success(ctx, adminInfo)
}
func (ac *AdminController) GetAdminPermissionList(ctx *gin.Context) {
	adminId, err := utils.ParseToken(ctx)
	if err != nil {
		ac.Error(ctx, err.Error())
		return
	}
	admin := adminService.GetAdminById(adminId)
	role := models.Role{}
	role.ID = admin.RoleId
	permissionList, _ := roleService.GetRolePermission(role)
	ac.Success(ctx, permissionList)
}
