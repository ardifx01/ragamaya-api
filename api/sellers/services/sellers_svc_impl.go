package services

import (
	"ragamaya-api/api/sellers/dto"
	"ragamaya-api/api/sellers/repositories"
	"ragamaya-api/models"
	"ragamaya-api/pkg/exceptions"
	"ragamaya-api/pkg/helpers"
	"ragamaya-api/pkg/mapper"

	userRepositories "ragamaya-api/api/users/repositories"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo     repositories.CompRepositories
	DB       *gorm.DB
	validate *validator.Validate

	userRepo userRepositories.CompRepositories
}

func NewComponentServices(compRepositories repositories.CompRepositories, db *gorm.DB, validate *validator.Validate, userRepo userRepositories.CompRepositories) CompServices {
	return &CompServicesImpl{
		repo:     compRepositories,
		DB:       db,
		validate: validate,
		userRepo: userRepo,
	}
}

func (s *CompServicesImpl) Register(ctx *gin.Context, data dto.RegisterReq) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		return err
	}

	input := mapper.MapSellerITM(data)
	input.UUID = uuid.NewString()
	input.UserUUID = userData.UUID

	if input.AvatarURL == "" {
		input.AvatarURL = userData.AvatarURL
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err = s.repo.Create(ctx, tx, input)
	if err != nil {
		return err
	}

	err = s.userRepo.Update(ctx, tx, models.Users{
		UUID: userData.UUID,
		Role: models.Seller,
	})

	return nil
}

func (s *CompServicesImpl) FindByUUID(ctx *gin.Context, uuid string) (*dto.SellerRes, *exceptions.Exception) {
	result, err := s.repo.FindByUUID(ctx, s.DB, uuid)
	if err != nil {
		return nil, err
	}

	output := mapper.MapSellerMTO(*result)
	return &output, nil
}

func (s *CompServicesImpl) FindByUserUUID(ctx *gin.Context, uuid string) (*dto.SellerRes, *exceptions.Exception) {
	result, err := s.repo.FindByUserUUID(ctx, s.DB, uuid)
	if err != nil {
		return nil, err
	}

	output := mapper.MapSellerMTO(*result)
	return &output, nil
}

func (s *CompServicesImpl) Update(ctx *gin.Context, data dto.UpdateReq) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		return err
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	sellerData, err := s.repo.FindByUserUUID(ctx, tx, userData.UUID)
	if err != nil {
		return err
	}

	input := mapper.MapSellerUTM(data)
	input.UUID = sellerData.UUID

	err = s.repo.Update(ctx, tx, input)
	if err != nil {
		return err
	}

	return nil
}

func (s *CompServicesImpl) Delete(ctx *gin.Context) *exceptions.Exception {
	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		return err
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	sellerData, err := s.repo.FindByUserUUID(ctx, tx, userData.UUID)
	if err != nil {
		return err
	}

	err = s.repo.Delete(ctx, tx, sellerData.UUID)
	if err != nil {
		return err
	}

	err = s.userRepo.Update(ctx, tx, models.Users{
		UUID: userData.UUID,
		Role: models.User,
	})

	return nil
}

func (s *CompServicesImpl) FindOrders(ctx *gin.Context, params dto.OrderQueryParams) ([]dto.OrderRes, *exceptions.Exception) {
	validateErr := s.validate.Struct(params)
	if validateErr != nil {
		return nil, exceptions.NewValidationException(validateErr)
	}

	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		return nil, err
	}

	if userData.SellerProfile.UUID == "" {
		return nil, exceptions.NewException(403, "unauthorized access to seller orders")
	}

	result, err := s.repo.FindOrderBySellerUUID(ctx, s.DB, userData.SellerProfile.UUID, params)
	if err != nil {
		return nil, err
	}

	var output []dto.OrderRes
	for _, order := range result {
		output = append(output, mapper.MapSellerOrderMTO(order))
	}

	return output, nil
}
