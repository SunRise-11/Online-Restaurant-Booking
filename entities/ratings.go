package entities

type Rating struct {
	ID                 uint
	RestaurantDetailID uint `gorm:"primaryKey"`
	UserID             uint `gorm:"primaryKey"`
	Rating             int  `gorm:"not null"`
	Comment            string
	User               User
	RestaurantDetail   RestaurantDetail
}
