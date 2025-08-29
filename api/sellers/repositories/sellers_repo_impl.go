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
		Preload("Product.Thumbnails").
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

func (r *CompRepositoriesImpl) Analytics(ctx *gin.Context, tx *gorm.DB, sellerUUID string) (*dto.AnalyticsRes, *exceptions.Exception) {
	var res dto.AnalyticsRes

	if err := tx.Model(&models.Products{}).
		Where("seller_uuid = ?", sellerUUID).
		Count(&res.TotalProducts).Error; err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}

	if err := tx.Model(&models.Orders{}).
		Joins("JOIN products ON products.uuid = orders.product_uuid").
		Where("products.seller_uuid = ?", sellerUUID).
		Count(&res.TotalOrders).Error; err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}

	if err := tx.Model(&models.Orders{}).
		Select("COALESCE(SUM(gross_amt),0)").
		Joins("JOIN products ON products.uuid = orders.product_uuid").
		Where("products.seller_uuid = ? AND orders.status = ?", sellerUUID, "settlement").
		Scan(&res.TotalRevenue).Error; err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	res.TotalRevenueCurrency = "IDR"

	var monthly []dto.MonthlyRevenueRes
	query := `
				WITH months AS (
					SELECT TO_CHAR(date_trunc('month', NOW()) - INTERVAL '1 month' * gs, 'YYYY-MM') as month
					FROM generate_series(0, 11) gs
				)
				SELECT m.month,
					COALESCE(SUM(o.gross_amt), 0) as revenue,
					'IDR' as currency
				FROM months m
				LEFT JOIN orders o 
					ON TO_CHAR(o.created_at, 'YYYY-MM') = m.month
				AND o.status = 'settlement'
				LEFT JOIN products p
					ON p.uuid = o.product_uuid
				AND p.seller_uuid = ?
				GROUP BY m.month
				ORDER BY m.month ASC
			`
	if err := tx.Raw(query, sellerUUID).Scan(&monthly).Error; err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	res.MonthlyRevenue = monthly

	query = `
				WITH months AS (
					SELECT TO_CHAR(date_trunc('month', NOW()) - INTERVAL '1 month' * gs, 'YYYY-MM') as month
					FROM generate_series(0, 11) gs
				)
				SELECT m.month,
					COALESCE(COUNT(DISTINCT o.uuid), 0) as total_orders
				FROM months m
				LEFT JOIN orders o 
					ON TO_CHAR(o.created_at, 'YYYY-MM') = m.month
				LEFT JOIN products p
					ON p.uuid = o.product_uuid
				AND p.seller_uuid = ?
				GROUP BY m.month
				ORDER BY m.month ASC
			`
	var monthlyOrders []dto.MonthlyOrdersRes
	if err := tx.Raw(query, sellerUUID).Scan(&monthlyOrders).Error; err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	res.MonthlyOrders = monthlyOrders

	return &res, nil
}
