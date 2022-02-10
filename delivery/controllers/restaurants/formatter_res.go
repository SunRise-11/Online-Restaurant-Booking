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

type ExportPDFResponseFormat struct {
	Date    string      `json:"date"`
	Name    string      `json:"name"`
	Address string      `json:"address"`
	Orders  interface{} `json:"Orders"`
	Seats   interface{} `json:"Seats"`
	Total   interface{} `json:"Total"`
}

type ExportPDF_Order_Response struct {
	Number_of_success_orders  string `json:"number_of_success_orders"`
	Number_of_fail_orders     string `json:"number_of_fail_orders"`
	Number_of_cancel_orders   string `json:"number_of_cancel_orders"`
	Total_orders              string `json:"total_orders"`
	Number_of_rejected_orders string `json:"number_of_rejected_orders"`
}

type ExportPDF_Seats_Response struct {
	Number_of_success_seats  string `json:"number_of_success_seats"`
	Number_of_fail_seats     string `json:"number_of_fail_seats"`
	Number_of_cancel_seats   string `json:"number_of_cancel_seats"`
	Total_seats              string `json:"total_seats"`
	Number_of_rejected_seats string `json:"number_of_rejected_seats"`
}

type ExportPDF_Total_Response struct {
	Number_of_success_total string `json:"number_of_success_total"`
	Number_of_fail_total    string `json:"number_of_fail_total"`
	Number_of_cancel_total  string `json:"number_of_cancel_total"`
	Total                   string `json:"total"`
}
