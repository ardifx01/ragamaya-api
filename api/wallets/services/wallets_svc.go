package services

import (
	"ragamaya-api/api/wallets/dto"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompServices interface {
	FindByUserUUID(ctx *gin.Context) (*dto.WalletRes, *exceptions.Exception)
	FindTransactionHistoryByUserUUID(ctx *gin.Context) ([]dto.WalletTransactionRes, *exceptions.Exception)
	CreateTransaction(ctx *gin.Context, input dto.WalletTransactionReq) *exceptions.Exception
	CreateTransactionWithTx(ctx *gin.Context, tx *gorm.DB,  input dto.WalletTransactionReq) *exceptions.Exception 
	RequestPayout(ctx *gin.Context, data dto.WalletPayoutReq) *exceptions.Exception
	ResponsePayout(ctx *gin.Context, data dto.WalletPayoutAcceptReq) *exceptions.Exception
}
