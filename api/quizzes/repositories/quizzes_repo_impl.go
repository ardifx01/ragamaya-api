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

func (r *CompRepositoriesImpl) CreateCategory(ctx *gin.Context, tx *gorm.DB, category models.QuizCategory) *exceptions.Exception {
	result := tx.Create(&category)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}
	return nil
}

func (r *CompRepositoriesImpl) FindCategoryByName(ctx *gin.Context, tx *gorm.DB, name string) (*models.QuizCategory, *exceptions.Exception) {
	var category models.QuizCategory
	err := tx.Where("LOWER(name) = LOWER(?)", name).First(&category).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return &category, nil
}

func (r *CompRepositoriesImpl) FindAllCategories(ctx *gin.Context, tx *gorm.DB) ([]models.QuizCategory, *exceptions.Exception) {
	var categories []models.QuizCategory
	err := tx.Find(&categories).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return categories, nil
}

func (r *CompRepositoriesImpl) Create(ctx *gin.Context, tx *gorm.DB, data models.Quiz) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}
