package routers

import (
	"ragamaya-api/internal/auth/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup, internalAuthController controllers.CompControllers) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", internalAuthController.Login)
	}
}
