package services

import (
	"ragamaya-api/midtrans/notifications/dto"
	"ragamaya-api/pkg/exceptions"
	"ragamaya-api/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	DB       *gorm.DB
	validate *validator.Validate
}

func NewComponentServices(
	db *gorm.DB,
	validate *validator.Validate,
) CompServices {
	return &CompServicesImpl{
		DB:       db,
		validate: validate,
	}
}

func (s *CompServicesImpl) Payment(ctx *gin.Context, data dto.PaymentNotificationReq) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	return nil
}

func (s *CompServicesImpl) isValidStatusTransition(currentStatus, newStatus string) bool {
	validTransitions := map[string][]string{
		"pending":    {"capture", "settlement", "deny", "cancel", "expire", "failure"},
		"capture":    {"settlement", "cancel", "expire", "failure"},
		"settlement": {},
		"deny":       {},
		"cancel":     {},
		"expire":     {},
		"failure":    {"pending"},
	}

	allowedTransitions, exists := validTransitions[currentStatus]
	if !exists {
		return false
	}

	for _, allowed := range allowedTransitions {
		if allowed == newStatus {
			return true
		}
	}
	return false
}

func (s *CompServicesImpl) shouldRestoreStock(currentStatus string) bool {
	restoreStates := []string{"pending", "capture", "processing"}
	for _, state := range restoreStates {
		if state == currentStatus {
			return true
		}
	}
	return false
}
