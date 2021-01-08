package seeders

import (
	"fmt"
	"gin-admin/app/models"
	"gin-admin/db"
	"gin-admin/pkg/utils"
)

type UserSeeder struct {
}

func NewUserSeeder() *UserSeeder {
	return &UserSeeder{}
}

func (u *UserSeeder) Run() {
	for i := 1; i < 100; i++ {
		user := models.User{}
		user.Name = fmt.Sprintf("user%v", i)
		user.Password = utils.EncodeMD5("123456")
		db.DB.Create(&user)
	}
}

func (u *UserSeeder) Drop() {
	db.DB.Exec("DROP TABLE IF EXISTS users")
}
