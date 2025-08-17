package models

import (
	"time"

	"gorm.io/gorm"
)

type Files struct {
	gorm.Model

	ID               uint   `gorm:"primaryKey"`
	UUID             string `gorm:"not null;unique;index"`
	PublicURL        string `gorm:"not null"`
	OriginalFileName string `gorm:"not null"`
	Size             string `gorm:"not null"`
	Extension        string `gorm:"not null"`
	MimeType         string `gorm:"not null"`
	MimeSubType      string
	Meta             string `gorm:"not null"`
	
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time `gorm:"null;default:null"`
}
