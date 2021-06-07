package entities

import (
	"gorm.io/gorm"
	"time"
)

type ProductImages struct {
	ID         uint           `gorm:"not null;primaryKey;auto_increment" json:"id"`
	File       string         `gorm:"not null" json:"file"`
	Title      string         `gorm:"not null" json:"title"`
	Size       string         `gorm:"not null" json:"size"`
	Dimensions string         `gorm:"not null" json:"dimensions"`
	ProductId  int            `gorm:"not null" json:"productId"`
	Product    Product        `gorm:"" json:"product"`
	CreatedAt  time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt  time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

func (ProductImages) TableName() string {
	return "products_images"
}
