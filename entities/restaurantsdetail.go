package entities

import (
	"time"

	"gorm.io/gorm"
)

type RestaurantDetail struct {
	gorm.Model
	ID             uint
	Name           string
	Open           time.Time
	Close          time.Time
	Price          int
	Latitude       float64
	Longitude      float64
	City           string
	Address        string
	PhoneNumber    string
	ProfilePicture string
	Seats          int
	Description    string
	Status         string `gorm:"default:DISABLED"`
}
