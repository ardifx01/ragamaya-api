package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductType string

const (
	Digital  ProductType = "digital"
	Physical ProductType = "physical"
)

type Products struct {
	gorm.Model

	ID          uint        `gorm:"primaryKey"`
	UUID        string      `gorm:"not null;unique;index"`
	SellerUUID  string      `gorm:"not null;index"`
	ProductType ProductType `gorm:"not null;index"`
	Name        string      `gorm:"not null"`
	Description string      `gorm:"not null"`
	Price       int64       `gorm:"not null"`
	Stock       int         `gorm:"not null"`
	Keywords    string      `gorm:"not null"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`

	Thumbnails   []ProductThumbnails   `gorm:"foreignKey:ProductUUID;references:UUID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	DigitalFiles []ProductDigitalFiles `gorm:"foreignKey:ProductUUID;references:UUID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type ProductThumbnails struct {
	gorm.Model

	ID           uint   `gorm:"primaryKey"`
	ProductUUID  string `gorm:"not null;index"`
	ThumbnailURL string `gorm:"not null"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`
}

type ProductDigitalFiles struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey"`
	ProductUUID string `gorm:"not null;index"`

	FileURL     string `gorm:"not null"`
	Description string `gorm:"not null"`
	Extension   string `gorm:"not null"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`
}

type ProductDigitalOwned struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey"`
	ProductUUID string `gorm:"not null;index;uniqueIndex:idx_user_product"`
	UserUUID    string `gorm:"not null;index;uniqueIndex:idx_user_product"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`

	Product Products `gorm:"foreignKey:ProductUUID;references:UUID"`
	User    Users    `gorm:"foreignKey:UserUUID;references:UUID"`
}

