package dto

import "time"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}

type OrderRes struct {
	UUID        string `json:"uuid"`
	UserUUID    string `json:"user_uuid"`
	ProductUUID string `json:"product_uuid"`
	Quantity    int    `json:"quantity"`
	GrossAmt    int64  `json:"amount"`
	Status      string `json:"status"`

	Payments []OrderPaymentRes `json:"payments,omitempty"`
}

type OrderChargeRes struct {
	UUID        string    `json:"uuid,omitempty"`
	UserUUID    string    `json:"user_uuid,omitempty"`
	ProductUUID string    `json:"product_uuid,omitempty"`
	Status      string    `json:"status,omitempty"`
	Quantity    int       `json:"quantity,omitempty"`
	GrossAmt    int64     `json:"gross_amount,omitempty"`
	PaymentType string    `json:"payment_type,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`

	Payments []OrderPaymentRes `json:"payments,omitempty"`
}

type OrderPaymentRes struct {
	UUID                   string                    `json:"uuid,omitempty"`
	UserUUID               string                    `json:"user_uuid,omitempty"`
	ProductUUID            string                    `json:"product_uuid,omitempty"`
	OrderUUID              string                    `json:"order_uuid,omitempty"`
	GrossAmount            uint                      `json:"gross_amount,omitempty"`
	PaymentType            string                    `json:"payment_type,omitempty"`
	TransactionTime        string                    `json:"transaction_time,omitempty"`
	TransactionStatus      string                    `json:"transaction_status,omitempty"`
	FraudStatus            string                    `json:"fraud_status,omitempty"`
	MaskedCard             string                    `json:"masked_card,omitempty"`
	StatusCode             string                    `json:"status_code,omitempty"`
	Bank                   string                    `json:"bank,omitempty"`
	StatusMessage          string                    `json:"status_message,omitempty"`
	ApprovalCode           string                    `json:"approval_code,omitempty"`
	ChannelResponseCode    string                    `json:"channel_response_code,omitempty"`
	ChannelResponseMessage string                    `json:"channel_response_message,omitempty"`
	Currency               string                    `json:"currency,omitempty"`
	CardType               string                    `json:"card_type,omitempty"`
	RedirectURL            string                    `json:"redirect_url,omitempty"`
	InstallmentTerm        string                    `json:"installment_term,omitempty"`
	Eci                    string                    `json:"eci,omitempty"`
	SavedTokenID           string                    `json:"saved_token_id,omitempty"`
	SavedTokenIDExpiredAt  string                    `json:"saved_token_id_expired_at,omitempty"`
	PointRedeemAmount      int                       `json:"point_redeem_amount,omitempty"`
	PointRedeemQuantity    int                       `json:"point_redeem_quantity,omitempty"`
	PointBalanceAmount     string                    `json:"point_balance_amount,omitempty"`
	PermataVaNumber        string                    `json:"permata_va_number,omitempty"`
	BillKey                string                    `json:"bill_key,omitempty"`
	BillerCode             string                    `json:"biller_code,omitempty"`
	Acquirer               string                    `json:"acquirer,omitempty"`
	PaymentCode            string                    `json:"payment_code,omitempty"`
	Store                  string                    `json:"store,omitempty"`
	QRString               string                    `json:"qr_string,omitempty"`
	OnUs                   bool                      `json:"on_us,omitempty"`
	ThreeDsVersion         string                    `json:"three_ds_version,omitempty"`
	ExpiryTime             string                    `json:"expiry_time,omitempty"`
	
	PaymentActions         []OrderPaymentActionRes   `json:"payment_actions,omitempty" mapstructure:"actions"`
	PaymentVANumbers       []OrderPaymentVANumberRes `json:"payment_va_numbers,omitempty" mapstructure:"va_numbers"`
}

type OrderPaymentActionRes struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	Url    string `json:"url"`
}

type OrderPaymentVANumberRes struct {
	Bank     string `json:"bank"`
	VANumber string `json:"va_number"`
}

type OrderStreamRes struct {
	Type    string      `json:"type"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}
