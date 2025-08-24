package dto

import "time"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}

type PaymentRes struct {
	UUID                   string `json:"uuid,omitempty"`
	UserUUID               string `json:"user_uuid,omitempty"`
	ProductUUID            string `json:"product_uuid,omitempty"`
	OrderUUID              string `json:"order_uuid,omitempty"`
	GrossAmount            uint   `json:"gross_amount,omitempty"`
	PaymentType            string `json:"payment_type,omitempty"`
	TransactionTime        string `json:"transaction_time,omitempty"`
	TransactionStatus      string `json:"transaction_status,omitempty"`
	FraudStatus            string `json:"fraud_status,omitempty"`
	MaskedCard             string `json:"masked_card,omitempty"`
	StatusCode             string `json:"status_code,omitempty"`
	Bank                   string `json:"bank,omitempty"`
	StatusMessage          string `json:"status_message,omitempty"`
	ApprovalCode           string `json:"approval_code,omitempty"`
	ChannelResponseCode    string `json:"channel_response_code,omitempty"`
	ChannelResponseMessage string `json:"channel_response_message,omitempty"`
	Currency               string `json:"currency,omitempty"`
	CardType               string `json:"card_type,omitempty"`
	RedirectURL            string `json:"redirect_url,omitempty"`
	InstallmentTerm        string `json:"installment_term,omitempty"`
	Eci                    string `json:"eci,omitempty"`
	SavedTokenID           string `json:"saved_token_id,omitempty"`
	SavedTokenIDExpiredAt  string `json:"saved_token_id_expired_at,omitempty"`
	PointRedeemAmount      int    `json:"point_redeem_amount,omitempty"`
	PointRedeemQuantity    int    `json:"point_redeem_quantity,omitempty"`
	PointBalanceAmount     string `json:"point_balance_amount,omitempty"`
	PermataVaNumber        string `json:"permata_va_number,omitempty"`
	BillKey                string `json:"bill_key,omitempty"`
	BillerCode             string `json:"biller_code,omitempty"`
	Acquirer               string `json:"acquirer,omitempty"`
	PaymentCode            string `json:"payment_code,omitempty"`
	Store                  string `json:"store,omitempty"`
	QRString               string `json:"qr_string,omitempty"`
	OnUs                   bool   `json:"on_us,omitempty"`
	ThreeDsVersion         string `json:"three_ds_version,omitempty"`
	ExpiryTime             string `json:"expiry_time,omitempty"`

	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	PaymentActions   []PaymentActionOutput   `json:"payment_actions,omitempty" mapstructure:"actions"`
	PaymentVANumbers []PaymentVANumberOutput `json:"payment_va_numbers,omitempty" mapstructure:"va_numbers"`

	Product ProductRes `json:"product,omitempty"`
}

type PaymentActionOutput struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	Url    string `json:"url"`
}

type PaymentVANumberOutput struct {
	Bank     string `json:"bank"`
	VANumber string `json:"va_number"`
}

type ProductType string

const (
	Digital  ProductType = "digital"
	Physical ProductType = "physical"
)

type ProductRes struct {
	UUID        string      `json:"uuid"`
	SellerUUID  string      `json:"seller_uuid"`
	ProductType ProductType `json:"product_type" validate:"required,oneof=digital physical"`
	Name        string      `json:"name" validate:"required"`
	Description string      `json:"description" validate:"required"`
	Price       uint        `json:"price" validate:"required"`
	Stock       int         `json:"stock" validate:"required"`
	Keywords    string      `json:"keywords" validate:"required"`

	Thumbnails []ProductThumbnailsRes `json:"thumbnails" validate:"required,dive"`
}

type ProductThumbnailsRes struct {
	ThumbnailURL string `json:"thumbnail_url" validate:"required,url"`
}
