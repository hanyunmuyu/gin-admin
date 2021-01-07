package admins

import (
	"gin-admin/app/models"
	"gin-admin/db"
	"gin-admin/pkg/utils"
)

type RoleService struct {
}

func (r *RoleService) GetRoleList(page int, limit int) *utils.Paginate {
	var roleList []models.Role
	var total int64
	db.DB.Offset((page - 1) * limit).Limit(limit).Find(&roleList).Offset(0).Count(&total)
	return utils.NewPaginate(total, page, limit, roleList)
}
func (r *RoleService) GetRoleById(roleId uint) (role models.Role) {
	db.DB.First(&role, roleId)
	return
}

// 如果角色id为1则返回所有的权限，就是说拥有所有的权限
func (r *RoleService) GetRolePermission(role models.Role) (permissionList []models.Permission, err error) {
	if role.ID == 1 {
		db.DB.Find(&permissionList)
		return permissionList, nil
	}
	err = db.DB.Model(&role).Association("PermissionList").Find(&permissionList)
	return
}
