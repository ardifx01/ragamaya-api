package services

import (
	"ragamaya-api/api/articles/dto"
	"ragamaya-api/api/articles/repositories"
	"ragamaya-api/pkg/exceptions"
	"ragamaya-api/pkg/mapper"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
