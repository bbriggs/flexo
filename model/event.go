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
	Targets     pq.Int64Array  `gorm:"type:integer[]"`
	Teams       pq.Int64Array  `gorm:"type:integer[]"`
	Category    int
	Description string `faker:"sentence"`
}
