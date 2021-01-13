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
	db.DB.Offset((page - 1) * limit).Limit(limit).Find(&roleList).Offset(-1).Count(&total)
	return utils.NewPaginate(total, page, limit, roleList)
}
func (r *RoleService) GetRoleById(roleId uint) (role models.Role) {
	db.DB.First(&role, roleId)
	return
}
func (r *RoleService) GetRoleByRoleName(roleName string) (role models.Role) {
	db.DB.First(&role, "role_name=?", roleName)
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
func (r *RoleService) UpdateRole(role models.Role) int64 {
	return db.DB.Save(&role).RowsAffected
}

func (r *RoleService) DeleteRolePermission(roleId uint) int64 {
	return db.DB.Where("role_id=?", roleId).Delete(&models.RolePermission{}).RowsAffected
}
func (r *RoleService) DeleteRole(roleId uint) int64 {
	return db.DB.Delete(&models.Role{}, roleId).RowsAffected
}
func (r *RoleService) AddRolePermission(rolePermissionList []models.RolePermission) {
	db.DB.Create(&rolePermissionList)
}
func (r *RoleService) AddRole(role models.Role) int64 {
	return db.DB.Create(&role).RowsAffected
}
