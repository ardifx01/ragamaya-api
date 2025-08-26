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

func (r *CompRepositoriesImpl) FindByUserUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.Wallet, *exceptions.Exception) {
	var data models.Wallet

	err := tx.Where("user_uuid = ?", uuid).First(&data).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}

	return &data, nil
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