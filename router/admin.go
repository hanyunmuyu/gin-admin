package router

import (
	"gin-admin/app/http/controllers/admins"
	"gin-admin/app/middleware"
)

var (
	adminController = admins.AdminController{}
	loginController = admins.LoginController{}
	userController  = admins.UserController{}
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
	}
}
