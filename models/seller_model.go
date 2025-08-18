package models

import (
	"time"

	"gorm.io/gorm"
)

type Sellers struct {
	gorm.Model

	ID       uint   `gorm:"primaryKey"`
	UUID     string `gorm:"not null;unique;index"`
	UserUUID string `gorm:"not null;unique;index"`

	Name      string `gorm:"not null"`
	Desc      string `gorm:"not null"`
	Address   string `gorm:"not null"`
	AvatarURL string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"null;default:null"`
}
