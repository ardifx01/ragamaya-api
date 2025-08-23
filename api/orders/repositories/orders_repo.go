package repositories

import (
	"ragamaya-api/models"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data models.Orders) *exceptions.Exception
	FindByUserUUID(ctx *gin.Context, tx *gorm.DB, uuid string) ([]models.Orders, *exceptions.Exception)
	FindByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.Orders, *exceptions.Exception)
	Update(ctx *gin.Context, tx *gorm.DB, data models.Orders) *exceptions.Exception
	LockForUpdateWithTimeout(ctx *gin.Context, tx *gorm.DB, orderUUID string, timeoutSeconds int) *exceptions.Exception
}
