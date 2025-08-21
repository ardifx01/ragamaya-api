package dto

type TransactionStatus string

const (
	Capture       TransactionStatus = "capture"
	Settlement    TransactionStatus = "settlement"
	Pending       TransactionStatus = "pending"
	Deny          TransactionStatus = "deny"
	Cancel        TransactionStatus = "cancel"
	Expire        TransactionStatus = "expire"
	Failure       TransactionStatus = "failure"
	Refund        TransactionStatus = "refund"
	PartialRefund TransactionStatus = "partial_refund"
	Authorize     TransactionStatus = "authorize"
)

type PaymentNotificationReq struct {
	TransactionTime   string            `json:"transaction_time"`
	TransactionStatus TransactionStatus `json:"transaction_status" validate:"required,oneof='capture' 'settlement' 'pending' 'deny' 'cancel' 'expire' 'failure' 'refund' 'partial_refund' 'authorize'"`
	TransactionId     string            `json:"transaction_id"`
	StatusMessage     string            `json:"status_message"`
	StatusCode        string            `json:"status_code"`
	GrossAmount       string            `json:"gross_amount" validate:"required"`
	SignatureKey      string            `json:"signature_key" validate:"required"`
	SettlementTime    string            `json:"settlement_time"`
	PaymentType       string            `json:"payment_type"`
	OrderId           string            `json:"order_id" validate:"required"`
	MerchantId        string            `json:"merchant_id"`
	FraudStatus       string            `json:"fraud_status" validate:"required"`
	Currency          string            `json:"currency"`
}
