package restaurants

import "Restobook/entities"

type RestaurantsInterface interface {
	Register(newRestaurant entities.Restaurant) (entities.Restaurant, error)
	Login(email, password string) (entities.Restaurant, error)
	Update(restaurantId uint, newRestaurant entities.Restaurant) (entities.Restaurant, error)
	Get(restaurantId uint) (entities.Restaurant, entities.RestaurantDetail, error)
	CreateDetail(restaurantId uint, updateRestaurantD entities.RestaurantDetail) (entities.RestaurantDetail, error)
	UpdateDetail(restaurantId uint, updateRestaurantD entities.RestaurantDetail) (entities.RestaurantDetail, error)
	GetsWaiting() ([]entities.RestaurantDetail, error)
	Approve(restaurantId uint, status string) (entities.RestaurantDetail, error)
	Gets() ([]entities.RestaurantDetail, error)
	GetsByOpen(open int) ([]entities.RestaurantDetail, error)
	GetExistSeat(restauranId uint, date_time string) ([]entities.Transaction, int, error)
	Delete(restaurantId uint) (entities.Restaurant, error)
}
