package entities

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID             uint           ` gorm:"not null;primaryKey;auto_increment" json:"id"`
	Name           string         `validate:"required" gorm:"not null" json:"name"`
	LastName       string         `validate:"required" gorm:"not null" json:"lastName"`
	SecondLastName string         `validate:"required" gorm:"not null" json:"secondLastName"`
	Email          string         `validate:"required" gorm:"not null" json:"email"`
	Password       string         `validate:"required" gorm:"not null" json:"-"`
	Active         bool           ` gorm:"not null;default:true" json:"active"`
	IsVerified     bool           ` gorm:"not null;default:false" json:"isVerified"`
	CreatedAt      time.Time      ` gorm:"not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt      time.Time      ` gorm:"not null;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt      gorm.DeletedAt ` gorm:"index" json:"-"`
}

func (User) TableName() string {
	return "clients"
}
