package main

import (
	"gin-admin/router"
	"gin-admin/seeds"
	"os"
)

// @title gin-admin
// @version 1.0
// @description Gin框架实现的内容分享系统
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Security ApiKeyAuth
// @contact.name hanyun
// @contact.url http://xiangshike.com
// @contact.email 1355081829@qq.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:886
func main() {
	if len(os.Args) > 1 {
		seeds.Run(os.Args)
	} else {
		router.Run()
	}
}
