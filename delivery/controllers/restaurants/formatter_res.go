package restaurants

import "time"

type LoginResponseFormat struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type RestaurantResponseFormat struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type RestaurantResponse struct {
	ID    uint        `json:"id"`
	Email string      `json:"email"`
	Data  interface{} `json:"data"`
}

type RestaurantDResponse struct {
	Name           string    `json:"name"`
	Open           time.Time `json:"open"`
	Close          time.Time `json:"close"`
	Price          int       `json:"price"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	City           string    `json:"city"`
	Address        string    `json:"address"`
	PhoneNumber    string    `json:"phone"`
	ProfilePicture string    `json:"profile_picture"`
	Seats          int       `json:"json"`
	Description    string    `json:"description"`
	Status         string    `json:"status"`
}
