package routers

import (
	"ragamaya-api/api/sellers/controllers"
	"ragamaya-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func SellerRoutes(r *gin.RouterGroup, compControllers controllers.CompControllers) {
	sellerGroup := r.Group("/seller")
	{
		sellerGroup.GET("/:uuid", compControllers.FindByUUID)
		sellerGroup.POST("/register", middleware.AuthMiddleware(), compControllers.Register)
		sellerGroup.PATCH("/update", middleware.AuthMiddleware(), compControllers.Update)
		sellerGroup.DELETE("/delete", middleware.AuthMiddleware(), compControllers.Delete)
	}
}
