package repositories

import (
	"ragamaya-api/models"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositoriesImpl struct {
}

func NewComponentRepository() CompRepositories {
	return &CompRepositoriesImpl{}
}

func (r *CompRepositoriesImpl) Create(ctx *gin.Context, tx *gorm.DB, data models.Files) (*models.Files, *exceptions.Exception) {
	result := tx.Create(&data)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return &data, nil
}

func (r *CompRepositoriesImpl) FindAllImages(ctx *gin.Context, tx *gorm.DB) (*[]models.Files, *exceptions.Exception) {
	var data []models.Files

	result := tx.Find(&data).Where("mime_type = ?", "image")
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return &data, nil
}
