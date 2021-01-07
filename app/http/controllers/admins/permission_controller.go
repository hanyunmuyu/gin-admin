package admins

import (
	"gin-admin/app/http"
	"gin-admin/app/services/admins"
	"github.com/gin-gonic/gin"
)

var (
	permissionService = admins.PermissionService{}
)

type PermissionController struct {
	http.BaseController
}

// @Summary 权限列表
// @Security ApiKeyAuth
// @Description | 参数 | 说明 |备注|
// @Description | :-----: | :----: | :----: |
// @Description |apiPath|API地址||
// @Description |rule|规则||
// @Description |method|请求方法||
// @Description |title|标题||
// @Description |parentId|父级id||
// @Description |isMenu|是否是菜单||
// @Description |path|路径||
// @Tags  admin
// @version 1.0
// @success 200 {object} utils.JSONResult{data=[]models.Permission}
// @Router /admin/v1/permission/list [GET]
func (p *PermissionController) GetPermissionList(ctx *gin.Context) {
	permissionList := permissionService.GetPermissionList()
	p.Success(ctx, permissionList)
}
