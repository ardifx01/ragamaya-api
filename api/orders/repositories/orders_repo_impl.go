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

func (r *CompRepositoriesImpl) Create(ctx *gin.Context, tx *gorm.DB, data models.Orders) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) FindByUserUUID(ctx *gin.Context, tx *gorm.DB, uuid string) ([]models.Orders, *exceptions.Exception) {
	var order []models.Orders
	err := tx.
		Where("user_uuid = ?", uuid).Find(&order).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return order, nil
}

func (r *CompRepositoriesImpl) FindByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.Orders, *exceptions.Exception) {
	var order models.Orders

	result := tx.
		Preload("Payments").
		Preload("Payments.PaymentActions").
		Preload("Payments.PaymentVANumbers").
		Where("uuid = ?", uuid).
		Order("created_at DESC").
		First(&order)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return &order, nil
}

func (r *CompRepositoriesImpl) Update(ctx *gin.Context, tx *gorm.DB, data models.Orders) *exceptions.Exception {
	result := tx.Where("uuid = ?", data.UUID).Updates(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}
	return nil
}
