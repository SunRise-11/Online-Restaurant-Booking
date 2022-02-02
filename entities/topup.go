package entities

import "gorm.io/gorm"

type TopUp struct {
	gorm.Model
	ID         uint
	UserID     uint
	InvoiceID  string
	PaymentUrl string
	Total      int `gorm:"not null"`
	Status     string
	User       User
}
