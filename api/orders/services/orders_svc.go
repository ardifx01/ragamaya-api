package services

import (
	"ragamaya-api/api/orders/dto"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Create(ctx *gin.Context, data dto.OrderReq) (*dto.OrderChargeRes, *exceptions.Exception)
	RemoveStreamClient(ctx *gin.Context, orderUUID string, client dto.StreamClient)
	SendStreamEvent(ctx *gin.Context, orderUUID string, data dto.OrderStreamRes)
	FindByUUID(ctx *gin.Context, uuid string) (*dto.OrderRes, *exceptions.Exception)
	FindByUserUUID(ctx *gin.Context) ([]dto.OrderRes, *exceptions.Exception)
}
