package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          uint
	Email       string `gorm:"not null"`
	Password    string `gorm:"not null"`
	Name        string `gorm:"not null"`
	PhoneNumber string
	Reputation  int
	Balance     int
}
