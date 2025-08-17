package routers

import (
	"ragamaya-api/api/storages/controllers"

	"github.com/gin-gonic/gin"
)

func StorageRoutes(r *gin.RouterGroup, compControllers controllers.CompControllers) {
	storageGroup := r.Group("/storage")
	{
		imageGroup := storageGroup.Group("/image")
		{
			imageGroup.POST("/upload", compControllers.Images)
			imageGroup.POST("/upload/single", compControllers.Image)
			imageGroup.GET("/getall", compControllers.FindAllImages)
		}
	}
}
