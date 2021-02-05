package seeds

import (
	"gin-admin/seeds/seeders"
)

var seederList = []Seeder{
	seeders.NewPermissionSeeder(),
	seeders.NewRoleSeeder(),
	seeders.NewAdminSeeder(),
	seeders.NewUserSeeder(),
	seeders.NewMessageSeeder(),
	seeders.NewActivitySeeder(),
	seeders.NewProductSeeder(),
}

type Seeder interface {
	Run()
	Drop()
}

func Run(args []string) {
	switch args[1] {
	case "seed":
		//删除所有表
		drop()
		migrate()
		// 重新跑种子
		run()
		break
	}
}
func drop() {
	for _, s := range seederList {
		s.Drop()
	}
}
func run() {
	for _, s := range seederList {
		s.Run()
	}
}
