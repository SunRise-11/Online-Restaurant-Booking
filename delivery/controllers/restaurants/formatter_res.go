package restaurants

type LoginResponseFormat struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type RestaurantsResponseFormat struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type RestaurantDetailResponseFormat struct {
	ID             uint    `json:"id"`
	Status         string  `json:"status"`
	ProfilePicture string  `json:"profile_picture"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Rating         int     `json:"rating"`
	Open           string  `json:"open"`
	Close          string  `json:"close"`
	Open_Hour      string  `json:"open_hour"`
	Close_Hour     string  `json:"close_hour"`
	Address        string  `json:"address"`
	City           string  `json:"city"`
	PhoneNumber    string  `json:"phone"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	Seats          int     `json:"seats"`
	Price          int     `json:"price"`
}

type RestaurantResponseFormat struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone"`
	Status      string `json:"status"`
}
