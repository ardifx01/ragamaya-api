package routers

import (
	"ragamaya-api/api/products/controllers"
	"ragamaya-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.RouterGroup, compControllers controllers.CompControllers) {
	productGroup := r.Group("/product")
	{
		productGroup.GET("/:uuid", compControllers.FindByUUID)
		productGroup.POST("/register", middleware.SellerMiddleware(), compControllers.Register)
		productGroup.DELETE("/delete/:uuid", middleware.SellerMiddleware(), compControllers.Delete)
	}
}
