package ratings

type RatingResponse struct {
	RestaurantDetailID int    `json:"restaurantdetail_id"`
	UserID             int    `json:"user_id"`
	Username           string `json:"username"`
	Rating             int    `json:"rating"`
	Comment            string `json:"comment"`
}

type RatingResponseFormat struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
