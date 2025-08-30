package services

import (
	"ragamaya-api/api/products/dto"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Register(ctx *gin.Context, data dto.RegisterReq) *exceptions.Exception
	FindByUUID(ctx *gin.Context, uuid string) (*dto.ProductRes, *exceptions.Exception)
	Update(ctx *gin.Context, uuid string, data dto.ProductUpdateReq) *exceptions.Exception
	Delete(ctx *gin.Context, uuid string) *exceptions.Exception
	Search(ctx *gin.Context, data dto.ProductSearchReq) ([]dto.ProductRes, *exceptions.Exception)
	DeleteThumbnail(ctx *gin.Context, productUUID string, id uint) *exceptions.Exception
	FindProductDigitalOwned(ctx *gin.Context) ([]dto.ProductRes, *exceptions.Exception)
}
