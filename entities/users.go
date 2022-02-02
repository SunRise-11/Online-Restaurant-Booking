package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string
	Password     string
	Name         string
	Phone_number string
	Reputation   int
	Balance      int
}
