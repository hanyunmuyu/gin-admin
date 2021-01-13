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
	db.DB.Offset((page - 1) * limit).Limit(limit).Find(&adminList).Offset(-1).Count(&count)
	return utils.NewPaginate(count, page, limit, adminList)
}
func (adminService *AdminService) UpdateAdmin(admin models.Admin) int64 {
	return db.DB.Save(&admin).RowsAffected
}
func (adminService *AdminService) GetAdminById(adminId uint) (admin models.Admin) {
	db.DB.First(&admin, adminId)
	return
}
func (adminService AdminService) GetAdminByName(name string) (admin models.Admin) {
	db.DB.Where("name=?", name).First(&admin)
	return
}
func (adminService AdminService) DeleteAdmin(adminId uint) int64 {
	return db.DB.Delete(&models.Admin{}, adminId).RowsAffected
}
