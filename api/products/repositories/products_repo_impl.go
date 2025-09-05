package repositories

import (
	"fmt"
	"net/http"
	"ragamaya-api/api/products/dto"
	"ragamaya-api/models"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
		Preload("Seller").
		Preload("Thumbnails").
		Preload("DigitalFiles", func(db *gorm.DB) *gorm.DB {
			return db.Omit("file_url")
		}).
		First(&seller).
		Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return &seller, nil
}

func (r *CompRepositoriesImpl) FindByUUIDWithFile(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.Products, *exceptions.Exception) {
	var seller models.Products
	err := tx.
		Where("uuid = ?", uuid).
		Preload("Seller").
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
		Preload("DigitalFiles", func(db *gorm.DB) *gorm.DB {
			return db.Omit("file_url")
		})

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

func (r *CompRepositoriesImpl) DecrementStockByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) *exceptions.Exception {
	var product models.Products
	result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("uuid = ?", uuid).
		First(&product)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	if product.Stock == 0 {
		tx.Rollback()
		return exceptions.NewException(http.StatusBadRequest, exceptions.ErrCheckoutQuantityMoreThanStocks)
	}

	result = tx.Model(&product).
		Update("stock", product.Stock-1)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) RestoreStockByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) *exceptions.Exception {
	var product models.Products
	result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("uuid = ?", uuid).
		First(&product)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	result = tx.Model(&product).
		Update("stock", product.Stock+1)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) CreateProductDigitalOwned(ctx *gin.Context, tx *gorm.DB, data models.ProductDigitalOwned) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) FindProductDigitalOwned(ctx *gin.Context, tx *gorm.DB, userUUID string) ([]models.Products, *exceptions.Exception) {
	var products []models.Products
	err := tx.
		Joins("JOIN product_digital_owneds ON product_digital_owneds.product_uuid = products.uuid").
		Where("product_digital_owneds.user_uuid = ?", userUUID).
		Preload("Thumbnails").
		Preload("DigitalFiles").
		Find(&products).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}

	return products, nil
}

func (r *CompRepositoriesImpl) IsProductDigitalOwned(ctx *gin.Context, tx *gorm.DB, userUUID string, productUUID string) bool {
	var productDigitalOwned models.ProductDigitalOwned
	result := tx.
		Where("user_uuid = ?", userUUID).
		Where("product_uuid = ?", productUUID).
		First(&productDigitalOwned)
	if result.Error != nil {
		return false
	}
	return true
}

func (r *CompRepositoriesImpl) DeleteThumbnail(ctx *gin.Context, tx *gorm.DB, productUUID string, id uint) *exceptions.Exception {
	err := tx.Where("product_uuid = ?", productUUID).Where("id = ?", id).Where("deleted_at IS NULL").Delete(&models.ProductThumbnails{}).Error
	if err != nil {
		return exceptions.ParseGormError(tx, err)
	}
	return nil
}
