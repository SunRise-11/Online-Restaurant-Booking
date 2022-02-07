package restaurants

import "Restobook/entities"

type RestaurantsInterface interface {
	Register(newRestaurant entities.Restaurant) (entities.Restaurant, error)
	LoginRestaurant(email, password string) (entities.Restaurant, error)
	GetsWaiting() ([]entities.RestaurantDetail, error)
	Approve(restaurantId uint, status string) (entities.RestaurantDetail, error)
	Get(restaurantId uint) (entities.Restaurant, entities.RestaurantDetail, error)
	GetsByOpen(open int) ([]entities.RestaurantDetail, error)
	GetExistSeat(restauranId uint, date_time string) ([]entities.Transaction, int, error)
	Gets() ([]entities.RestaurantDetail, error)
	Update(restaurantId uint, newRestaurant entities.Restaurant) (entities.Restaurant, error)
	UpdateDetail(restaurantId uint, updateRestaurantD entities.RestaurantDetail) (entities.RestaurantDetail, error)
	Delete(restaurantId uint) (entities.Restaurant, error)
}
