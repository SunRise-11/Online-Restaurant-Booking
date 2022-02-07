package restaurants

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
	ID             uint    `json:"id"`
	Name           string  `json:"name"`
	Open           string  `json:"open"`
	Close          string  `json:"close"`
	Price          int     `json:"price"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	City           string  `json:"city"`
	Address        string  `json:"address"`
	PhoneNumber    string  `json:"phone"`
	ProfilePicture string  `json:"profile_picture"`
	Seats          int     `json:"json"`
	Description    string  `json:"description"`
	Status         string  `json:"status"`
}

type RestaurantDApproveResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone"`
	Status      string `json:"status"`
}
