package restaurants

import (
	"Restobook/entities"
	"fmt"
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
	restaurantD := entities.RestaurantDetail{
		Name: "Restaurant Name",
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

func (rr *RestaurantRepository) GetsWaiting() ([]entities.RestaurantDetail, error) {
	restaurantD := []entities.RestaurantDetail{}

	if err := rr.db.Where("status=?", "Waiting for approval").Find(&restaurantD).Error; err != nil {
		return restaurantD, err
	} else {
		return restaurantD, nil
	}
}

func (rr *RestaurantRepository) Approve(restaurantId uint, status string) (entities.RestaurantDetail, error) {
	restaurantD := entities.RestaurantDetail{}

	if err := rr.db.First(&restaurantD, "id=?", restaurantId).Error; err != nil {
		return restaurantD, err
	} else {
		updateStatus := entities.RestaurantDetail{
			Status: status,
		}
		rr.db.Model(&restaurantD).Updates(updateStatus)
		return restaurantD, nil
	}

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

func (rr *RestaurantRepository) GetsByOpen(open, oh string) ([]entities.RestaurantDetail, error) {
	restaurantD := []entities.RestaurantDetail{}

	if err := rr.db.Not("status=?", "DISABLED").Not("status=?", "CLOSED").Where("open=? AND operational_hour lIKE ?", open, oh+"%").Find(&restaurantD).Error; err != nil {
		return restaurantD, err
	} else {

		now := time.Now()
		fmt.Println("now", now)
		fmt.Println("today", now.Day())
		fmt.Println("month", now.Month())
		fmt.Println("year", now.Year())

		for i := 0; i < len(restaurantD); i++ {
			fmt.Println("open", restaurantD[i].Open)
			fmt.Println("close", restaurantD[i].Close)
			fmt.Println("oh", restaurantD[i].OperationalHour)
		}

		return restaurantD, nil
	}

}

func (rr *RestaurantRepository) Gets() ([]entities.RestaurantDetail, error) {
	restaurantD := []entities.RestaurantDetail{}

	if err := rr.db.Not("status=?", "DISABLED").Not("status=?", "CLOSED").Not("status=?", "Waiting for approval").Find(&restaurantD).Error; err != nil {
		return restaurantD, err
	} else {

		now := time.Now()
		fmt.Println("now", now)
		fmt.Println("today", now.Day())
		fmt.Println("month", now.Month())
		fmt.Println("year", now.Year())

		for i := 0; i < len(restaurantD); i++ {
			fmt.Println("open", restaurantD[i].Open)
			fmt.Println("close", restaurantD[i].Close)
			fmt.Println("oh", restaurantD[i].OperationalHour)
		}

		return restaurantD, nil
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
