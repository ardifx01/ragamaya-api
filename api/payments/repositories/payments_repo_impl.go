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

func (r *CompRepositoriesImpl) Create(ctx *gin.Context, tx *gorm.DB, data models.Payments) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) FindAll(ctx *gin.Context, tx *gorm.DB) (*[]models.Payments, *exceptions.Exception) {
	var payments []models.Payments
	result := tx.
		Preload("Product").
		Preload("Product.Thumbnails").
		Order("created_at DESC").
		Find(&payments)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return &payments, nil
}

func (r *CompRepositoriesImpl) FindByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.Payments, *exceptions.Exception) {
	var payment models.Payments

	result := tx.
		Preload("PaymentActions").
		Preload("PaymentVANumbers").
		Preload("Product").
		Preload("Product.Thumbnails").
		Where("uuid = ?", uuid).
		Order("created_at DESC").
		First(&payment)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return &payment, nil
}

func (r *CompRepositoriesImpl) FindByUserUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*[]models.Payments, *exceptions.Exception) {
	var payments []models.Payments

	result := tx.
		Preload("PaymentActions").
		Preload("PaymentVANumbers").
		Preload("Product").
		Preload("Product.Thumbnails").
		Where("user_uuid = ?", uuid).
		Order("created_at DESC").
		Find(&payments)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return &payments, nil
}

func (r *CompRepositoriesImpl) FindPendingByUserUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*[]models.Payments, *exceptions.Exception) {
	var payments []models.Payments

	result := tx.
		Preload("PaymentActions").
		Preload("PaymentVANumbers").
		Preload("Product").
		Preload("Product.Thumbnails").
		Where("user_uuid = ?", uuid).
		Where("transaction_status = ?", "pending").
		Order("created_at DESC").
		Find(&payments)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return &payments, nil
}

func (r *CompRepositoriesImpl) FindFailedByUserUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*[]models.Payments, *exceptions.Exception) {
	var payments []models.Payments

	result := tx.
		Preload("Product").
		Preload("Product.Thumbnails").
		Where("user_uuid = ?", uuid).
		Where("transaction_status = ? OR transaction_status = ? OR transaction_status = ? OR transaction_status = ?", "expire", "deny", "cancel", "failure").
		Order("created_at DESC").
		Find(&payments)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return &payments, nil
}

func (r *CompRepositoriesImpl) FindSuccessByUserUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*[]models.Payments, *exceptions.Exception) {
	var payments []models.Payments

	result := tx.
		Preload("Product").
		Preload("Product.Thumbnails").
		Where("user_uuid = ?", uuid).
		Where("transaction_status = ? OR transaction_status = ?", "capture", "settlement").
		Order("created_at DESC").
		Find(&payments)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return &payments, nil
}

func (r *CompRepositoriesImpl) FindRefundByUserUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*[]models.Payments, *exceptions.Exception) {
	var payments []models.Payments

	result := tx.
		Preload("Product").
		Preload("Product.Thumbnails").
		Where("user_uuid = ?", uuid).
		Where("transaction_status = ? OR transaction_status = ?", "refund", "partial_refund").
		Order("created_at DESC").
		Find(&payments)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return &payments, nil
}

func (r *CompRepositoriesImpl) FindByOrderUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*[]models.Payments, *exceptions.Exception) {
	var payments []models.Payments

	result := tx.
		Preload("PaymentActions").
		Preload("PaymentVANumbers").
		Preload("Product").
		Preload("Product.Thumbnails").
		Where("order_uuid = ?", uuid).
		Order("created_at DESC").
		Find(&payments)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return &payments, nil
}

func (r *CompRepositoriesImpl) FindByProductUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*[]models.Payments, *exceptions.Exception) {
	var payments []models.Payments

	result := tx.
		Preload("Product").
		Preload("Product.Thumbnails").
		Where("product_uuid = ?", uuid).
		Order("created_at DESC").
		Find(&payments)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return &payments, nil
}

func (r *CompRepositoriesImpl) FindSuccessByProductUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*[]models.Payments, *exceptions.Exception) {
	var payments []models.Payments

	result := tx.
		Preload("Product").
		Preload("Product.Thumbnails").
		Preload("Order").
		Where("product_uuid = ?", uuid).
		Where("transaction_status = ? OR transaction_status = ?", "capture", "settlement").
		Order("created_at DESC").
		Find(&payments)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return &payments, nil
}

func (r *CompRepositoriesImpl) Update(ctx *gin.Context, tx *gorm.DB, data models.Payments) *exceptions.Exception {
	result := tx.Where("uuid = ?", data.UUID).Updates(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}
