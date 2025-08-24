package services

import (
	"ragamaya-api/api/payments/dto"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	FindByUserUUID(ctx *gin.Context) (*[]dto.PaymentRes, *exceptions.Exception)
	FindPendingByUserUUID(ctx *gin.Context) (*[]dto.PaymentRes, *exceptions.Exception)
	FindFailedByUserUUID(ctx *gin.Context) (*[]dto.PaymentRes, *exceptions.Exception)
	FindSuccessByUserUUID(ctx *gin.Context) (*[]dto.PaymentRes, *exceptions.Exception)
	FindRefundByUserUUID(ctx *gin.Context) (*[]dto.PaymentRes, *exceptions.Exception)
	FindSuccessByProductUUID(ctx *gin.Context) (*[]dto.PaymentRes, *exceptions.Exception)
}
