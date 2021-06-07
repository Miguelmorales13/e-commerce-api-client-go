package entities

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID         uint       `gorm:"not null;primaryKey;auto_increment" json:"id"`
	Name       string     `gorm:"not null" json:"name"`
	Nivel      int        `gorm:"not null" json:"nivel"`
	CategoryID uint       `gorm:"not null" json:"categoryId"`
	Categories []Category `gorm:"not null" json:"categories"`
	
	CreatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Category) TableName() string {
	return "categories_products"
}
