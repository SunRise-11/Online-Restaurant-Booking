package restaurants

type LoginRequestFormat struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequestFormat struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateRestaurantDetailRequestFormat struct {
	Status         string  `json:"status"`
	ProfilePicture string  `json:"profile_picture"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
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

type UpdateRestaurantDetailRequestFormat struct {
	Open           string `json:"open"`
	Close          string `json:"close"`
	Open_Hour      string `json:"open_hour"`
	Close_Hour     string `json:"close_hour"`
	Price          int    `json:"price"`
	PhoneNumber    string `json:"phone"`
	ProfilePicture string `json:"profile_picture"`
	Seats          int    `json:"seats"`
	Description    string `json:"description"`
}

type ApproveRequestFormat struct {
	ID     uint   `json:"resto_id"`
	Status string `json:"status"`
}

type DeleteRequestFormat struct {
	ID uint `json:"resto_id"`
}
