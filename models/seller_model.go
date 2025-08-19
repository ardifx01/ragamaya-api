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
	Whatsapp  string `gorm:"not null;unique"`
	AvatarURL string

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`
}

type SellerJWTPayload struct {
	UUID      string `json:"uuid"`
	UserUUID  string `json:"user_uuid"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

func (s *Sellers) ToJWTPayload() SellerJWTPayload {
	return SellerJWTPayload{
		UUID:      s.UUID,
		UserUUID:  s.UserUUID,
		Name:      s.Name,
		AvatarURL: s.AvatarURL,
	}
}
