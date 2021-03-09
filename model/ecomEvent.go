package model

import (
	"time"

	"gorm.io/gorm"
)

// Register transactions from the ecom API
type EcomEvent struct {
	ID        uint           `gorm:"primaryKey" faker:"-"`
	CreatedAt time.Time      `faker:"-"`
	UpdatedAt time.Time      `faker:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" faker:"-"`
	Team      string         // Team short-code/alias
	Success   bool
}
