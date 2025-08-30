package routers

import (
	"ragamaya-api/api/storages/controllers"
	"ragamaya-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func StorageRoutes(r *gin.RouterGroup, compControllers controllers.CompControllers) {
	storageGroup := r.Group("/storage")
	storageGroup.Use(middleware.AuthMiddleware())
	{
		imageGroup := storageGroup.Group("/image")
		{
			imageGroup.POST("/upload", compControllers.Images)
			imageGroup.POST("/upload/single", compControllers.Image)
		}

		generalGroup := storageGroup.Group("/general")
		generalGroup.Use(middleware.SellerMiddleware())
		{
			generalGroup.POST("/upload", compControllers.General)
		}
	}
}
