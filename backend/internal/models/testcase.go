package models

import "gorm.io/gorm"

type TestCase struct {
	gorm.Model

	InputFile  string `gorm:"not null"`
	OutputFile string `gorm:"not null"`

	IsHidden bool `gorm:"default:false"`

	ProblemRef uint
}
