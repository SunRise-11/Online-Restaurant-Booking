package entities

import (
	"time"

	"gorm.io/gorm"
)

type RestaurantDetail struct {
	gorm.Model
	ID             uint
	Name           string    `gorm:"not null"`
	Open           time.Time `gorm:"not null"`
	Close          time.Time `gorm:"not null"`
	Price          int       `gorm:"not null"`
	Latitude       float64   `gorm:"not null"`
	Longitude      float64   `gorm:"not null"`
	City           string    `gorm:"not null"`
	Address        string    `gorm:"not null"`
	PhoneNumber    string    `gorm:"not null"`
	ProfilePicture string    `gorm:"not null"`
	Seats          int       `gorm:"not null"`
	Description    string    `gorm:"not null"`
}
