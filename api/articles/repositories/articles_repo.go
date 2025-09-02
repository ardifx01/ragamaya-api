package repositories

import (
	"ragamaya-api/api/articles/dto"
	"ragamaya-api/models"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	CreateCategory(ctx *gin.Context, tx *gorm.DB, category models.ArticleCategory) *exceptions.Exception
	FindCategoryByName(ctx *gin.Context, tx *gorm.DB, name string) (*models.ArticleCategory, *exceptions.Exception)
	FindAllCategories(ctx *gin.Context, tx *gorm.DB) ([]models.ArticleCategory, *exceptions.Exception)
	Create(ctx *gin.Context, tx *gorm.DB, article models.Article) *exceptions.Exception
	FindBySlug(ctx *gin.Context, tx *gorm.DB, slug string) (*models.Article, *exceptions.Exception)
	FindByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.Article, *exceptions.Exception)
	Search(ctx *gin.Context, tx *gorm.DB, data dto.SearchReq) ([]models.Article, *exceptions.Exception)
	Update(ctx *gin.Context, tx *gorm.DB, data models.Article) *exceptions.Exception
	Delete(ctx *gin.Context, tx *gorm.DB, uuid string) *exceptions.Exception
}
