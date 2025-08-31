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

func (r *CompRepositoriesImpl) CreateCategory(ctx *gin.Context, tx *gorm.DB, category models.ArticleCategory) *exceptions.Exception {
	result := tx.Create(&category)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}
	return nil
}

func (r *CompRepositoriesImpl) FindCategoryByName(ctx *gin.Context, tx *gorm.DB, name string) (*models.ArticleCategory, *exceptions.Exception) {
	var category models.ArticleCategory
	err := tx.Where("LOWER(name) = LOWER(?)", name).First(&category).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return &category, nil
}

func (r *CompRepositoriesImpl) FindAllCategories(ctx *gin.Context, tx *gorm.DB) ([]models.ArticleCategory, *exceptions.Exception) {
	var categories []models.ArticleCategory
	err := tx.Find(&categories).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return categories, nil
}

func (r *CompRepositoriesImpl) Create(ctx *gin.Context, tx *gorm.DB, article models.Article) *exceptions.Exception {
	result := tx.Create(&article)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}
	return nil
}

func (r *CompRepositoriesImpl) FindBySlug(ctx *gin.Context, tx *gorm.DB, slug string) (*models.Article, *exceptions.Exception) {
	var article models.Article
	err := tx.Where("slug = ?", slug).
		Preload("Category").
		First(&article).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return &article, nil
}

func (r *CompRepositoriesImpl) FindByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.Article, *exceptions.Exception) {
	var article models.Article
	err := tx.Where("uuid = ?", uuid).
		Preload("Category").
		First(&article).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return &article, nil
}
