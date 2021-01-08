package seeds

import (
	"gin-admin/seeds/seeders"
)

var seederList = []Seeder{
	seeders.NewAdminSeeder(),
}

type Seeder interface {
	Run()
	Drop()
}

func Run(args []string) {
	switch args[1] {
	case "seed":
		//删除所有表
		dropTables()
		// 重新跑种子
		run()

		break
	}
}

func dropTables() {
	migrate()
}
func run() {
	for _, s := range seederList {
		s.Run()
	}
}
