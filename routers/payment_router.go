package routers

import (
	"ragamaya-api/api/payments/controllers"
	"ragamaya-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.RouterGroup, paymentController controllers.CompControllers) {
	paymentGroup := r.Group("/payment")
	paymentGroup.Use(middleware.AuthMiddleware())
	{
		paymentGroup.GET("/getall", paymentController.FindByUserUUID)
		paymentGroup.GET("/pending", paymentController.FindPendingByUserUUID)
		paymentGroup.GET("/failed", paymentController.FindFailedByUserUUID)
		paymentGroup.GET("/success", paymentController.FindSuccessByUserUUID)
		paymentGroup.GET("/refund", paymentController.FindRefundByUserUUID)
		// paymentGroup.POST("/free", paymentController.PayFreeTransaction)
	}
}
