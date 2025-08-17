package models

import (
	"time"

	"gorm.io/gorm"
)

type Roles string

const (
	User   Roles = "user"
	Seller Roles = "seller"
)

type Users struct {
	gorm.Model

	ID              int64  `gorm:"primaryKey"`
	UUID            string `gorm:"not null;unique;index"`
	Email           string `gorm:"not null;unique;index"`
	IsEmailVerified bool   `gorm:"not null;default:false"`
	SUB             string `gorm:"not null;unique;index"`
	Name            string `gorm:"not null"`
	Role            Roles  `gorm:"not null;default:user"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`
}
