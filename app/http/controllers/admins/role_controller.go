package admins

import (
	"gin-admin/app/http"
	"gin-admin/app/models"
	"gin-admin/app/services/admins"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	roleService = admins.RoleService{}
)

type RoleController struct {
	http.BaseController
}

// @Summary 角色列表
// @Security ApiKeyAuth
// @Description | 参数 | 说明 |备注|
// @Description | :-----: | :----: | :----: |
// @Description |roleName|角色名称||
// @Tags  admin
// @version 1.0
// @Param page query int false "页码" default(1)
// @success 200 {object} utils.JSONResult{data=[]utils.Paginate}
// @Router /admin/role/list [GET]
func (r *RoleController) GetRoleList(ctx *gin.Context) {
	page := 1
	if p, err := strconv.Atoi(ctx.DefaultQuery("page", "1")); err == nil {
		page = p
	}
	roleList := roleService.GetRoleList(page, 15)
	r.Success(ctx, roleList)
}

// @Summary 角色详情
// @Security ApiKeyAuth
// @Description | 参数 | 说明 |备注|
// @Description | :-----: | :----: | :----: |
// @Description |roleName|角色名称||
// @Description |permissionList|权限id||
// @Description |id|权限列表|permissionList列表下面的|
// @Description |apiPath|API地址||
// @Description |rule|匹配规则||
// @Description |method|请求方法||
// @Description |title|权限名称||
// @Description |parentId|parentId||
// @Description |isMenu|是否是菜单|1是0不是|
// @Description |path|路由地址||
// @Param roleId path int true "角色id" minimum(1)
// @Tags  admin
// @version 1.0
// @success 200 {object} utils.JSONResult{data=models.Role}
// @Router /admin/role/detail/{roleId} [GET]
func (r *RoleController) GetRoleDetail(ctx *gin.Context) {
	form := struct {
		RoleId uint `json:"roleId" uri:"roleId" binding:"required,gte=1"`
	}{}
	if err := ctx.ShouldBindUri(&form); err != nil {
		lang := map[string]string{}
		lang["roleId"] = "角色id"
		err = r.Translate(err, lang)
		if err != nil {
			r.Error(ctx, err.Error())
			return
		}
		r.Error(ctx, "")
		return
	}
	role := roleService.GetRoleById(form.RoleId)
	if role.ID == 0 {
		r.Error(ctx, "角色不存在")
		return
	}
	rolePermissionList, err := roleService.GetRolePermission(role)
	permissionAll := permissionService.GetPermissionList()
	if err != nil {
		r.Error(ctx, err.Error())
	}
	r.Success(ctx, gin.H{"permissionList": rolePermissionList, "permissionAll": permissionAll})
}

// @Summary 更新角色
// @Security ApiKeyAuth
// @accept x-www-form-urlencoded
// @Description | 参数 | 说明 |备注|
// @Description | :-----: | :----: | :----: |
// @Param roleId path uint true "角色id" minimum(1)
// @Param permissionId formData []int true "权限id" collectionFormat(multi)
// @Tags  admin
// @version 1.0
// @success 200 {object} utils.JSONResult{}
// @Router /admin/role/update/{roleId} [PUT]
func (r *RoleController) UpdateRole(ctx *gin.Context) {
	form := struct {
		RoleId uint `uri:"roleId" binding:"required,gte=1"`
	}{}
	permissionList := struct {
		PermissionId []uint `form:"permissionList" json:"permissionList" binding:"required"`
	}{}
	if err := ctx.ShouldBindUri(&form); err != nil {
		lang := map[string]string{}
		lang["roleId"] = "角色id"
		err = r.Translate(err, lang)
		if err != nil {
			r.Error(ctx, err.Error())
			return
		}
		r.Error(ctx, "")
		return
	}
	if err := ctx.ShouldBind(&permissionList); err != nil {
		lang := map[string]string{}
		lang["permissionList"] = "权限列表"
		err = r.Translate(err, lang)
		if err != nil {
			r.Error(ctx, err.Error())
			return
		}
		r.Error(ctx, "")
		return
	}
	if form.RoleId == 1 {
		r.Success(ctx, "更新成功")
		return
	}
	role := roleService.GetRoleById(form.RoleId)
	if role.ID == 0 {
		r.Error(ctx, "角色不存在")
		return
	}
	row := roleService.UpdateRole(role)
	roleService.DeleteRolePermission(role.ID)

	var rolePermissionList []models.RolePermission
	for _, permission := range permissionService.GetPermissionListByIdList(permissionList.PermissionId) {
		rolePermissionList = append(rolePermissionList, models.RolePermission{
			RoleId:       role.ID,
			PermissionId: permission.ID,
		})
	}
	roleService.AddRolePermission(rolePermissionList)
	if row <= 0 {
		r.Error(ctx, "更新失败")
	}
	r.Success(ctx, "更新成功")
}

// @Summary 删除角色
// @Security ApiKeyAuth
// @Description | 参数 | 说明 |备注|
// @Description | :-----: | :----: | :----: |
// @Param roleId path uint true "角色id" minimum(1)
// @Tags  admin
// @version 1.0
// @success 200 {object} utils.JSONResult{}
// @Router /admin/role/delete/{roleId} [DELETE]
func (r *RoleController) DeleteRole(ctx *gin.Context) {
	form := struct {
		RoleId uint `uri:"roleId" binding:"required,gte=1"`
	}{}
	if err := ctx.ShouldBindUri(&form); err != nil {
		lang := map[string]string{}
		lang["roleId"] = "角色id"
		err = r.Translate(err, lang)
		if err != nil {
			r.Error(ctx, err.Error())
			return
		}
		r.Error(ctx, "")
		return
	}
	if form.RoleId == 1 {
		r.Error(ctx, "角色id为1的不可以删除")
		return
	}
	role := roleService.GetRoleById(form.RoleId)
	if role.ID == 0 {
		r.Error(ctx, "角色不存在")
		return
	}
	if row := roleService.DeleteRole(role.ID); row <= 0 {
		r.Error(ctx, "")
	} else {
		roleService.DeleteRolePermission(role.ID)
		r.Success(ctx, "删除成功")
	}
}
func (r *RoleController) AddRole(ctx *gin.Context) {
	form := struct {
		RoleName     string `form:"roleName" binding:"required"`
		PermissionId []uint `form:"permissionList" json:"permissionList" binding:"required"`
	}{}
	if err := ctx.ShouldBind(&form); err != nil {
		lang := map[string]string{}
		lang["permissionList"] = "权限列表"
		lang["RoleName"] = "角色名称"
		err = r.Translate(err, lang)
		if err != nil {
			r.Error(ctx, err.Error())
			return
		}
		r.Error(ctx, "")
	}
	role := roleService.GetRoleByRoleName(form.RoleName)
	if role.ID > 0 {
		r.Error(ctx, "角色已经存在")
		return
	}
	role.RoleName = form.RoleName
	row := roleService.AddRole(role)

	if row <= 0 {
		r.Error(ctx, "新怎失败")
	}
	var rolePermissionList []models.RolePermission
	for _, permission := range permissionService.GetPermissionListByIdList(form.PermissionId) {
		rolePermissionList = append(rolePermissionList, models.RolePermission{
			RoleId:       role.ID,
			PermissionId: permission.ID,
		})
	}
	roleService.AddRolePermission(rolePermissionList)
	r.Success(ctx, gin.H{})
}
func (r RoleController) GetAllRole(ctx *gin.Context) {
	roleList := roleService.GetAllRole()
	r.Success(ctx, roleList)
}
