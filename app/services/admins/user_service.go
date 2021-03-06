package admins

import (
	"gin-admin/app/models"
	"gin-admin/db"
	"gin-admin/pkg/utils"
)

type UserService struct {
}

func (userService *UserService) GetUserList(page int, keyword string, startDate, endDate string, limit int) *utils.Paginate {
	var userList []models.User
	var count int64
	d := db.DB.Offset((page - 1) * limit)
	if keyword != "" {
		d.Where(
			d.Where("name like ?", "%"+keyword+"%"),
		)
	}
	if startDate != "" {
		d.Where("created_at>=?", startDate)
	}
	if endDate != "" {
		d.Where("created_at<=?", endDate)
	}
	d.Limit(limit).Order("id desc").Find(&userList).Offset(-1).Count(&count)
	return utils.NewPaginate(count, page, limit, userList)
}

func (userService *UserService) UpdateUser(user models.User) error {
	return db.DB.Save(&user).Error
}
func (userService *UserService) GetUserById(userId uint) (user models.User) {
	db.DB.First(&user, userId)
	return
}
func (userService *UserService) GetUserByName(userName string) (user models.User) {
	db.DB.Where("name=?", userName).First(&user)
	return
}
func (userService *UserService) GetUserByMobile(mobile string) (user models.User, row int64) {
	row = db.DB.Where("mobile=?", mobile).First(&user).RowsAffected
	return
}
func (userService *UserService) GetUserByEmail(email string) (user models.User) {
	db.DB.Where("email=?", email).First(&user)
	return
}
func (userService *UserService) DeleteUser(userId uint) int64 {
	return db.DB.Delete(&models.User{}, userId).RowsAffected
}
func (userService *UserService) AddUser(user models.User) int64 {
	return db.DB.Create(&user).RowsAffected
}
