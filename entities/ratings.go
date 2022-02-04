package entities

type Rating struct {
	ID           uint
	RestaurantID uint `gorm:"primaryKey"`
	UserID       uint `gorm:"primaryKey"`
	Rating       int  `gorm:"not null"`
	Comment      string
	User         User
	Restaurant   Restaurant
}
