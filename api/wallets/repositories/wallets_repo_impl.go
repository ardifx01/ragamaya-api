package repositories

import (
	"ragamaya-api/models"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositoriesImpl struct {
}

func NewComponentRepository() CompRepositories {
	return &CompRepositoriesImpl{}
}

func (r *CompRepositoriesImpl) Create(ctx *gin.Context, tx *gorm.DB, data models.Wallet) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) FindByID(ctx *gin.Context, tx *gorm.DB, id uint) (*models.Wallet, *exceptions.Exception) {
	var data models.Wallet

	err := tx.Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}

	return &data, nil
}

func (r *CompRepositoriesImpl) FindByUserUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.Wallet, *exceptions.Exception) {
	var data models.Wallet

	err := tx.Where("user_uuid = ?", uuid).First(&data).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}

	return &data, nil
}

func (r *CompRepositoriesImpl) UpdateBalance(ctx *gin.Context, tx *gorm.DB, data models.Wallet) *exceptions.Exception {
	result := tx.Where("user_uuid = ?", data.UserUUID).Updates(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}
	return nil
}

func (r *CompRepositoriesImpl) FindTransactionHistoryByUserUUID(ctx *gin.Context, tx *gorm.DB, uuid string) ([]models.WalletTransactionHistory, *exceptions.Exception) {
	var data []models.WalletTransactionHistory

	err := tx.Joins("JOIN wallets ON wallets.id = wallet_transaction_histories.wallet_id").
		Where("wallets.user_uuid = ?", uuid).
		Order("wallet_transaction_histories.created_at DESC").
		Find(&data).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}

	return data, nil
}

func (r *CompRepositoriesImpl) CreateTransaction(ctx *gin.Context, tx *gorm.DB, data models.WalletTransactionHistory) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) CreatePayoutRequest(ctx *gin.Context, tx *gorm.DB, data models.WalletPayoutRequest) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) FindPayoutByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.WalletPayoutRequest, *exceptions.Exception) {
	var data models.WalletPayoutRequest

	err := tx.Where("uuid = ?", uuid).First(&data).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}

	return &data, nil
}

func (r *CompRepositoriesImpl) FindPayoutByWalletID(ctx *gin.Context, tx *gorm.DB, id uint) (*models.WalletPayoutRequest, *exceptions.Exception) {
	var data models.WalletPayoutRequest

	err := tx.Where("wallet_id = ?", id).First(&data).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}

	return &data, nil
}

func (r *CompRepositoriesImpl) UpdatePayout(ctx *gin.Context, tx *gorm.DB, data models.WalletPayoutRequest) *exceptions.Exception {
	result := tx.Where("uuid = ?", data.UUID).Updates(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) CreatePayoutTransactionReceipt(ctx *gin.Context, tx *gorm.DB, data models.WalletPayoutTransactionReceipt) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}