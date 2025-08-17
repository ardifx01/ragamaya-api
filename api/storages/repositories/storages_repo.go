package repositories

import (
	"ragamaya-api/models"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data models.Files) (*models.Files, *exceptions.Exception)
	FindAllImages(ctx *gin.Context, tx *gorm.DB) (*[]models.Files, *exceptions.Exception)
}
