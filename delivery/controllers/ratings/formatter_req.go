package ratings

type PostRatingRequest struct {
	RestaurantID int    `json:"restaurant_id"`
	Rating       int    `json:"rating"`
	Comment      string `json:"comment"`
}

type UpdateRatingRequest struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}
