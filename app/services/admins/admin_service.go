package admins

import (
	"gin-admin/app/models"
	"gin-admin/db"
	"gin-admin/pkg/utils"
)

type AdminService struct {
}

func (adminService *AdminService) GetAdminList(page int, limit int) *utils.Paginate {
	var adminList []models.Admin
	var count int64
	db.DB.Offset((page - 1) * limit).Limit(limit).Find(&adminList).Offset(0).Count(&count)
	return utils.NewPaginate(count, page, limit, adminList)
}
func (adminService *AdminService) UpdateAdmin(adminId uint) {
	db.DB.Model(&models.Admin{}).Where("id=?", adminId).Update("name", "admin1")
}
func (adminService *AdminService) GetAdminById(adminId uint) (admin models.Admin) {
	db.DB.First(&admin, adminId)
	return
}
func (adminService AdminService) GetAdminByName(name string) (admin models.Admin) {
	db.DB.Where("name=?", name).First(&admin)
	return
}
