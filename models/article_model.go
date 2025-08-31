package models

import (
	"time"

	"gorm.io/gorm"
)

type ArticleCategory struct {
	gorm.Model

	ID   uint   `gorm:"primaryKey"`
	UUID string `gorm:"not null;unique;index"`

	Name string `gorm:"not null;unique;index"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`

	Articles []Article `gorm:"foreignKey:CategoryUUID;references:UUID"`
}

type Article struct {
	gorm.Model

	ID           uint   `gorm:"primaryKey"`
	UUID         string `gorm:"not null;unique;index"`
	CategoryUUID string `gorm:"not null;index"`
	Slug         string `gorm:"not null;unique;index"`

	Title     string `gorm:"not null"`
	Thumbnail string `gorm:"not null"`
	Content   string `gorm:"not null"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`

	Category *ArticleCategory `gorm:"foreignKey:CategoryUUID;references:UUID"`
}
