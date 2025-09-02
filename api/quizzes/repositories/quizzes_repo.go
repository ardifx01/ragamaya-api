package repositories

import (
	"ragamaya-api/api/quizzes/dto"
	"ragamaya-api/models"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data models.Quiz) *exceptions.Exception
	CreateCategory(ctx *gin.Context, tx *gorm.DB, category models.QuizCategory) *exceptions.Exception
	FindCategoryByName(ctx *gin.Context, tx *gorm.DB, name string) (*models.QuizCategory, *exceptions.Exception)
	FindAllCategories(ctx *gin.Context, tx *gorm.DB) ([]models.QuizCategory, *exceptions.Exception)
	Search(ctx *gin.Context, tx *gorm.DB, data dto.SearchReq) ([]models.Quiz, *exceptions.Exception)
	FindBySlug(ctx *gin.Context, tx *gorm.DB, slug string) (*models.Quiz, *exceptions.Exception)
}
