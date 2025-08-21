package services

import (
	"ragamaya-api/midtrans/notifications/dto"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Payment(ctx *gin.Context, data dto.PaymentNotificationReq) *exceptions.Exception
}
