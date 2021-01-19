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
		PermissionList: []models.Permission{
			models.Permission{
				ApiPath:  "/role/update/:roleId",
				Rule:     "",
				Method:   "PUT",
				Title:    "编辑角色",
				ParentId: 0,
				IsMenu:   0,
				Path:     "editRole",
			},
			models.Permission{
				ApiPath:  "/role/update/:roleId",
				Rule:     "",
				Method:   "DELETE",
				Title:    "删除角色",
				ParentId: 0,
				IsMenu:   0,
				Path:     "deleteRole",
			},
		},
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
		PermissionList: []models.Permission{
			models.Permission{
				ApiPath:  "/admin/admin/update/:adminId",
				Rule:     "",
				Method:   "PUT",
				Title:    "编辑管理员信息",
				ParentId: 0,
				IsMenu:   0,
				Path:     "editAdmin",
			},
			models.Permission{
				ApiPath:  "/admin/admin/update/:adminId",
				Rule:     "",
				Method:   "PUT",
				Title:    "删除管理员",
				ParentId: 0,
				IsMenu:   0,
				Path:     "deleteAdmin",
			},
		},
	})
	adminPermission.PermissionList = adminPermissionList

	var userPermission models.Permission
	var userPermissionList []models.Permission
	userPermission = models.Permission{
		ApiPath:  "/admin/user",
		Rule:     "/admin/user",
		Method:   "GET",
		Title:    "用户管理",
		ParentId: 0,
		IsMenu:   1,
		Path:     "/admin/user",
	}
	userPermissionList = append(userPermissionList, models.Permission{
		ApiPath:  "",
		Rule:     "",
		Method:   "GET",
		Title:    "用户列表",
		ParentId: 0,
		IsMenu:   0,
		Path:     "/admin/user/list",
		PermissionList: []models.Permission{
			models.Permission{
				ApiPath:  "",
				Rule:     "",
				Method:   "DELETE",
				Title:    "删除用户",
				ParentId: 0,
				IsMenu:   0,
				Path:     "deleteUser",
			},
			models.Permission{
				ApiPath:  "",
				Rule:     "",
				Method:   "PUT",
				Title:    "编辑用户信息",
				ParentId: 0,
				IsMenu:   0,
				Path:     "editUser",
			},
		},
	})
	userPermission.PermissionList = userPermissionList

	permissionList = append(permissionList, rolePermission)
	permissionList = append(permissionList, adminPermission)
	permissionList = append(permissionList, userPermission)

	initPermission(permissionList, 0)
}
func initPermission(permissionList []models.Permission, parentId uint) {
	for _, permission := range permissionList {
		permission.ParentId = parentId
		db.DB.Create(&permission)
		if len(permission.PermissionList) > 0 {
			initPermission(permission.PermissionList, permission.ID)
		}
	}
}
func (p PermissionSeeder) Drop() {
	db.DB.Exec("DROP TABLE IF EXISTS permissions")
}
