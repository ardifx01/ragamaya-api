package services

import (
	"net/http"
	"ragamaya-api/api/products/dto"
	"ragamaya-api/api/products/repositories"
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

func (s *CompServicesImpl) Register(ctx *gin.Context, data dto.RegisterReq) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	sellerData, err := helpers.GetUserData(ctx)
	if err != nil {
		return err
	}

	input := mapper.MapProductITM(data)
	input.UUID = uuid.NewString()
	input.SellerUUID = sellerData.SellerProfile.UUID

	result := s.repo.Create(ctx, s.DB, input)
	if result != nil {
		return result
	}

	return nil
}

func (s *CompServicesImpl) FindByUUID(ctx *gin.Context, uuid string) (*dto.ProductRes, *exceptions.Exception) {
	product, result := s.repo.FindByUUID(ctx, s.DB, uuid)
	if result != nil {
		return nil, result
	}

	output := mapper.MapProductMTO(*product)
	return &output, nil
}

func (s *CompServicesImpl) Delete(ctx *gin.Context, uuid string) *exceptions.Exception {
	productData, err := s.repo.FindByUUID(ctx, s.DB, uuid)
	if err != nil {
		return err
	}

	sellerData, err := helpers.GetUserData(ctx)
	if err != nil {
		return err
	}

	if productData.SellerUUID != sellerData.SellerProfile.UUID {
		return exceptions.NewException(http.StatusForbidden, exceptions.ErrNotTheOwner)
	}

	result := s.repo.Delete(ctx, s.DB, uuid)
	if result != nil {
		return result
	}

	return nil
}

func (s *CompServicesImpl) Search(ctx *gin.Context, data dto.ProductSearchReq) ([]dto.ProductRes, *exceptions.Exception) {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return nil, exceptions.NewValidationException(validateErr)
	}

	product, _, err := s.repo.Search(ctx, s.DB, data)
	if err != nil {
		return nil, err
	}

	var output []dto.ProductRes

	for _, item := range product {
		output = append(output, mapper.MapProductMTO(item))
	}

	return output, nil
}
