package repositories

import (
	"ragamaya-api/api/products/dto"
	"ragamaya-api/models"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data models.Products) *exceptions.Exception
	FindByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.Products, *exceptions.Exception)
	Update(ctx *gin.Context, tx *gorm.DB, data models.Products) *exceptions.Exception
	Delete(ctx *gin.Context, tx *gorm.DB, uuid string) *exceptions.Exception
	Search(ctx *gin.Context, tx *gorm.DB, searchReq dto.ProductSearchReq) ([]models.Products, int64, *exceptions.Exception)
	DecrementStockByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) *exceptions.Exception
	RestoreStockByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) *exceptions.Exception
	CreateProductDigitalOwned(ctx *gin.Context, tx *gorm.DB, data models.ProductDigitalOwned) *exceptions.Exception
	IsProductDigitalOwned(ctx *gin.Context, tx *gorm.DB, userUUID string, productUUID string) bool
}
