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
	ProfilePicture string  `json:"profile_picture" form:"profile_picture"`
	Name           string  `json:"name" form:"name"`
	Description    string  `json:"description" form:"description"`
	Open           string  `json:"open" form:"open"`
	Close          string  `json:"close" form:"close"`
	Open_Hour      string  `json:"open_hour" form:"open_hour"`
	Close_Hour     string  `json:"close_hour" form:"close_hour"`
	Address        string  `json:"address" form:"address"`
	City           string  `json:"city" form:"city"`
	PhoneNumber    string  `json:"phone" form:"phone"`
	Latitude       float64 `json:"latitude" form:"latitude"`
	Longitude      float64 `json:"longitude" form:"longitude"`
	Seats          int     `json:"seats" form:"seats"`
	Price          int     `json:"price" form:"price"`
}

type UpdateRestaurantDetailRequestFormat struct {
	Open           string `json:"open" form:"open"`
	Close          string `json:"close" form:"close"`
	Open_Hour      string `json:"open_hour" form:"open_hour"`
	Close_Hour     string `json:"close_hour" form:"close_hour"`
	Price          int    `json:"price" form:"price"`
	PhoneNumber    string `json:"phone" form:"phone"`
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
	Seats          int    `json:"seats" form:"seats"`
	Description    string `json:"description" form:"description"`
}

type ApproveRequestFormat struct {
	ID     uint   `json:"resto_id"`
	Status string `json:"status"`
}

type DeleteRequestFormat struct {
	ID uint `json:"resto_id"`
}
