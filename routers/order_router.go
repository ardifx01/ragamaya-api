package routers

import (
	"ragamaya-api/api/orders/controllers"
	"ragamaya-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.RouterGroup, orderController controllers.CompControllers) {
	orderGroup := r.Group("/order")
	orderGroup.Use(middleware.AuthMiddleware())
	{
		orderGroup.POST("/create/:payment_type", orderController.Create)
		orderGroup.GET("/stream", orderController.StreamInfo)
		// orderGroup.GET("/get", orderController.FindByUUID)
		// orderGroup.GET("/getall", orderController.FindByUserUUID)
	}
}
