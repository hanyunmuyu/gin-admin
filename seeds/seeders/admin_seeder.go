package seeders

import (
	"fmt"
	"gin-admin/app/models"
	"gin-admin/db"
	"gin-admin/pkg/utils"
)

type AdminSeeder struct {
}

func NewAdminSeeder() *AdminSeeder {
	return &AdminSeeder{}
}

func (a *AdminSeeder) Run() {
	for i := 1; i <= 100; i++ {
		admin := models.Admin{}
		admin.Name = fmt.Sprintf("admin%v", i)
		admin.Password = utils.EncodeMD5("123456")
		db.DB.Create(&admin)
	}
}

func (a *AdminSeeder) Drop() {
	db.DB.Exec("DROP TABLE IF EXISTS admins")
}
