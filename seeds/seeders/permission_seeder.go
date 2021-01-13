package seeders

import (
	"gin-admin/app/models"
	"gin-admin/db"
)

type PermissionSeeder struct {
}

func NewPermissionSeeder() *PermissionSeeder {
	return &PermissionSeeder{}
}

func (p PermissionSeeder) Run() {
	var permissionList []models.Permission
	permissionList = append(permissionList, models.Permission{
		ApiPath:        "",
		Rule:           "",
		Method:         "GET",
		Title:          "仪表盘",
		ParentId:       0,
		IsMenu:         1,
		Path:           "/admin/dashboard",
		PermissionList: nil,
	})
	var rolePermission models.Permission
	var rolePermissionList []models.Permission
	rolePermission = models.Permission{
		ApiPath:  "/admin/v1/role/list",
		Rule:     "/admin/role/list",
		Method:   "get",
		Title:    "角色管理",
		ParentId: 0,
		IsMenu:   1,
		Path:     "/admin/role",
	}
	rolePermission.PermissionList = append(rolePermissionList, models.Permission{
		ApiPath:  "/admin/v1/role/list",
		Rule:     "/admin/role/list",
		Method:   "get",
		Title:    "角色列表",
		ParentId: 0,
		IsMenu:   1,
		Path:     "/admin/role/list",
	})

	rolePermissionList = append(rolePermissionList, rolePermission)

	var adminPermission = models.Permission{
		ApiPath:  "/admin/v1/role/list",
		Rule:     "/admin/role/list",
		Method:   "get",
		Title:    "管理员管理",
		ParentId: 0,
		IsMenu:   1,
		Path:     "/admin/list",
	}
	var adminPermissionList []models.Permission
	adminPermissionList = append(adminPermissionList, models.Permission{
		ApiPath:  "/admin/v1/role/list",
		Rule:     "/admin/role/list",
		Method:   "get",
		Title:    "管理员列表",
		ParentId: 0,
		IsMenu:   0,
		Path:     "/admin/admin/list",
	})
	adminPermission.PermissionList = adminPermissionList

	permissionList = append(permissionList, rolePermission)
	permissionList = append(permissionList, adminPermission)

	for _, permission := range permissionList {
		db.DB.Create(&permission)
		if len(permission.PermissionList) > 0 {
			for i, p := range permission.PermissionList {
				p.ParentId = permission.ID
				permission.PermissionList[i] = p
			}
			db.DB.Create(permission.PermissionList)
		}
	}
}

func (p PermissionSeeder) Drop() {
	db.DB.Exec("DROP TABLE IF EXISTS permissions")
}
