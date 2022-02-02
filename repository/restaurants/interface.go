package restaurants

import "Restobook/entities"

type RestaurantsInterface interface {
	Register(newRestaurant entities.Restaurant) (entities.Restaurant, error)
	LoginRestaurant(email, password string) (entities.Restaurant, error)
	Get(restaurantId uint) (entities.Restaurant, error)
	Update(restaurantId uint, newRestaurant entities.Restaurant) (entities.Restaurant, error)
	Delete(restaurantId uint) (entities.Restaurant, error)
}
