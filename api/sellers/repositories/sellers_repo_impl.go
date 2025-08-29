package repositories

import (
	"ragamaya-api/api/sellers/dto"
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

func (r *CompRepositoriesImpl) Create(ctx *gin.Context, tx *gorm.DB, data models.Sellers) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) FindByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.Sellers, *exceptions.Exception) {
	var seller models.Sellers
	err := tx.Where("uuid = ?", uuid).First(&seller).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return &seller, nil
}

func (r *CompRepositoriesImpl) FindByUserUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.Sellers, *exceptions.Exception) {
	var seller models.Sellers
	err := tx.Where("user_uuid = ?", uuid).First(&seller).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return &seller, nil
}

func (r *CompRepositoriesImpl) Update(ctx *gin.Context, tx *gorm.DB, data models.Sellers) *exceptions.Exception {
	result := tx.Where("uuid = ?", data.UUID).Updates(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}
	return nil
}

func (r *CompRepositoriesImpl) Delete(ctx *gin.Context, tx *gorm.DB, uuid string) *exceptions.Exception {
	err := tx.Where("uuid = ?", uuid).Delete(&models.Sellers{}).Error
	if err != nil {
		return exceptions.ParseGormError(tx, err)
	}
	return nil
}

func (r *CompRepositoriesImpl) FindOrderBySellerUUID(ctx *gin.Context, tx *gorm.DB, uuid string, params dto.OrderQueryParams) ([]models.Orders, *exceptions.Exception) {
	var orders []models.Orders
	query := tx.WithContext(ctx).
		Model(&models.Orders{}).
		Joins("JOIN products ON products.uuid = orders.product_uuid").
		Joins("JOIN sellers ON sellers.uuid = products.seller_uuid").
		Where("sellers.uuid = ?", uuid).
		Preload("Product").
		Preload("User")

	if params.Status != "" {
		if params.Status == "success" {
			query = query.Where("orders.status = ? OR orders.status = ?", "capture", "settlement")
		} else if params.Status == "pending" {
			query = query.Where("orders.status = ?", "pending")
		} else if params.Status == "failed" {
			query = query.Where("orders.status = ? OR orders.status = ? OR orders.status = ? OR orders.status = ?", "expire", "deny", "cancel", "failure")
		}
	}

	if params.ProductUUID != "" {
		query = query.Where("orders.product_uuid = ?", params.ProductUUID)
	}

	err := query.Order("orders.created_at DESC").
		Find(&orders).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return orders, nil
}
