package router

import (
	"gin-admin/app/http/controllers/admins"
	"gin-admin/app/middleware"
	_ "gin-admin/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	adminController      = admins.AdminController{}
	loginController      = admins.LoginController{}
	userController       = admins.UserController{}
	permissionController = admins.PermissionController{}
	roleController       = admins.RoleController{}
	uploadController     = admins.UploadController{}
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
func adminRouter() {
	r := router()

	url := ginSwagger.URL("http://localhost:886/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	adminRouter := r.Group("/admin/v1")
	adminRouter.POST("/login", loginController.Login)
	adminRouter.POST("/upload", uploadController.Upload)
	adminRouter.POST("/upload/multi", uploadController.UploadMulti)
	adminRouter.Use(middleware.AdminJwt())
	{
		adminRouter.GET("/admin/list", adminController.GetAdminList)
		adminRouter.GET("/admin/update", adminController.UpdateAdmin)
		adminRouter.GET("/user/list", userController.GetUserList)
		adminRouter.PUT("/user/:userId", userController.UpdateUser)
		adminRouter.DELETE("/user/:userId", userController.DeleteUser)
		adminRouter.GET("/permission/list", permissionController.GetPermissionList)
		adminRouter.GET("/role/list", roleController.GetRoleList)
		adminRouter.GET("/role/detail/:roleId", roleController.GetRoleDetail)
		adminRouter.PUT("/role/update/:roleId", roleController.UpdateRole)
		adminRouter.DELETE("/role/delete/:roleId", roleController.DeleteRole)
	}
}
