package repositories

import (
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
}
