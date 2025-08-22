package models

import "time"

type Orders struct {
	ID          int64  `gorm:"primaryKey"`
	UUID        string `gorm:"not null;unique;index"`
	UserUUID    string `gorm:"not null;index"`
	ProductUUID string `gorm:"not null;index"`

	Quantity    int    `gorm:"not nul;default:1"`
	GrossAmt    int64  `gorm:"not null"`
	Status      string `gorm:"not null;default:pending"`
	PaymentType string `gorm:"not null"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`

	Payments []Payments `gorm:"foreignKey:OrderUUID;references:UUID" mapstructure:"payments"`
	User     Users      `gorm:"foreignKey:UserUUID;references:UUID"`
	Product  Products   `gorm:"foreignKey:ProductUUID;references:UUID"`
}
