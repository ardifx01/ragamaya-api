package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"ragamaya-api/api/orders/dto"
	"ragamaya-api/api/orders/repositories"
	"ragamaya-api/models"
	"ragamaya-api/pkg/exceptions"
	"ragamaya-api/pkg/helpers"
	"ragamaya-api/pkg/mapper"
	"strconv"
	"time"

	paymentRepo "ragamaya-api/api/payments/repositories"
	productRepo "ragamaya-api/api/products/repositories"
	userDTO "ragamaya-api/api/users/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo         repositories.CompRepositories
	DB           *gorm.DB
	validate     *validator.Validate
	midtransCore *coreapi.Client
	paymentRepo  paymentRepo.CompRepositories
	productRepo  productRepo.CompRepositories
}

func NewComponentServices(compRepositories repositories.CompRepositories, db *gorm.DB, validate *validator.Validate, midtransCore *coreapi.Client, paymentRepo paymentRepo.CompRepositories, productRepo productRepo.CompRepositories) CompServices {
	return &CompServicesImpl{
		repo:         compRepositories,
		DB:           db,
		validate:     validate,
		midtransCore: midtransCore,
		paymentRepo:  paymentRepo,
		productRepo:  productRepo,
	}
}

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.OrderReq) (*dto.OrderChargeRes, *exceptions.Exception) {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return nil, exceptions.NewValidationException(validateErr)
	}

	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		return nil, err
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	order, err := s.prepareOrder(ctx, tx, data, &userData)
	if err != nil {
		return nil, err
	}

	err = s.repo.Create(ctx, tx, *order)
	if err != nil {
		return nil, err
	}

	err = s.productRepo.DecrementStockByUUID(ctx, tx, data.ProductUUID)
	if err != nil {
		return nil, err
	}

	if order.GrossAmt == 0 {
		err = s.processFreePayment(ctx, tx, order, &userData)
		if err != nil {
			return nil, err
		}
	} else {
		err = s.processPayment(ctx, tx, order, &userData)
		if err != nil {
			return nil, err
		}
	}

	orderResult, err := s.repo.FindByUUID(ctx, tx, order.UUID)
	if err != nil {
		return nil, err
	}

	output := mapper.MapOrderModelToChargeOutput(*orderResult)

	return &output, nil
}

func (s *CompServicesImpl) prepareOrder(ctx *gin.Context, tx *gorm.DB, data dto.OrderReq, userData *userDTO.UserOutput) (*models.Orders, *exceptions.Exception) {
	productData, err := s.productRepo.FindByUUID(ctx, tx, data.ProductUUID)
	if err != nil {
		return nil, err
	}

	if productData.ProductType == models.Digital {
		isOwned := s.productRepo.IsProductDigitalOwned(ctx, tx, userData.UUID, productData.UUID)
		if isOwned {
			return nil, exceptions.NewException(http.StatusForbidden, exceptions.ErrAlreadyOwned)
		}

		data.Quantity = 1
	}

	err = s.validateOrderQuantity(data.Quantity, productData)
	if err != nil {
		return nil, err
	}

	order := mapper.MapOrderITM(data)
	order.UUID = uuid.NewString()
	order.UserUUID = userData.UUID
	order.ProductUUID = productData.UUID
	order.Status = "pending"

	if productData.Price > 0 {
		serviceFee := 0
		if os.Getenv("ORDER_SERVICE_FEE") != "" {
			serviceFeeEnv, exc := strconv.Atoi(os.Getenv("ORDER_SERVICE_FEE"))
			if exc != nil {
				return nil, exceptions.NewException(http.StatusInternalServerError, exc.Error())
			}

			serviceFee = serviceFeeEnv
		}

		order.GrossAmt = (productData.Price * int64(data.Quantity)) + int64(serviceFee)
	} else {
		order.GrossAmt = 0
		order.PaymentType = "free"
	}

	return &order, nil
}

func (s *CompServicesImpl) validateOrderQuantity(quantity int, ticketCategory *models.Products) *exceptions.Exception {
	if quantity > ticketCategory.Stock {
		return exceptions.NewException(http.StatusBadRequest, exceptions.ErrCheckoutQuantityMoreThanStocks)
	}
	if quantity > ticketCategory.Stock {
		return exceptions.NewException(http.StatusBadRequest, exceptions.ErrCheckoutQuantityMoreThanAllowed)
	}
	return nil
}

func (s *CompServicesImpl) processPayment(ctx *gin.Context, tx *gorm.DB, order *models.Orders, userData *userDTO.UserOutput) *exceptions.Exception {
	chargeReq := s.createChargeRequest(order, userData)

	midtransRes, midtransErr := s.midtransCore.ChargeTransaction(chargeReq)
	if midtransErr != nil {
		tx.Rollback()
		return helpers.FormatMidtransErrorToException(midtransErr)
	}

	payment := mapper.MapChargeResponseToPaymentModel(*midtransRes)
	payment.OrderUUID = order.UUID
	payment.UserUUID = userData.UUID
	payment.ProductUUID = order.ProductUUID

	err := s.paymentRepo.Create(ctx, tx, payment)
	if err != nil {
		return err
	}

	return nil
}

func (s *CompServicesImpl) processFreePayment(ctx *gin.Context, tx *gorm.DB, order *models.Orders, userData *userDTO.UserOutput) *exceptions.Exception {
	var data models.Payments

	data.UUID = uuid.NewString()
	data.UserUUID = userData.UUID
	data.ProductUUID = order.ProductUUID
	data.OrderUUID = order.UUID
	data.GrossAmount = 0
	data.PaymentType = "free"
	data.TransactionTime = time.Now().UTC().Add(time.Hour * 7).Format("2006-01-02 15:04:05")
	data.TransactionStatus = "pending"
	data.FraudStatus = "accept"
	data.StatusCode = "201"
	data.StatusMessage = "Free transaction is created"
	data.Currency = "IDR"
	data.PointRedeemAmount = 0
	data.PointRedeemQuantity = 0
	data.OnUs = false
	data.ExpiryTime = time.Now().Add(time.Hour * 24).String()

	err := s.paymentRepo.Create(ctx, tx, data)
	if err != nil {
		return err
	}

	return nil
}

func (s *CompServicesImpl) createChargeRequest(order *models.Orders, userData *userDTO.UserOutput) *coreapi.ChargeReq {
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
		ShopeePay: &coreapi.ShopeePayDetails{
			CallbackUrl: fmt.Sprintf("%s/payments/%s", os.Getenv("FRONTEND_BASE_URL"), order.UUID),
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

func (s *CompServicesImpl) RemoveStreamClient(ctx *gin.Context, orderUUID string, client dto.StreamClient) {
	dto.ClientsMutex.Lock()
	defer dto.ClientsMutex.Unlock()

	newClients := []dto.StreamClient{}
	for _, c := range dto.Clients[orderUUID] {
		if c != client {
			newClients = append(newClients, c)
		}
	}
	dto.Clients[orderUUID] = newClients
}

func (s *CompServicesImpl) SendStreamEvent(ctx *gin.Context, orderUUID string, data dto.OrderStreamRes) {
	dto.ClientsMutex.Lock()
	defer dto.ClientsMutex.Unlock()

	if clients, exists := dto.Clients[orderUUID]; exists {
		for _, client := range clients {
			jsonData, _ := json.Marshal(data)
			_, err := fmt.Fprintf(client.Writer, "data:%s\n\n", jsonData)
			if err != nil {
				log.Println(err)
				s.RemoveStreamClient(ctx, orderUUID, client)
				continue
			}
			client.Flusher.Flush()
		}
	}
}

func (s *CompServicesImpl) FindByUUID(ctx *gin.Context, uuid string) (*dto.OrderRes, *exceptions.Exception) {
	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		return nil, err
	}

	orderData, err := s.repo.FindByUUID(ctx, s.DB, uuid)
	if err != nil {
		return nil, err
	}

	if orderData.UserUUID != userData.UUID {
		return nil, exceptions.NewException(http.StatusForbidden, exceptions.ErrForbidden)
	}

	output := mapper.MapOrderMTO(*orderData)

	return &output, nil
}

func (s *CompServicesImpl) FindByUserUUID(ctx *gin.Context) ([]dto.OrderRes, *exceptions.Exception) {
	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		return nil, err
	}

	orderData, err := s.repo.FindByUserUUID(ctx, s.DB, userData.UUID)
	if err != nil {
		return nil, err
	}

	var result []dto.OrderRes
	for _, order := range orderData {
		output := mapper.MapOrderMTO(order)
		result = append(result, output)
	}

	return result, nil
}
