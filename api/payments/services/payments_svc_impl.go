package services

import (
	"net/http"
	"ragamaya-api/api/payments/dto"
	"ragamaya-api/api/payments/repositories"
	"ragamaya-api/pkg/exceptions"
	"ragamaya-api/pkg/helpers"
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


func (s *CompServicesImpl) FindByUserUUID(ctx *gin.Context) (*[]dto.PaymentRes, *exceptions.Exception) {
	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		return nil, err
	}

	result, err := s.repo.FindByUserUUID(ctx, s.DB, userData.UUID)
	if err != nil {
		return nil, err
	}

	var output []dto.PaymentRes
	for _, data := range *result {
		output = append(output, mapper.MapPaymentMTO(data))
	}

	return &output, nil
}

func (s *CompServicesImpl) FindPendingByUserUUID(ctx *gin.Context) (*[]dto.PaymentRes, *exceptions.Exception) {
	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		return nil, err
	}

	result, err := s.repo.FindPendingByUserUUID(ctx, s.DB, userData.UUID)
	if err != nil {
		return nil, err
	}

	var output []dto.PaymentRes
	for _, data := range *result {
		output = append(output, mapper.MapPaymentMTO(data))
	}

	return &output, nil
}

func (s *CompServicesImpl) FindFailedByUserUUID(ctx *gin.Context) (*[]dto.PaymentRes, *exceptions.Exception) {
	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		return nil, err
	}

	result, err := s.repo.FindFailedByUserUUID(ctx, s.DB, userData.UUID)
	if err != nil {
		return nil, err
	}

	var output []dto.PaymentRes
	for _, data := range *result {
		output = append(output, mapper.MapPaymentMTO(data))
	}

	return &output, nil
}

func (s *CompServicesImpl) FindSuccessByUserUUID(ctx *gin.Context) (*[]dto.PaymentRes, *exceptions.Exception) {
	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		return nil, err
	}

	result, err := s.repo.FindSuccessByUserUUID(ctx, s.DB, userData.UUID)
	if err != nil {
		return nil, err
	}

	var output []dto.PaymentRes
	for _, data := range *result {
		output = append(output, mapper.MapPaymentMTO(data))
	}

	return &output, nil
}

func (s *CompServicesImpl) FindRefundByUserUUID(ctx *gin.Context) (*[]dto.PaymentRes, *exceptions.Exception) {
	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		return nil, err
	}

	result, err := s.repo.FindRefundByUserUUID(ctx, s.DB, userData.UUID)
	if err != nil {
		return nil, err
	}

	var output []dto.PaymentRes
	for _, data := range *result {
		output = append(output, mapper.MapPaymentMTO(data))
	}

	return &output, nil
}

func (s *CompServicesImpl) FindSuccessByProductUUID(ctx *gin.Context) (*[]dto.PaymentRes, *exceptions.Exception) {
	productUUID := ctx.Query("product_uuid")
	if productUUID == "" {
		return nil, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest)
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	payments, err := s.repo.FindSuccessByProductUUID(ctx, tx, productUUID)
	if err != nil {
		return nil, err
	}

	var output []dto.PaymentRes
	for _, data := range *payments {
		result := mapper.MapPaymentMTO(data)
		output = append(output, result)
	}

	return &output, nil
}
