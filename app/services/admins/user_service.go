package admins

import (
	"gin-admin/app/models"
	"gin-admin/db"
	"gin-admin/pkg/utils"
)

type UserService struct {
}

func (userService *UserService) GetUserList(page, limit int) *utils.Paginate {
	var userList []models.User
	var count int64
	db.DB.Offset((page - 1) * limit).Limit(limit).Find(&userList).Offset(0).Count(&count)
	return utils.NewPaginate(count, page, limit, userList)
}

func (userService *UserService) UpdateUser(user models.User) error {
	return db.DB.Save(&user).Error
}
func (userService *UserService) GetUserById(userId uint) (user models.User) {
	db.DB.First(&user, userId)
	return
}
func (userService *UserService) DeleteUser(userId uint) int64 {
	return db.DB.Delete(&models.User{}, userId).RowsAffected
}
