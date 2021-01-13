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
	var rolePermission []models.Permission
	rolePermission = append(rolePermission, models.Permission{
		ApiPath:  "/admin/v1/role/list",
		Rule:     "/admin/role/list",
		Method:   "get",
		Title:    "角色列表",
		ParentId: 0,
		IsMenu:   0,
		Path:     "/admin/role/list",
	})
	permissionList = append(permissionList, rolePermission...)
	db.DB.Create(permissionList)
}

func (p PermissionSeeder) Drop() {
	db.DB.Exec("DROP TABLE IF EXISTS permissions")
}
