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
	messageController    = admins.MessageController{}
	configController     = admins.ConfigController{}
	activityController   = admins.ActivityController{}
	productController    = admins.ProductController{}
)

func adminRouter() {
	r := router()
	url := ginSwagger.URL("http://localhost:886/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	adminRouter := r.Group("/admin")
	adminRouter.POST("/login", loginController.Login)
	adminRouter.POST("/config", configController.Init)
	adminRouter.GET("/config/check", configController.Check)
	adminRouter.POST("/upload", uploadController.Upload)
	adminRouter.POST("/upload/multi", uploadController.UploadMulti)
	adminRouter.Use(middleware.AdminJwt())
	{
		adminRouter.GET("/admin/list", adminController.GetAdminList)
		adminRouter.GET("/admin/info", adminController.GetAdminInfo)
		adminRouter.PUT("/admin/update/:adminId", adminController.UpdateAdmin)
		adminRouter.DELETE("/admin/delete/:adminId", adminController.DeleteAdmin)
		adminRouter.GET("/admin/permission", adminController.GetAdminPermissionList)
		adminRouter.GET("/user/list", userController.GetUserList)
		adminRouter.PUT("/user/:userId", userController.UpdateUser)
		adminRouter.DELETE("/user/:userId", userController.DeleteUser)
		adminRouter.GET("/permission/list", permissionController.GetPermissionList)
		adminRouter.GET("/role/list", roleController.GetRoleList)
		adminRouter.GET("/role/detail/:roleId", roleController.GetRoleDetail)
		adminRouter.PUT("/role/update/:roleId", roleController.UpdateRole)
		adminRouter.POST("/role/add", roleController.AddRole)
		adminRouter.GET("/role/all", roleController.GetAllRole)
		adminRouter.DELETE("/role/delete/:roleId", roleController.DeleteRole)
		adminRouter.GET("/message", messageController.GetMessageList)
		adminRouter.GET("/message/:id", messageController.GetMessageDetail)
		adminRouter.GET("/activity", activityController.GetActivityList)

		adminRouter.GET("/product/list", productController.GetProductList)
	}
}
