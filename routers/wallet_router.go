package routers

import (
	"ragamaya-api/api/wallets/controllers"
	"ragamaya-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func WalletRoutes(r *gin.RouterGroup, walletController controllers.CompControllers) {
	walletGroup := r.Group("/wallet")
	walletGroup.Use(middleware.AuthMiddleware())
	{
		walletGroup.GET("/info", walletController.FindByUserUUID)
		walletGroup.GET("/history", walletController.FindTransactionHistoryByUserUUID)
		walletGroup.POST("/payout", walletController.RequestPayout)
	}
}
