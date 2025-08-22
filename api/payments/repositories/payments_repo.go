package repositories

import (
	"ragamaya-api/models"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data models.Payments) *exceptions.Exception
	FindAll(ctx *gin.Context, tx *gorm.DB) (*[]models.Payments, *exceptions.Exception)
	FindByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.Payments, *exceptions.Exception)
	FindByUserUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*[]models.Payments, *exceptions.Exception)
	FindPendingByUserUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*[]models.Payments, *exceptions.Exception)
	FindFailedByUserUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*[]models.Payments, *exceptions.Exception)
	FindSuccessByUserUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*[]models.Payments, *exceptions.Exception)
	FindRefundByUserUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*[]models.Payments, *exceptions.Exception)
	FindByOrderUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*[]models.Payments, *exceptions.Exception)
	FindByProductUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*[]models.Payments, *exceptions.Exception)
	FindSuccessByProductUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*[]models.Payments, *exceptions.Exception) 
	Update(ctx *gin.Context, tx *gorm.DB, data models.Payments) *exceptions.Exception
}