package router

import (
	"gin-admin/app/http/controllers/admins"
	"gin-admin/app/middleware"
)

var (
	adminController      = admins.AdminController{}
	loginController      = admins.LoginController{}
	userController       = admins.UserController{}
	permissionController = admins.PermissionController{}
	roleController       = admins.RoleController{}
)

func adminRouter() {
	r := router()
	adminRouter := r.Group("/admin/v1")
	adminRouter.POST("/login", loginController.Login)

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
