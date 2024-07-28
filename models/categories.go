package models

import (

	"gorm.io/gorm"
)

// Category represents the category model
type Category struct {
	gorm.Model
	Name string `gorm:"size:255;not null"`
}
