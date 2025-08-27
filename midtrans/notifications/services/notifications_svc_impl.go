package services

import (
	"fmt"
	"ragamaya-api/midtrans/notifications/dto"
	"ragamaya-api/models"
	"ragamaya-api/pkg/exceptions"
	"ragamaya-api/pkg/helpers"
	"ragamaya-api/pkg/mapper"

	orderDTO "ragamaya-api/api/orders/dto"
	orderRepo "ragamaya-api/api/orders/repositories"
	orderService "ragamaya-api/api/orders/services"
	paymentRepo "ragamaya-api/api/payments/repositories"
	productRepo "ragamaya-api/api/products/repositories"
	walletDTO "ragamaya-api/api/wallets/dto"
	walletService "ragamaya-api/api/wallets/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	DB            *gorm.DB
	validate      *validator.Validate
	orderService  orderService.CompServices
	orderRepo     orderRepo.CompRepositories
	paymentRepo   paymentRepo.CompRepositories
	productRepo   productRepo.CompRepositories
	walletService walletService.CompServices
}

func NewComponentServices(
	db *gorm.DB,
	validate *validator.Validate,
	orderService orderService.CompServices,
	orderRepo orderRepo.CompRepositories,
	paymentRepo paymentRepo.CompRepositories,
	productRepo productRepo.CompRepositories,
	walletService walletService.CompServices,
) CompServices {
	return &CompServicesImpl{
		DB:            db,
		validate:      validate,
		orderService:  orderService,
		orderRepo:     orderRepo,
		paymentRepo:   paymentRepo,
		productRepo:   productRepo,
		walletService: walletService,
	}
}

func (s *CompServicesImpl) Payment(ctx *gin.Context, data dto.PaymentNotificationReq) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	orderData, err := s.orderRepo.FindByUUID(ctx, tx, data.OrderId)
	if err != nil {
		return err
	}

	if !s.isValidStatusTransition(orderData.Status, string(data.TransactionStatus)) {
		return exceptions.NewValidationException(fmt.Errorf("invalid status transition from %s to %s",
			orderData.Status, string(data.TransactionStatus)))
	}

	err = s.orderRepo.LockForUpdateWithTimeout(ctx, tx, data.OrderId, 5)
	if err != nil {
		return err
	}

	if orderData.Status == string(data.TransactionStatus) {
		return nil
	}

	err = s.orderRepo.Update(ctx, tx, models.Orders{UUID: data.OrderId, Status: string(data.TransactionStatus)})
	if err != nil {
		return err
	}

	err = s.paymentRepo.Update(ctx, tx, models.Payments{UUID: data.TransactionId, TransactionStatus: string(data.TransactionStatus)})
	if err != nil {
		return err
	}

	result := mapper.MapOrderMTO(*orderData)

	if data.TransactionStatus == dto.Capture || data.TransactionStatus == dto.Settlement {
		isOwned := s.productRepo.IsProductDigitalOwned(ctx, tx, orderData.UserUUID, orderData.ProductUUID)

		if isOwned {
			result.Status = "settlement"
			err = s.orderRepo.Update(ctx, tx, models.Orders{UUID: data.OrderId, Status: result.Status})
			if err != nil {
				return err
			}
			return nil
		}

		result.Status = "processing"
		err = s.orderRepo.Update(ctx, tx, models.Orders{UUID: data.OrderId, Status: result.Status})
		if err != nil {
			return err
		}

		s.orderService.SendStreamEvent(ctx, data.OrderId, orderDTO.OrderStreamRes{
			Type:    "info",
			Message: "Payment status updated",
			Body:    result,
		})

		err = s.productRepo.CreateProductDigitalOwned(ctx, tx, models.ProductDigitalOwned{
			ProductUUID: orderData.ProductUUID,
			UserUUID:    orderData.UserUUID,
		})
		if err != nil {
			if helpers.IsDuplicateKeyError(err) {
				return nil
			}
			return err
		}

		productData, err := s.productRepo.FindByUUID(ctx, tx, orderData.ProductUUID)
		if err != nil {
			return err
		}

		err = s.walletService.CreateTransactionWithTx(ctx, tx, walletDTO.WalletTransactionReq{
			UserUUID: productData.Seller.UserUUID,
			Amount: productData.Price,
			Type: string(models.Debit),
			Reference: "Order " + data.OrderId,
			Note: "Income from sales of product " + productData.Name,
		})
		if err != nil {
			return err
		}

		result.Status = "settlement"
		err = s.orderRepo.Update(ctx, tx, models.Orders{UUID: data.OrderId, Status: result.Status})
		if err != nil {
			return err
		}

		s.orderService.SendStreamEvent(ctx, data.OrderId, orderDTO.OrderStreamRes{
			Type:    "info",
			Message: "Ticket generated",
			Body:    result,
		})

	} else if data.TransactionStatus == dto.Deny ||
		data.TransactionStatus == dto.Cancel ||
		data.TransactionStatus == dto.Expire ||
		data.TransactionStatus == dto.Failure {

		if s.shouldRestoreStock(orderData.Status) {
			err := s.productRepo.RestoreStockByUUID(ctx, tx, orderData.ProductUUID)
			if err != nil {
				return err
			}
		}

		s.orderService.SendStreamEvent(ctx, data.OrderId, orderDTO.OrderStreamRes{
			Type:    "info",
			Message: "Payment status updated",
			Body:    result,
		})

	} else {
		s.orderService.SendStreamEvent(ctx, data.OrderId, orderDTO.OrderStreamRes{
			Type:    "info",
			Message: "Payment status updated",
			Body:    result,
		})
	}

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
