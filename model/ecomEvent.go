package model

import (
	"fmt"
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

func (e EcomEvent) String() string {
	status := "Fail"
	if e.Success {
		status = "Success"
	}

	return fmt.Sprintf("by team %s (%s)", e.Team, status)
}
