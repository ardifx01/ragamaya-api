package routers

import (
	"ragamaya-api/api/users/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup, userController controllers.CompControllers) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", userController.Login)
		userGroup.POST("/refresh", userController.Refresh)
		userGroup.POST("/logout", userController.Logout)
	}
}
