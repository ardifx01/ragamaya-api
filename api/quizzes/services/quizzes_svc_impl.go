package services

import (
	"encoding/json"
	"net/http"
	"ragamaya-api/api/quizzes/dto"
	"ragamaya-api/api/quizzes/repositories"
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
		result = append(result, mapper.MapQuizCategoryMTO(v))
	}

	return result, nil
}

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.QuizReq) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	existCategory, err := s.repo.FindCategoryByName(ctx, tx, data.Category)
	if err != nil {
		if err.Status == http.StatusNotFound && existCategory == nil {
			existCategory = &models.QuizCategory{
				UUID: uuid.NewString(),
				Name: data.Category,
			}
			err = s.repo.CreateCategory(ctx, tx, *existCategory)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	input := mapper.MapQuizITM(data)
	input.UUID = uuid.NewString()
	input.Slug = helpers.SlugifyUnique(data.Title)
	input.Questions = helpers.FormatToJSON(data.Questions)
	input.CategoryUUID = existCategory.UUID

	result := s.repo.Create(ctx, tx, input)
	if result != nil {
		return result
	}

	return nil
}

func (s *CompServicesImpl) Search(ctx *gin.Context, data dto.SearchReq) ([]dto.QuizRes, *exceptions.Exception) {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return nil, exceptions.NewValidationException(validateErr)
	}

	result, err := s.repo.Search(ctx, s.DB, data)
	if err != nil {
		return nil, err
	}

	var output []dto.QuizRes

	for _, v := range result {
		item := mapper.MapQuizMTO(v)
		var questions []dto.QuizQuestionRes
		err := json.Unmarshal([]byte(v.Questions), &questions)
		if err != nil {
			return nil, exceptions.NewException(http.StatusInternalServerError, "try again, if still error: report to customer service")
		}

		item.TotalQuestions = len(questions)
		output = append(output, item)
	}

	return output, nil
}
