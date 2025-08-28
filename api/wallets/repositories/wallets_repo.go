package repositories

import (
	"ragamaya-api/models"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data models.Wallet) *exceptions.Exception
	FindByUserUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.Wallet, *exceptions.Exception)
	UpdateBalance(ctx *gin.Context, tx *gorm.DB, data models.Wallet) *exceptions.Exception
	CreateTransaction(ctx *gin.Context, tx *gorm.DB, data models.WalletTransactionHistory) *exceptions.Exception
	FindTransactionHistoryByUserUUID(ctx *gin.Context, tx *gorm.DB, uuid string) ([]models.WalletTransactionHistory, *exceptions.Exception)
	CreatePayoutRequest(ctx *gin.Context, tx *gorm.DB, data models.WalletPayoutRequest) *exceptions.Exception
}
