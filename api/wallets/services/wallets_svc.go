package services

import (
	"ragamaya-api/api/wallets/dto"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	FindByUserUUID(ctx *gin.Context) (*dto.WalletRes, *exceptions.Exception)
	FindTransactionHistoryByUserUUID(ctx *gin.Context) ([]dto.WalletTransactionRes, *exceptions.Exception)
}
