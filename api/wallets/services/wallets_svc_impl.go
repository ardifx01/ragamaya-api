package services

import (
	"ragamaya-api/api/wallets/dto"
	"ragamaya-api/api/wallets/repositories"
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

func (s *CompServicesImpl) FindByUserUUID(ctx *gin.Context) (*dto.WalletRes, *exceptions.Exception) {
	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		return nil, err
	}

	result, err := s.repo.FindByUserUUID(ctx, s.DB, userData.UUID)
	if err != nil {
		return nil, err
	}

	output := mapper.MapWalletMTO(*result)
	output.Currency = "IDR"
	return &output, nil
}

func (s *CompServicesImpl) FindTransactionHistoryByUserUUID(ctx *gin.Context) ([]dto.WalletTransactionRes, *exceptions.Exception) {
	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		return nil, err
	}

	result, err := s.repo.FindTransactionHistoryByUserUUID(ctx, s.DB, userData.UUID)
	if err != nil {
		return nil, err
	}

	var output []dto.WalletTransactionRes
	for _, data := range result {
		output = append(output, mapper.MapWalletTransactionMTO(data))
	}

	return output, nil
}
