package models

import (
	"time"

	"gorm.io/gorm"
)

type Payments struct {
	gorm.Model

	ID                     uint   `gorm:"primaryKey"`
	UUID                   string `gorm:"not null;unique;index"`
	UserUUID               string `gorm:"not null;index"`
	ProductUUID            string `gorm:"not null;index"`
	OrderUUID              string `gorm:"not null;unique;index"`
	GrossAmount            uint   `gorm:"not null"`
	PaymentType            string `gorm:"not null"`
	TransactionTime        string
	TransactionStatus      string `gorm:"not null"`
	FraudStatus            string
	MaskedCard             string
	StatusCode             string
	Bank                   string
	StatusMessage          string
	ApprovalCode           string
	ChannelResponseCode    string
	ChannelResponseMessage string
	Currency               string
	CardType               string
	RedirectURL            string
	InstallmentTerm        string
	Eci                    string
	SavedTokenID           string
	SavedTokenIDExpiredAt  string
	PointRedeemAmount      int
	PointRedeemQuantity    int
	PointBalanceAmount     string
	PermataVaNumber        string
	BillKey                string
	BillerCode             string
	Acquirer               string
	PaymentCode            string
	Store                  string
	QRString               string
	OnUs                   bool
	ThreeDsVersion         string
	ExpiryTime             string

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `gorm:"index"`

	PaymentActions   []PaymentActions   `gorm:"foreignKey:PaymentUUID;references:UUID" mapstructure:"actions"`
	PaymentVANumbers []PaymentVANumbers `gorm:"foreignKey:PaymentUUID;references:UUID" mapstructure:"va_numbers"`

	User    Users    `gorm:"foreignKey:UserUUID;references:UUID"`
	Product Products `gorm:"foreignKey:ProductUUID;references:UUID"`
	Order   Orders   `gorm:"foreignKey:OrderUUID;references:UUID"`
}

type PaymentActions struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey"`
	PaymentUUID string `gorm:"not null;index"`
	Name        string `gorm:"not null"`
	Method      string `gorm:"not null"`
	Url         string `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `gorm:"index"`

	Payment Payments `gorm:"foreignKey:PaymentUUID;references:UUID"`
}

type PaymentVANumbers struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey"`
	PaymentUUID string `gorm:"not null;index"`
	Bank        string `gorm:"not null"`
	VANumber    string `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `gorm:"index"`

	Payment Payments `gorm:"foreignKey:PaymentUUID;references:UUID"`
}
