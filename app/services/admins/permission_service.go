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
func (ps *PermissionService) GetPermissionListByIdList(idList []uint) (permissionList []models.Permission) {
	db.DB.Where("id IN (?)", idList).Find(&permissionList)
	return
}
