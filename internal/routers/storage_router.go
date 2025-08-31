package routers

import (
	"ragamaya-api/api/storages/controllers"

	"github.com/gin-gonic/gin"
)

func StorageRoutes(r *gin.RouterGroup, compControllers controllers.CompControllers) {
	storageGroup := r.Group("/storage")
	{
		storageGroup.POST("/upload", compControllers.General)
	}
}
