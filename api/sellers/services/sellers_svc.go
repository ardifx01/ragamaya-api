package services

import (
	"ragamaya-api/api/sellers/dto"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Register(ctx *gin.Context, data dto.RegisterReq) *exceptions.Exception
	FindByUUID(ctx *gin.Context, uuid string) (*dto.SellerRes, *exceptions.Exception)
	FindByUserUUID(ctx *gin.Context, uuid string) (*dto.SellerRes, *exceptions.Exception)
	Update(ctx *gin.Context, data dto.UpdateReq) *exceptions.Exception
	Delete(ctx *gin.Context) *exceptions.Exception
	FindOrders(ctx *gin.Context, params dto.OrderQueryParams) ([]dto.OrderRes, *exceptions.Exception) 
}
