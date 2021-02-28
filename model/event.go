package model

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Targets     []Target
	Teams       []Team
	Category    Category
	Description string `faker:"sentence"`
}
