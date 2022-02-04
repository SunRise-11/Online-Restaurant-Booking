package entities

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID           uint
	UserID       uint `gorm:"NOT NULL"`
	RestaurantID uint `gorm:"NOT NULL"`
	DateTime     time.Time
	Persons      int    `gorm:"NOT NULL;default:1"`
	Status       string `gorm:"NOT NULL;default:waiting for confirmation"`
	User         User
	Restaurant   Restaurant
}
