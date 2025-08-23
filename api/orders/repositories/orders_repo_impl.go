package repositories

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"ragamaya-api/models"
	"ragamaya-api/pkg/exceptions"
	"strings"
	"time"

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

func (r *CompRepositoriesImpl) LockForUpdateWithTimeout(ctx *gin.Context, tx *gorm.DB, orderUUID string, timeoutSeconds int) *exceptions.Exception {
	var order models.Orders

	lockCtx, cancel := context.WithTimeout(ctx, time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	err := tx.WithContext(lockCtx).
		Select("uuid").
		Where("uuid = ?", orderUUID).
		Set("gorm:query_option", "FOR UPDATE NOWAIT").
		First(&order).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exceptions.NewException(http.StatusNotFound, "Order not found")
		}

		if isLockTimeoutError(err) {
			return exceptions.NewException(http.StatusConflict, "Order is being processed by another request")
		}

		return exceptions.NewException(http.StatusInternalServerError, fmt.Sprintf("Failed to lock order: %v", err))
	}

	return nil
}

func isLockTimeoutError(err error) bool {
	errMsg := strings.ToLower(err.Error())

	if strings.Contains(errMsg, "could not obtain lock") ||
		strings.Contains(errMsg, "lock_not_available") ||
		strings.Contains(errMsg, "lock timeout") {
		return true
	}

	return false
}
