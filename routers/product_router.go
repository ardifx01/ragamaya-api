package routers

import (
	"ragamaya-api/api/products/controllers"
	"ragamaya-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.RouterGroup, compControllers controllers.CompControllers) {
	productGroup := r.Group("/product")
	{
		productGroup.GET("/search", compControllers.Search)
		productGroup.GET("/owned", middleware.AuthMiddleware(), compControllers.FindProductDigitalOwned)
		productGroup.GET("/:uuid", compControllers.FindByUUID)
		productGroup.POST("/register", middleware.SellerMiddleware(), compControllers.Register)
		productGroup.PUT("/update/:uuid", middleware.SellerMiddleware(), compControllers.Update)
		productGroup.DELETE("/delete/:uuid", middleware.SellerMiddleware(), compControllers.Delete)
		productGroup.DELETE("/delete/thumbnail", middleware.SellerMiddleware(), compControllers.DeleteThumbnail)
	}
}
