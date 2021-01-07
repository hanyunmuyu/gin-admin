package admins

import (
	"gin-admin/app/models"
	"gin-admin/db"
)

type PermissionService struct {
}

func (ps *PermissionService) GetPermissionList() (permissionList []models.Permission) {
	db.DB.Find(&permissionList)
	return
}
