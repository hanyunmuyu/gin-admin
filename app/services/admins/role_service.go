package admins

import (
	"gin-admin/app/models"
	"gin-admin/db"
	"gin-admin/pkg/utils"
)

type RoleService struct {
}

func (r RoleService) GetRoleList(page int, limit int) *utils.Paginate {
	var roleList []models.Role
	var total int64
	db.DB.Offset((page - 1) * limit).Limit(limit).Find(&roleList).Offset(0).Count(&total)
	return utils.NewPaginate(total, page, limit, roleList)
}
