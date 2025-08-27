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
