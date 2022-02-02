package entities

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID           uint
	UserID       uint
	RestaurantID uint
	Date         time.Time
	Time         time.Time
	Persons      uint
	Status       string `gorm:"NOT NULL;default:waiting for confirmation"`
	User         User
	Restaurant   Restaurant
}
