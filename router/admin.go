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

func adminRouter() {
	r := router()
	url := ginSwagger.URL("http://localhost:886/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	adminRouter := r.Group("/admin")
	adminRouter.POST("/login", loginController.Login)
	adminRouter.POST("/upload", uploadController.Upload)
	adminRouter.POST("/upload/multi", uploadController.UploadMulti)
	adminRouter.Use(middleware.AdminJwt())
	{
		adminRouter.GET("/admin/list", adminController.GetAdminList)
		adminRouter.GET("/admin/info", adminController.GetAdminInfo)
		adminRouter.GET("/admin/update", adminController.UpdateAdmin)
		adminRouter.GET("/admin/permission", adminController.GetAdminPermissionList)
		adminRouter.GET("/user/list", userController.GetUserList)
		adminRouter.PUT("/user/:userId", userController.UpdateUser)
		adminRouter.DELETE("/user/:userId", userController.DeleteUser)
		adminRouter.GET("/permission/list", permissionController.GetPermissionList)
		adminRouter.GET("/role/list", roleController.GetRoleList)
		adminRouter.GET("/role/detail/:roleId", roleController.GetRoleDetail)
		adminRouter.PUT("/role/update/:roleId", roleController.UpdateRole)
		adminRouter.POST("/role/add", roleController.AddRole)
		adminRouter.DELETE("/role/delete/:roleId", roleController.DeleteRole)
	}
}
