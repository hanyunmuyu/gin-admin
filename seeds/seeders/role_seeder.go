package seeders

import (
	"fmt"
	"gin-admin/app/models"
	"gin-admin/db"
)

type RoleSeeder struct {
}

func NewRoleSeeder() *RoleSeeder {
	return &RoleSeeder{}
}

func (r *RoleSeeder) Run() {
	for i := 1; i < 100; i++ {
		role := models.Role{}
		role.RoleName = fmt.Sprintf("role%v", i)
		db.DB.Create(&role)
	}
}

func (r *RoleSeeder) Drop() {
	db.DB.Exec("DROP TABLE IF EXISTS roles")
}
