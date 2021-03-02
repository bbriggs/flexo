package model

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID          uint           `gorm:"primaryKey" faker:"-"`
	CreatedAt   time.Time      `faker:"-"`
	UpdatedAt   time.Time      `faker:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" faker:"-"`
	Targets     []Target
	Teams       []Team
	Category    Category
	Description string `faker:"sentence"`
}
