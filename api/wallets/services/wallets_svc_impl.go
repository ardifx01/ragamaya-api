package services

import (
	"fmt"
	"ragamaya-api/api/wallets/dto"
	"ragamaya-api/api/wallets/repositories"
	"ragamaya-api/models"
	"ragamaya-api/pkg/exceptions"
	"ragamaya-api/pkg/helpers"
	"ragamaya-api/pkg/logger"
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

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	walletData, err := s.repo.FindByUserUUID(ctx, tx, input.UserUUID)
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
			UserUUID: walletData.UserUUID,
			Balance:  walletData.Balance + input.Amount,
		})
		if err != nil {
			return err
		}
	} else if input.Type == string(models.Credit) {
		err = s.repo.UpdateBalance(ctx, tx, models.Wallet{
			UserUUID: walletData.UserUUID,
			Balance:  walletData.Balance - input.Amount,
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
	logger.Info("CreateTransactionWithTx called")
	validateErr := s.validate.Struct(input)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	walletData, err := s.repo.FindByUserUUID(ctx, tx, input.UserUUID)
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
			UserUUID: walletData.UserUUID,
			Balance:  walletData.Balance + input.Amount,
		})
		if err != nil {
			return err
		}
	} else if input.Type == string(models.Credit) {
		err = s.repo.UpdateBalance(ctx, tx, models.Wallet{
			UserUUID: walletData.UserUUID,
			Balance:  walletData.Balance - input.Amount,
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

func (s *CompServicesImpl) RequestPayout(ctx *gin.Context, data dto.WalletPayoutReq) *exceptions.Exception {
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

	walletData, err := s.repo.FindByUserUUID(ctx, tx, userData.UUID)
	if err != nil {
		return err
	}

	if walletData.Balance < data.Amount {
		return exceptions.NewValidationException(fmt.Errorf("insufficient balance"))
	}

	input := mapper.MapPayoutITM(data)
	input.UUID = uuid.NewString()
	input.WalletID = walletData.ID

	err = s.repo.CreateTransaction(ctx, tx, models.WalletTransactionHistory{
		UUID:      uuid.NewString(),
		WalletID:  walletData.ID,
		Amount:    input.Amount,
		Type:      models.Credit,
		Reference: "Payout " + input.UUID,
		Note:      "Payout request",
	})
	if err != nil {
		return err
	}

	err = s.repo.UpdateBalance(ctx, tx, models.Wallet{
		UserUUID: walletData.UserUUID,
		Balance:  walletData.Balance - input.Amount,
	})
	if err != nil {
		return err
	}

	err = s.repo.CreatePayoutRequest(ctx, tx, input)
	if err != nil {
		return err
	}

	return nil
}

func (s *CompServicesImpl) FindAllPayouts(ctx *gin.Context) ([]dto.WalletPayoutRes, *exceptions.Exception) {
	result, err := s.repo.FindAllPayouts(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	var output []dto.WalletPayoutRes
	for _, v := range result {
		data := mapper.MapPayoutMTO(v)
		data.User.UUID = v.Wallet.User.UUID
		data.User.Email = v.Wallet.User.Email
		data.User.Name = v.Wallet.User.Name
		data.User.AvatarURL = v.Wallet.User.AvatarURL
		output = append(output, data)
	}

	return output, nil
}

func (s *CompServicesImpl) ResponsePayout(ctx *gin.Context, data dto.WalletPayoutAcceptReq) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	payoutData, err := s.repo.FindPayoutByUUID(ctx, tx, data.PayoutUUID)
	if err != nil {
		return err
	}

	if payoutData.Status != models.Pending {
		return exceptions.NewValidationException(fmt.Errorf("payout already processed"))
	}

	switch data.Status {
	case string(models.Completed):
		payoutData.Status = models.Completed
		err = s.repo.UpdatePayout(ctx, tx, *payoutData)
		if err != nil {
			return err
		}

		err = s.repo.CreatePayoutTransactionReceipt(ctx, tx, models.WalletPayoutTransactionReceipt{
			PayoutUUID: payoutData.UUID,
			ReceiptURL: data.ReceiptURL,
			Note:       data.Note,
		})
		if err != nil {
			return err
		}

	case string(models.Failed):
		payoutData.Status = models.Failed
		err = s.repo.UpdatePayout(ctx, tx, *payoutData)
		if err != nil {
			return err
		}

		err = s.repo.CreatePayoutTransactionReceipt(ctx, tx, models.WalletPayoutTransactionReceipt{
			PayoutUUID: payoutData.UUID,
			ReceiptURL: data.ReceiptURL,
			Note:       data.Note,
		})
		if err != nil {
			return err
		}

		walletData, err := s.repo.FindByID(ctx, tx, payoutData.WalletID)
		if err != nil {
			return err
		}

		err = s.repo.CreateTransaction(ctx, tx, models.WalletTransactionHistory{
			UUID:      uuid.NewString(),
			WalletID:  walletData.ID,
			Amount:    payoutData.Amount,
			Type:      models.Debit,
			Reference: "Payout Rejected " + payoutData.UUID,
			Note:      "Payout rejected, refund to wallet",
		})
		if err != nil {
			return err
		}

		err = s.repo.UpdateBalance(ctx, tx, models.Wallet{
			UserUUID: walletData.UserUUID,
			Balance:  walletData.Balance + payoutData.Amount,
		})
		if err != nil {
			return err
		}

	default:
		tx.Rollback()
		return exceptions.NewValidationException(fmt.Errorf("invalid status"))
	}

	return nil
}

func (s *CompServicesImpl) FindPayoutsByUserUUID(ctx *gin.Context) ([]dto.WalletPayoutRes, *exceptions.Exception) {
	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		return nil, err
	}

	result, err := s.repo.FindPayoutByUserUUID(ctx, s.DB, userData.UUID)
	if err != nil {
		return nil, err
	}

	var output []dto.WalletPayoutRes
	for _, data := range result {
		output = append(output, mapper.MapPayoutMTO(data))
	}

	return output, nil
}
