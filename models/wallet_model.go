package models

import (
	"time"

	"gorm.io/gorm"
)

type TransactionType string

const (
	Debit  TransactionType = "debit"
	Credit TransactionType = "credit"
)

type PayoutStatus string

const (
	Pending   PayoutStatus = "pending"
	Completed PayoutStatus = "completed"
	Failed    PayoutStatus = "failed"
)

type PayoutBank string

const (
	BCA     PayoutBank = "bca"
	BNI     PayoutBank = "bni"
	BRI     PayoutBank = "bri"
	Mandiri PayoutBank = "mandiri"
)

type Wallet struct {
	gorm.Model

	ID       uint   `gorm:"primaryKey" json:"-"`
	UserUUID string `gorm:"not null;unique;index"`

	Balance int64 `gorm:"not null;default:0"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`

	TransactionHistory []WalletTransactionHistory `gorm:"foreignKey:WalletID;references:ID"`
}

type WalletTransactionHistory struct {
	gorm.Model

	ID       uint   `gorm:"primaryKey" json:"-"`
	UUID     string `gorm:"not null;unique;index"`
	WalletID uint   `gorm:"not null;index"`

	Amount    int64           `gorm:"not null"`
	Type      TransactionType `gorm:"not null"` // e.g., "credit" or "debit"
	Reference string          `gorm:"not null"` // e.g., "order_payment", "refund", etc.
	Note      string          `gorm:"type:text"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`
}

type WalletPayoutRequest struct {
	gorm.Model

	ID       uint   `gorm:"primaryKey" json:"-"`
	UUID     string `gorm:"not null;unique;index"`
	WalletID uint   `gorm:"not null;index"`

	Amount          int64        `gorm:"not null"`
	Status          PayoutStatus `gorm:"not null;default:pending"`
	BankName        string       `gorm:"not null"`
	BankAccount     string       `gorm:"not null"`
	BankAccountName string       `gorm:"not null"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`

	TransactionReceipt WalletPayoutTransactionReceipt `gorm:"foreignKey:PayoutUUID;references:UUID"`

	User Users `gorm:"-:all"` // to hold user info when joining with wallet
}

type WalletPayoutTransactionReceipt struct {
	gorm.Model

	ID uint `gorm:"primaryKey"`

	PayoutUUID string `gorm:"not null;unique;index"`
	ReceiptURL string `gorm:"not null"`
	Note       string

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`
}
