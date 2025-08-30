package mapper

import (
	"ragamaya-api/api/orders/dto"
	"ragamaya-api/models"
	"strconv"

	"github.com/go-viper/mapstructure/v2"
	"github.com/midtrans/midtrans-go/coreapi"
)

func MapChargeResponseToPaymentModel(input coreapi.ChargeResponse) models.Payments {
	var data models.Payments
	grossAmount, _ := strconv.ParseFloat(input.GrossAmount, 64)

	data.UUID = input.TransactionID
	data.OrderUUID = input.OrderID
	data.GrossAmount = uint(grossAmount)
	data.PaymentType = input.PaymentType
	data.TransactionTime = input.TransactionTime
	data.TransactionStatus = input.TransactionStatus
	data.FraudStatus = input.FraudStatus
	data.MaskedCard = input.MaskedCard
	data.StatusCode = input.StatusCode
	data.Bank = input.Bank
	data.StatusMessage = input.StatusMessage
	data.ApprovalCode = input.ApprovalCode
	data.ChannelResponseCode = input.ChannelResponseCode
	data.ChannelResponseMessage = input.ChannelResponseMessage
	data.Currency = input.Currency
	data.CardType = input.CardType
	data.RedirectURL = input.RedirectURL
	data.InstallmentTerm = input.InstallmentTerm
	data.Eci = input.Eci
	data.SavedTokenID = input.SavedTokenID
	data.SavedTokenIDExpiredAt = input.SavedTokenIDExpiredAt
	data.PointRedeemAmount = input.PointRedeemAmount
	data.PointRedeemQuantity = input.PointRedeemQuantity
	data.PointBalanceAmount = input.PointBalanceAmount
	data.PermataVaNumber = input.PermataVaNumber
	data.BillKey = input.BillKey
	data.BillerCode = input.BillerCode
	data.Acquirer = input.Acquirer
	data.PaymentCode = input.PaymentCode
	data.Store = input.Store
	data.QRString = input.QRString
	data.OnUs = input.OnUs
	data.ThreeDsVersion = input.ThreeDsVersion
	data.ExpiryTime = input.ExpiryTime

	mapstructure.Decode(input.Actions, &data.PaymentActions)
	mapstructure.Decode(input.VaNumbers, &data.PaymentVANumbers)

	return data
}

func MapOrderMTCO(model models.Orders) dto.OrderChargeRes {
	var output dto.OrderChargeRes
	mapstructure.Decode(model, &output)
	output.CreatedAt = model.CreatedAt
	output.UpdatedAt = model.UpdatedAt
	return output
}
