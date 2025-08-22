package dto

import (
	"net/http"
	"sync"
)

type OrderReq struct {
	ProductUUID      string           `json:"product_uuid" validate:"required,uuid4"`
	Quantity         int              `json:"quantity" validate:"omitempty,number"`
	PaymentType OrderPaymentType `json:"order_payment_type" validate:"required,oneof='gopay' 'shopeepay' 'qris' 'bni' 'mandiri' 'cimb' 'bca' 'bri' 'maybank' 'permata' 'mega' 'indomaret' 'alfamart'"`
}

type OrderPaymentType string
type SubscriptionPaymentType = OrderPaymentType

const (
	PaymentTypeGopay         OrderPaymentType = "gopay"
	PaymentTypeShopeepay     OrderPaymentType = "shopeepay"
	PaymentTypeQris          OrderPaymentType = "qris"
	PaymentTypeVABankBni     OrderPaymentType = "bni"
	PaymentTypeVABankMandiri OrderPaymentType = "mandiri"
	PaymentTypeVABankCimb    OrderPaymentType = "cimb"
	PaymentTypeVABankBca     OrderPaymentType = "bca"
	PaymentTypeVABankBri     OrderPaymentType = "bri"
	PaymentTypeVABankMaybank OrderPaymentType = "maybank"
	PaymentTypeVABankPermata OrderPaymentType = "permata"
	PaymentTypeVABankMega    OrderPaymentType = "mega"
	PaymentTypeIndomart      OrderPaymentType = "indomaret"
	PaymentTypeAlfamart      OrderPaymentType = "alfamart"
)

type StreamClient struct {
	Writer  http.ResponseWriter
	Flusher http.Flusher
}

var Clients = make(map[string][]StreamClient)
var ClientsMutex = sync.Mutex{}
