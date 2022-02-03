package transactions

import "time"

type TransactionResponseFormat struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type TransactionResponse struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	RestaurantID uint      `json:"restaurant_id"`
	DateTime     time.Time `json:"date_time"`
	Person       int       `json:"person"`
}
