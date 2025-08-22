package services

import (
	"ragamaya-api/api/orders/dto"
	"ragamaya-api/api/orders/repositories"
	"ragamaya-api/models"
	"ragamaya-api/pkg/exceptions"
	"ragamaya-api/pkg/helpers"
	"ragamaya-api/pkg/mapper"

	userDTO "ragamaya-api/api/users/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo         repositories.CompRepositories
	DB           *gorm.DB
	validate     *validator.Validate
	midtransCore *coreapi.Client
}

func NewComponentServices(compRepositories repositories.CompRepositories, db *gorm.DB, validate *validator.Validate, midtransCore *coreapi.Client) CompServices {
	return &CompServicesImpl{
		repo:         compRepositories,
		DB:           db,
		validate:     validate,
		midtransCore: midtransCore,
	}
}

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.OrderReq) {

}

func (s *CompServicesImpl) processPayment(ctx *gin.Context, tx *gorm.DB, order *models.Orders, userData *userDTO.UserOutput) *exceptions.Exception {
	chargeReq := s.createChargeRequest(order, order.UUID, userData)

	midtransRes, midtransErr := s.midtransCore.ChargeTransaction(chargeReq)
	if midtransErr != nil {
		tx.Rollback()
		return helpers.FormatMidtransErrorToException(midtransErr)
	}

	payment := mapper.MapChargeResponseToPaymentModel(*midtransRes)
	payment.OrderUUID = order.UUID
	payment.UserUUID = userData.UUID
	payment.ProductUUID = order.ProductUUID

	// err := s.paymentRepo.Create(ctx, tx, payment)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (s *CompServicesImpl) createChargeRequest(order *models.Orders, orderUUID string, userData *userDTO.UserOutput) *coreapi.ChargeReq {
	baseCharge := &coreapi.ChargeReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  order.UUID,
			GrossAmt: order.GrossAmt,
		},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: userData.Name,
			Email: userData.Email,
		},
		CustomExpiry: &coreapi.CustomExpiry{
			ExpiryDuration: 30,
			Unit:           "minute",
		},
	}

	switch dto.OrderPaymentType(order.PaymentType) {
	case dto.PaymentTypeGopay, dto.PaymentTypeShopeepay, dto.PaymentTypeQris:
		baseCharge.PaymentType = coreapi.CoreapiPaymentType(order.PaymentType)
	case dto.PaymentTypeIndomart, dto.PaymentTypeAlfamart:
		baseCharge.PaymentType = coreapi.PaymentTypeConvenienceStore
		baseCharge.ConvStore = &coreapi.ConvStoreDetails{
			Store:             string(order.PaymentType),
			Message:           "Ragamaya #" + order.UUID,
			AlfamartFreeText1: "Ragamaya",
			AlfamartFreeText2: "Order receipt",
			AlfamartFreeText3: "Ref: " + order.UUID,
		}
	default:
		baseCharge.PaymentType = coreapi.PaymentTypeBankTransfer
		baseCharge.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.Bank(order.PaymentType),
		}
	}

	return baseCharge
}
