package services

import (
	"fmt"
	"ragamaya-api/api/wallets/dto"
	"ragamaya-api/api/wallets/repositories"
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

func (s *CompServicesImpl) CreateTransaction(ctx *gin.Context, input dto.WalletTransactionReq) *exceptions.Exception {
	validateErr := s.validate.Struct(input)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		return err
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	walletData, err := s.repo.FindByUserUUID(ctx, tx, userData.UUID)
	if err != nil {
		return err
	}

	err = s.repo.CreateTransaction(ctx, tx, models.WalletTransactionHistory{
		UUID:      uuid.NewString(),
		WalletID:  walletData.ID,
		Amount:    input.Amount,
		Type:      models.TransactionType(input.Type),
		Reference: input.Reference,
		Note:      input.Note,
	})
	if err != nil {
		return err
	}

	if input.Type == string(models.Debit) {
		err = s.repo.UpdateBalance(ctx, tx, models.Wallet{
			ID:      walletData.ID,
			Balance: walletData.Balance + input.Amount,
		})
		if err != nil {
			return err
		}
	} else if input.Type == string(models.Credit) {
		err = s.repo.UpdateBalance(ctx, tx, models.Wallet{
			ID:      walletData.ID,
			Balance: walletData.Balance - input.Amount,
		})
		if err != nil {
			return err
		}
	} else {
		tx.Rollback()
		return exceptions.NewValidationException(fmt.Errorf("type must be debit or credit"))
	}

	return nil
}

func (s *CompServicesImpl) CreateTransactionWithTx(ctx *gin.Context, tx *gorm.DB, input dto.WalletTransactionReq) *exceptions.Exception {
	validateErr := s.validate.Struct(input)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		return err
	}

	walletData, err := s.repo.FindByUserUUID(ctx, tx, userData.UUID)
	if err != nil {
		return err
	}

	err = s.repo.CreateTransaction(ctx, tx, models.WalletTransactionHistory{
		UUID:      uuid.NewString(),
		WalletID:  walletData.ID,
		Amount:    input.Amount,
		Type:      models.TransactionType(input.Type),
		Reference: input.Reference,
		Note:      input.Note,
	})
	if err != nil {
		return err
	}

	if input.Type == string(models.Debit) {
		err = s.repo.UpdateBalance(ctx, tx, models.Wallet{
			ID:      walletData.ID,
			Balance: walletData.Balance + input.Amount,
		})
		if err != nil {
			return err
		}
	} else if input.Type == string(models.Credit) {
		err = s.repo.UpdateBalance(ctx, tx, models.Wallet{
			ID:      walletData.ID,
			Balance: walletData.Balance - input.Amount,
		})
		if err != nil {
			return err
		}
	} else {
		tx.Rollback()
		return exceptions.NewValidationException(fmt.Errorf("type must be debit or credit"))
	}

	return nil
}
