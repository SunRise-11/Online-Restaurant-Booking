package transactions

type LoginRequestFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
type TransactionRequestFormat struct {
	UserID       uint   `json:"user_id" form:"user_id"`
	RestaurantID uint   `json:"restaurant_id" form:"restaurant_id"`
	DateTime     string `json:"date_time" form:"date_time"`
	Persons      int    `json:"person" form:"person"`
}
