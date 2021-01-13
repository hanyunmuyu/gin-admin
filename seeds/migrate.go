package seeds

import (
	"gin-admin/app/models"
	"gin-admin/db"
)

func migrate() {
	_ = db.DB.AutoMigrate(
		models.Role{},
		models.Permission{},
		models.RolePermission{},
		models.Admin{},
		models.User{},
	)
}
