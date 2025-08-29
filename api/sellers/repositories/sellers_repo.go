package repositories

import (
	"ragamaya-api/api/sellers/dto"
	"ragamaya-api/models"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data models.Sellers) *exceptions.Exception
	FindByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.Sellers, *exceptions.Exception)
	FindByUserUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.Sellers, *exceptions.Exception)
	Update(ctx *gin.Context, tx *gorm.DB, data models.Sellers) *exceptions.Exception
	Delete(ctx *gin.Context, tx *gorm.DB, uuid string) *exceptions.Exception
	FindOrderBySellerUUID(ctx *gin.Context, tx *gorm.DB, uuid string, params dto.OrderQueryParams) ([]models.Orders, *exceptions.Exception) 
}
