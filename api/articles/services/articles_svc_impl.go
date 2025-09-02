package services

import (
	"net/http"
	"ragamaya-api/api/articles/dto"
	"ragamaya-api/api/articles/repositories"
	"ragamaya-api/models"
	"ragamaya-api/pkg/exceptions"
	"ragamaya-api/pkg/helpers"
	"ragamaya-api/pkg/mapper"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo     repositories.CompRepositories
	DB       *gorm.DB
	validate *validator.Validate
}

func NewComponentServices(compRepositories repositories.CompRepositories, db *gorm.DB, validate *validator.Validate) CompServices {
	return &CompServicesImpl{
		repo:     compRepositories,
		DB:       db,
		validate: validate,
	}
}

func (s *CompServicesImpl) FindAllCategories(ctx *gin.Context) ([]dto.CategoryRes, *exceptions.Exception) {
	data, err := s.repo.FindAllCategories(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	var result []dto.CategoryRes

	for _, v := range data {
		result = append(result, mapper.MapCategoryMTO(v))
	}

	return result, nil
}

func (s *CompServicesImpl) FindOrCreateCategory(ctx *gin.Context, tx *gorm.DB, category string) (*models.ArticleCategory, *exceptions.Exception) {
	existCategory, err := s.repo.FindCategoryByName(ctx, tx, category)
	if err != nil {
		if err.Status == http.StatusNotFound && existCategory == nil {
			existCategory = &models.ArticleCategory{
				UUID: uuid.NewString(),
				Name: category,
			}
			err = s.repo.CreateCategory(ctx, tx, *existCategory)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return existCategory, nil
}

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.ArticleReq) (*dto.ArticleRes, *exceptions.Exception) {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return nil, exceptions.NewValidationException(validateErr)
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	existCategory, err := s.FindOrCreateCategory(ctx, tx, data.Category)
	if err != nil {
		return nil, err
	}

	input := mapper.MapArticleITM(data)
	input.UUID = uuid.NewString()
	input.Slug = helpers.SlugifyUnique(data.Title)
	input.CategoryUUID = existCategory.UUID

	err = s.repo.Create(ctx, tx, input)
	if err != nil {
		return nil, err
	}

	result, err := s.repo.FindByUUID(ctx, tx, input.UUID)
	if err != nil {
		return nil, err
	}

	output := mapper.MapArticleMTO(*result)
	return &output, nil
}

func (s *CompServicesImpl) Search(ctx *gin.Context, data dto.SearchReq) ([]dto.ArticleRes, *exceptions.Exception) {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return nil, exceptions.NewValidationException(validateErr)
	}

	result, err := s.repo.Search(ctx, s.DB, data)
	if err != nil {
		return nil, err
	}

	var output []dto.ArticleRes

	for _, item := range result {
		output = append(output, mapper.MapArticleMTO(item))
	}

	return output, nil
}

func (s *CompServicesImpl) FindBySlug(ctx *gin.Context, slug string) (*dto.ArticleRes, *exceptions.Exception) {
	result, err := s.repo.FindBySlug(ctx, s.DB, slug)
	if err != nil {
		return nil, err
	}

	output := mapper.MapArticleMTO(*result)
	return &output, nil
}

func (s *CompServicesImpl) Update(ctx *gin.Context, data dto.ArticleUpdateReq) (*dto.ArticleRes, *exceptions.Exception) {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return nil, exceptions.NewValidationException(validateErr)
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	existCategory, err := s.FindOrCreateCategory(ctx, tx, data.Category)
	if err != nil {
		return nil, err
	}

	input := mapper.MapArticleUTM(data)
	input.Slug = helpers.SlugifyUnique(data.Title)
	input.CategoryUUID = existCategory.UUID

	err = s.repo.Update(ctx, tx, input)
	if err != nil {
		return nil, err
	}

	result, err := s.repo.FindByUUID(ctx, tx, input.UUID)
	if err != nil {
		return nil, err
	}

	output := mapper.MapArticleMTO(*result)
	return &output, nil
}

func (s *CompServicesImpl) Delete(ctx *gin.Context, uuid string) *exceptions.Exception {
	return s.repo.Delete(ctx, s.DB, uuid)
}
