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

func (r *CompRepositoriesImpl) Create(ctx *gin.Context, tx *gorm.DB, data models.Products) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) FindByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.Products, *exceptions.Exception) {
	var seller models.Products
	err := tx.Where("uuid = ?", uuid).First(&seller).Error
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
