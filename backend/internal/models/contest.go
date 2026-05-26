package models

import (
	"time"

	"gorm.io/gorm"
)

type Contest struct {
	gorm.Model

	ContestID string `gorm:"uniqueIndex;not null"`
	Title     string `gorm:"not null"`

	StartTime time.Time
	EndTime   time.Time

	Problems []Problem `gorm:"foreignKey:ContestRef;constraint:OnDelete:CASCADE"`
}
