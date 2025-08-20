package repositories

import (
	"fmt"
	"ragamaya-api/api/products/dto"
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

func (r *CompRepositoriesImpl) Create(ctx *gin.Context, tx *gorm.DB, data models.Products) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) FindByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.Products, *exceptions.Exception) {
	var seller models.Products
	err := tx.
		Where("uuid = ?", uuid).
		Preload("Thumbnails").
		Preload("DigitalFiles").
		First(&seller).
		Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return &seller, nil
}

func (r *CompRepositoriesImpl) Update(ctx *gin.Context, tx *gorm.DB, data models.Products) *exceptions.Exception {
	result := tx.Where("uuid = ?", data.UUID).Updates(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}
	return nil
}

func (r *CompRepositoriesImpl) Delete(ctx *gin.Context, tx *gorm.DB, uuid string) *exceptions.Exception {
	err := tx.Where("uuid = ?", uuid).Delete(&models.Products{}).Error
	if err != nil {
		return exceptions.ParseGormError(tx, err)
	}
	return nil
}

func (r *CompRepositoriesImpl) Search(ctx *gin.Context, tx *gorm.DB, searchReq dto.ProductSearchReq) ([]models.Products, int64, *exceptions.Exception) {
	var products []models.Products
	var total int64

	query := tx.WithContext(ctx).
		Model(&models.Products{}).
		Preload("Thumbnails").
		Preload("DigitalFiles")

	if searchReq.Keyword != nil && *searchReq.Keyword != "" {
		kw := fmt.Sprintf("%%%s%%", *searchReq.Keyword)
		query = query.Where(
			tx.Where("name ILIKE ?", kw).
				Or("description ILIKE ?", kw).
				Or("keywords ILIKE ?", kw),
		)
	}

	if searchReq.PriceMin != nil {
		query = query.Where("price >= ?", *searchReq.PriceMin)
	}
	if searchReq.PriceMax != nil {
		query = query.Where("price <= ?", *searchReq.PriceMax)
	}

	if searchReq.ProductType != nil {
		query = query.Where("product_type = ?", *searchReq.ProductType)
	}

	if searchReq.SellerUUID != nil {
		query = query.Where("seller_uuid = ?", *searchReq.SellerUUID)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, exceptions.ParseGormError(tx, err)
	}

	page := 1
	pageSize := 20
	if searchReq.Page != nil {
		page = *searchReq.Page
	}
	if searchReq.PageSize != nil {
		pageSize = *searchReq.PageSize
	}

	offset := (page - 1) * pageSize

	if err := query.
		Order("created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&products).Error; err != nil {
		return nil, 0, exceptions.ParseGormError(tx, err)
	}

	return products, total, nil
}
