package repositories

import (
	"fmt"
	"ragamaya-api/api/quizzes/dto"
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

func (r *CompRepositoriesImpl) Search(ctx *gin.Context, tx *gorm.DB, data dto.SearchReq) ([]models.Quiz, *exceptions.Exception) {
	var quizzes []models.Quiz

	query := tx.WithContext(ctx).
		Model(&models.Quiz{}).
		Preload("Category")

	if data.Keyword != nil && *data.Keyword != "" {
		kw := fmt.Sprintf("%%%s%%", *data.Keyword)
		query = query.Where(
			tx.Where("title ILIKE ?", kw),
		)
	}

	if data.Category != nil && *data.Category != "" {
		query = query.Joins("JOIN quiz_categories ON quizzes.category_uuid = quiz_categories.uuid").
			Where("LOWER(quiz_categories.name) = LOWER(?)", *data.Category)
	}

	if data.Level != nil {
		query = query.Where("level = ?", *data.Level)
	}

	err := query.
		Order("created_at DESC").
		Find(&quizzes).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return quizzes, nil
}

func (r *CompRepositoriesImpl) FindByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.Quiz, *exceptions.Exception) {
	var quiz models.Quiz
	err := tx.Where("uuid = ?", uuid).
		Preload("Category").
		First(&quiz).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return &quiz, nil
}

func (r *CompRepositoriesImpl) FindBySlug(ctx *gin.Context, tx *gorm.DB, slug string) (*models.Quiz, *exceptions.Exception) {
	var quiz models.Quiz
	err := tx.Where("slug = ?", slug).
		Preload("Category").
		First(&quiz).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return &quiz, nil
}

func (r *CompRepositoriesImpl) CreateCertificate(ctx *gin.Context, tx *gorm.DB, data models.QuizCertificate) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) FindCertificateByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.QuizCertificate, *exceptions.Exception) {
	var certificate models.QuizCertificate
	err := tx.Where("uuid = ?", uuid).
		First(&certificate).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return &certificate, nil
}

func (r *CompRepositoriesImpl) FindCertificateByUserUUID(ctx *gin.Context, tx *gorm.DB, uuid string) ([]models.QuizCertificate, *exceptions.Exception) {
	var certificate []models.QuizCertificate
	err := tx.Where("user_uuid = ?", uuid).
		Find(&certificate).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return certificate, nil
}

func (r *CompRepositoriesImpl) FindCertificateByQuizUUIDandUserUUID(ctx *gin.Context, tx *gorm.DB, quizUUID string, userUUID string) (*models.QuizCertificate, *exceptions.Exception) {
	var certificate models.QuizCertificate
	err := tx.
		Where("quiz_uuid = ?", quizUUID).
		Where("user_uuid = ?", userUUID).
		First(&certificate).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return &certificate, nil
}
