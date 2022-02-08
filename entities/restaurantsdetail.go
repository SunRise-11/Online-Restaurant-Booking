package entities

import (
	"gorm.io/gorm"
)

type RestaurantDetail struct {
	gorm.Model
	ID             uint
	Name           string
	Open_Hour      string `gorm:"default:NULL"`
	Close_Hour     string `gorm:"default:NULL"`
	Open           string
	Close          string
	Price          int
	Latitude       float64
	Longitude      float64
	City           string
	Address        string
	PhoneNumber    string
	ProfilePicture string
	Seats          int
	Status         string `gorm:"default:DISABLED"`
	Description    string
	Rating         []Rating
}
