package repositories

import (
	"ragamaya-api/models"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	CreateCategory(ctx *gin.Context, tx *gorm.DB, category models.ArticleCategory) *exceptions.Exception
	FindCategoryByName(ctx *gin.Context, tx *gorm.DB, name string) (*models.ArticleCategory, *exceptions.Exception)
	FindAllCategories(ctx *gin.Context, tx *gorm.DB) ([]models.ArticleCategory, *exceptions.Exception)
}
