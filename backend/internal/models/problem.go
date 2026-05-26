package models

import "gorm.io/gorm"

type Problem struct {
	gorm.Model

	ProblemID string `gorm:"uniqueIndex;not null"`

	ProblemTitle     string `gorm:"not null"`
	ProblemStatement string `gorm:"type:text"`

	TimeLimitMS   int `gorm:"not null"`
	MemoryLimitMB int `gorm:"not null"`

	ContestRef uint

	TestCases []TestCase `gorm:"foreignKey:ProblemRef;constraint:OnDelete:CASCADE"`
}
