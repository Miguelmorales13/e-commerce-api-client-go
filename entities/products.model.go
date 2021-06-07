package entities

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID                 uint           `gorm:"not null;primaryKey;auto_increment" json:"id"`
	Name               string         `gorm:"not null" json:"name"`
	MainImage          string         `gorm:"not null" json:"mainImage"`
	HasMultiplesImages bool           `gorm:"not null" json:"hasMultiplesImages"`
	Price              float32        `gorm:"not null;type:decimal(8,2)" json:"price"`
	PriceDiscount      float32        `gorm:"not null;type:decimal(8,2)" json:"priceDiscount"`
	Description        string         `gorm:"not null;type:TEXT" json:"description"`
	Brand              string         `gorm:"not null" json:"brand"`
	Active             bool           `gorm:"not null;default:true" json:"active"`
	IsInDiscount       bool           `gorm:"not null;default:false" json:"isInDiscount"`
	CategoryId         int            `gorm:"not null" json:"categoryId"`
	Category           Category       `gorm:"" json:"category"`
	CreatedAt          time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt          time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Product) TableName() string {
	return "products"
}
