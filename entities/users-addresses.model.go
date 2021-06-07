package entities

import (
	"gorm.io/gorm"
	"time"
)

type UsersAddresses struct {
	ID              uint           `gorm:"not null;primaryKey;auto_increment" json:"id"`
	StreetAndNumber string         `gorm:"not null" json:"streetAndNumber"`
	City            string         `gorm:"not null" json:"city"`
	ZipCode         string         `gorm:"not null" json:"zipCode"`
	State           string         `gorm:"not null" json:"state"`
	Lat             string         `gorm:"" json:"lat"`
	Lng             string         `gorm:"" json:"lng"`
	UserId          int            `gorm:"not null" json:"userId"`
	User            User           `gorm:"not null" json:"user"`
	CreatedAt       time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt       time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

func (UsersAddresses) TableName() string {
	return "clients_addresses"
}
