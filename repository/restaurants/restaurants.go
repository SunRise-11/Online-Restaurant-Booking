package restaurants

import (
	"Restobook/entities"
	"time"

	"gorm.io/gorm"
)

type RestaurantRepository struct {
	db *gorm.DB
}

func NewRestaurantsRepo(db *gorm.DB) *RestaurantRepository {
	return &RestaurantRepository{db: db}
}

func (rr *RestaurantRepository) LoginRestaurant(email, password string) (entities.Restaurant, error) {
	var restaurant entities.Restaurant

	if err := rr.db.Where("Email = ? AND Password=?", email, password).First(&restaurant).Error; err != nil {
		return restaurant, err
	}

	return restaurant, nil
}
func (rr *RestaurantRepository) Register(newRestaurant entities.Restaurant) (entities.Restaurant, error) {
	now := time.Now()
	restaurantD := entities.RestaurantDetail{
		Open:  now,
		Close: now,
	}

	if err := rr.db.Save(&restaurantD).Error; err != nil {
		return newRestaurant, err
	}

	newRestaurant.ID = restaurantD.ID
	newRestaurant.RestaurantDetailID = restaurantD.ID

	if err := rr.db.Save(&newRestaurant).Error; err != nil {
		return newRestaurant, err
	}

	return newRestaurant, nil
}
func (rr *RestaurantRepository) Delete(restauranId uint) (entities.Restaurant, error) {
	restaurant := entities.Restaurant{}

	if err := rr.db.First(&restaurant, "id=?", restauranId).Error; err != nil {
		return restaurant, err
	}

	rr.db.Delete(&restaurant)

	return restaurant, nil
}
func (rr *RestaurantRepository) Update(restauranId uint, newRestaurant entities.Restaurant) (entities.Restaurant, error) {
	restaurant := entities.Restaurant{}

	if err := rr.db.First(&restaurant, "id=?", restauranId).Error; err != nil {
		return restaurant, err
	}

	rr.db.Model(&restaurant).Updates(newRestaurant)

	return restaurant, nil
}
func (rr *RestaurantRepository) Get(restauranId uint) (entities.Restaurant, error) {
	restaurant := entities.Restaurant{}

	if err := rr.db.First(&restaurant, restauranId).Error; err != nil {
		return restaurant, err
	}

	return restaurant, nil
}
