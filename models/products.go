package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string    `gorm:"size:255;not null"`
	Brand         string    `gorm:"size:255"`
	CategoryID    uint
	Price         float64   `gorm:"type:decimal(10,2)"`
	StockQuantity int       `gorm:"column:stock_quantity"`
	Description   string    `gorm:"type:text"`
	ReleaseDate   time.Time `gorm:"column:release_date"`
	ImageURL      string    `gorm:"size:255"`
}

