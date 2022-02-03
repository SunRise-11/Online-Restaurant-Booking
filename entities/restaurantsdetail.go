package entities

import (
	"gorm.io/gorm"
)

type RestaurantDetail struct {
	gorm.Model
	ID              uint
	Name            string
	OperationalHour string
	Open            string
	Close           string
	Price           int
	Latitude        float64
	Longitude       float64
	City            string
	Address         string
	PhoneNumber     string
	ProfilePicture  string
	Seats           int
	Status          string `gorm:"default:close"`
	Description     string
}
