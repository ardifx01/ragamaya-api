package routers

import (
	"ragamaya-api/midtrans/notifications/controllers"

	"github.com/gin-gonic/gin"
)

func NotificationRoutes(r *gin.RouterGroup, notificationController controllers.CompControllers) {
	authGroup := r.Group("/notification")
	{
		authGroup.POST("/payment", notificationController.Payment)
	}
}
