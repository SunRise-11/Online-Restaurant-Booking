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

func (rr *RestaurantRepository) Register(newRestaurant entities.Restaurant) (entities.Restaurant, error) {
	now := time.Now()
	restaurantD := entities.RestaurantDetail{
		Open:  now,
		Close: now,
	}
	rr.db.Save(&restaurantD)

	// if err := rr.db.Save(&restaurantD).Error; err != nil {
	// 	return newRestaurant, err
	// }

	newRestaurant.RestaurantDetailID = restaurantD.ID

	if err := rr.db.Save(&newRestaurant).Error; err != nil {
		return newRestaurant, err
	}

	return newRestaurant, nil
}

func (rr *RestaurantRepository) LoginRestaurant(email, password string) (entities.Restaurant, error) {
	var restaurant entities.Restaurant

	if err := rr.db.Where("Email = ? AND Password=?", email, password).First(&restaurant).Error; err != nil {
		return restaurant, err
	}

	return restaurant, nil
}

func (rr *RestaurantRepository) Get(restaurantId uint) (entities.Restaurant, entities.RestaurantDetail, error) {
	restaurant := entities.Restaurant{}
	restaurantD := entities.RestaurantDetail{}

	if err := rr.db.First(&restaurant, restaurantId).Error; err != nil {
		return restaurant, restaurantD, err
	} else {

		rr.db.First(&restaurantD, restaurant.RestaurantDetailID)

		return restaurant, restaurantD, nil
	}

}

func (rr *RestaurantRepository) Update(restaurantId uint, updateRestaurant entities.Restaurant) (entities.Restaurant, error) {
	restaurant := entities.Restaurant{}

	if err := rr.db.First(&restaurant, "id=?", restaurantId).Error; err != nil {
		return restaurant, err
	}

	rr.db.Model(&restaurant).Updates(updateRestaurant)

	return restaurant, nil
}

func (rr *RestaurantRepository) UpdateDetail(restaurantId uint, updateRestaurantD entities.RestaurantDetail) (entities.RestaurantDetail, error) {
	restaurant := entities.Restaurant{}
	restaurantD := entities.RestaurantDetail{}

	if err := rr.db.First(&restaurant, "id=?", restaurantId).Error; err != nil {
		return restaurantD, err
	}

	rr.db.First(&restaurantD, "id=?", restaurant.RestaurantDetailID)
	rr.db.Model(&restaurantD).Updates(updateRestaurantD)
	return restaurantD, nil

}

func (rr *RestaurantRepository) Delete(restaurantId uint) (entities.Restaurant, error) {
	restaurant := entities.Restaurant{}

	// if err := rr.db.First(&restaurant, "id=?", restaurantId).Error; err != nil {
	// 	return restaurant, err
	// }

	if err := rr.db.First(&restaurant, "id=?", restaurantId).Delete(&restaurant).Error; err != nil {
		return restaurant, err
	} else {
		return restaurant, nil
	}

}
