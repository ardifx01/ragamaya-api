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
