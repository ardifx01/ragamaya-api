package services

import (
	"ragamaya-api/api/products/dto"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Register(ctx *gin.Context, data dto.RegisterReq) *exceptions.Exception
	FindByUUID(ctx *gin.Context, uuid string) (*dto.ProductRes, *exceptions.Exception)
	Delete(ctx *gin.Context, uuid string) *exceptions.Exception
}
