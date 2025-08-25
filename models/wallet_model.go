package models

import (
	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model

	ID       uint   `gorm:"primaryKey" json:"-"`
	UserUUID string `gorm:"not null;index"`

	Balance int64 `gorm:"not null;default:0"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`

	TransactionHistory []WalletTransactionHistory `gorm:"foreignKey:WalletID;references:ID"`
}

type WalletTransactionHistory struct {
	gorm.Model

	ID       uint `gorm:"primaryKey" json:"-"`
	WalletID uint `gorm:"not null;index"`

	Amount    int64  `gorm:"not null"`
	Type      string `gorm:"not null"` // e.g., "credit" or "debit"
	Reference string `gorm:"not null"` // e.g., "order_payment", "refund", etc.
	Note      string `gorm:"type:text"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`
}
