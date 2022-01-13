package model

import (
	"time"

	pq "github.com/lib/pq"
	"gorm.io/gorm"
)

type Event struct {
	ID          uint           `gorm:"primaryKey" faker:"-"`
	CreatedAt   time.Time      `faker:"-"`
	UpdatedAt   time.Time      `faker:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" faker:"-"`
	Targets     pq.Int64Array  `gorm:"type:integer[]" faker:"boundary_start=1, boundary_end=10"`
	Teams       pq.Int64Array  `gorm:"type:integer[]" faker:"boundary_start=1, boundary_end=10"`
	Category    int            `faker:"boundary_start=1, boundary_end=15"`
	Description string         `faker:"sentence"`
}
