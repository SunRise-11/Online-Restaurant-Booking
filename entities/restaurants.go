package entities

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	ID                 uint
	Email              string `gorm:"not null"`
	Password           string `gorm:"not null"`
	RestaurantDetailID uint
	RestaurantDetail   RestaurantDetail
}
