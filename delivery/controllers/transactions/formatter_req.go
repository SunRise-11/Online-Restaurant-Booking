package transactions

type TransactionRequestFormat struct {
	ID           uint   `json:"id" form:"id"`
	UserID       uint   `json:"user_id" form:"user_id"`
	RestaurantID uint   `json:"restaurant_id" form:"restaurant_id"`
	DateTime     string `json:"date_time" form:"date_time"`
	Persons      int    `json:"person" form:"person"`
	Status       string `json:"status" form:"status"`
}
