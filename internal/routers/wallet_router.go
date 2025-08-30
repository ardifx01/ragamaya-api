package routers

import (
	"ragamaya-api/api/wallets/controllers"

	"github.com/gin-gonic/gin"
)

func WalletRouter(r *gin.RouterGroup, walletController controllers.CompControllers) {
	walletGroup := r.Group("/wallet")
	{
		walletGroup.GET("/payout/getall", walletController.FindAllPayouts)
		walletGroup.POST("/payout/response/:status", walletController.ResponsePayout)
	}
}