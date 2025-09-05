package services

import (
	"net/http"
	"ragamaya-api/api/products/dto"
	"ragamaya-api/api/products/repositories"
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

func (s *CompServicesImpl) Update(ctx *gin.Context, uuid string, data dto.ProductUpdateReq) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

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

	productData.Name = data.Name
	productData.Description = data.Description
	productData.Price = int64(data.Price)
	productData.Stock = data.Stock
	productData.Keywords = data.Keywords

	productData.Thumbnails = nil
	for _, thumbnail := range data.Thumbnails {
		productData.Thumbnails = append(productData.Thumbnails, models.ProductThumbnails{
			ProductUUID:  uuid,
			ThumbnailURL: thumbnail.ThumbnailURL,
		})
	}

	result := s.repo.Update(ctx, s.DB, *productData)
	if result != nil {
		return result
	}

	return nil
}

func (s *CompServicesImpl) FindByUUID(ctx *gin.Context, uuid string) (*dto.ProductRes, *exceptions.Exception) {
	userData, _ := helpers.GetUserData(ctx)
	
	var product *models.Products
	var result *exceptions.Exception
	var isOwned bool
	
	if userData.UUID != "" {
		product, result = s.repo.FindByUUIDWithFile(ctx, s.DB, uuid)
		if result != nil {
			return nil, result
		}
		isOwned = s.repo.IsProductDigitalOwned(ctx, s.DB, userData.UUID, uuid)
	} else {
		product, result = s.repo.FindByUUID(ctx, s.DB, uuid)
		if result != nil {
			return nil, result
		}
	}
	
	output := mapper.MapProductMTO(*product)
	output.IsOwned = isOwned
	
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

func (s *CompServicesImpl) DeleteThumbnail(ctx *gin.Context, productUUID string, id uint) *exceptions.Exception {
	productData, err := s.repo.FindByUUID(ctx, s.DB, productUUID)
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

	result := s.repo.DeleteThumbnail(ctx, s.DB, productUUID, id)
	if result != nil {
		return result
	}

	return nil
}

func (s *CompServicesImpl) FindProductDigitalOwned(ctx *gin.Context) ([]dto.ProductRes, *exceptions.Exception) {
	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		return nil, err
	}

	products, exc := s.repo.FindProductDigitalOwned(ctx, s.DB, userData.UUID)
	if exc != nil {
		return nil, exc
	}

	var output []dto.ProductRes
	for _, item := range products {
		output = append(output, mapper.MapProductMTO(item))
	}

	return output, nil
}
