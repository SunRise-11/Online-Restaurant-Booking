package ratings

type PostRatingRequest struct {
	RestaurantDetailID int    `json:"restaurantdetail_id"`
	Rating             int    `json:"rating"`
	Comment            string `json:"comment"`
}

type UpdateRatingRequest struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}
