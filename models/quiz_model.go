package models

import (
	"time"

	"gorm.io/gorm"
)

type QuizCategory struct {
	gorm.Model

	ID   uint   `gorm:"primaryKey"`
	UUID string `gorm:"not null;unique;index"`

	Name string `gorm:"not null;unique;index"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`

	Quizzes []Quiz `gorm:"foreignKey:CategoryUUID;references:UUID"`
}

type QuizLevel string

const (
	Beginner     QuizLevel = "beginner"
	Intermediate QuizLevel = "intermediate"
	Advanced     QuizLevel = "advanced"
)

type Quiz struct {
	gorm.Model

	ID           uint   `gorm:"primaryKey"`
	UUID         string `gorm:"not null;unique;index"`
	Slug         string `gorm:"not null;unique;index"`
	CategoryUUID string `gorm:"not null"`

	Title        string `gorm:"not null"`
	Desc         string
	Level        QuizLevel `gorm:"not null;index"`
	Estimate     int       `gorm:"not null"`
	MinimumScore int       `gorm:"not null"`
	Questions    string    `gorm:"not null"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`

	Category *QuizCategory `gorm:"foreignKey:CategoryUUID;references:UUID"`
}
