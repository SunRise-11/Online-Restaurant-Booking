package restaurants

import (
	"Restobook/delivery/common"
	"Restobook/entities"
	"fmt"
	"strconv"
	"strings"

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

		for i := 0; i < len(restaurantD); i++ {
			openDay := strings.Split(restaurantD[i].Open, ",")
			closeDay := strings.Split(restaurantD[i].Close, ",")
			openStr := ""
			closeStr := ""
			for j := 0; j < len(openDay); j++ {
				for k := 0; k < len(common.Daytoint); k++ {
					if openDay[j] == strconv.Itoa(common.Daytoint[k].No) {
						openStr += fmt.Sprintf("%v,", common.Daytoint[k].Day)
					}
				}
			}
			for l := 0; l < len(closeDay); l++ {
				for m := 0; m < len(common.Daytoint); m++ {
					if closeDay[l] == strconv.Itoa(common.Daytoint[m].No) {
						closeStr += fmt.Sprintf("%v,", common.Daytoint[m].Day)
					}
				}
			}
			restaurantD[i].Open = openStr
			restaurantD[i].Close = closeStr
		}

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

func (rr *RestaurantRepository) GetsByOpen(open int) ([]entities.RestaurantDetail, error) {
	restaurantD := []entities.RestaurantDetail{}
	// newrestaurantD := []entities.RestaurantDetail{}
	fmt.Println("=>open", open)
	openstr := strconv.Itoa(open)

	if err := rr.db.Where("status=? AND open LIKE ?", "OPEN", "%"+openstr+"%").Find(&restaurantD).Error; err != nil {
		return restaurantD, err
	} else {

		for i := 0; i < len(restaurantD); i++ {
			openDay := strings.Split(restaurantD[i].Open, ",")
			closeDay := strings.Split(restaurantD[i].Close, ",")
			openStr := ""
			closeStr := ""

			for j := 0; j < len(openDay); j++ {
				for k := 0; k < len(common.Daytoint); k++ {
					if openDay[j] == strconv.Itoa(common.Daytoint[k].No) {
						openStr += fmt.Sprintf("%v,", common.Daytoint[k].Day)
					}
				}
			}
			for l := 0; l < len(closeDay); l++ {
				for m := 0; m < len(common.Daytoint); m++ {
					if closeDay[l] == strconv.Itoa(common.Daytoint[m].No) {
						closeStr += fmt.Sprintf("%v,", common.Daytoint[m].Day)
					}
				}
			}
			restaurantD[i].Open = openStr
			restaurantD[i].Close = closeStr

		}

		return restaurantD, nil
	}

}

func (rr *RestaurantRepository) Gets() ([]entities.RestaurantDetail, error) {
	restaurantD := []entities.RestaurantDetail{}
	if err := rr.db.Where("status=?", "OPEN").Find(&restaurantD).Error; err != nil {
		return restaurantD, err
	} else {

		fmt.Println("===> Semua resto yang open", restaurantD)
		for i := 0; i < len(restaurantD); i++ {
			openDay := strings.Split(restaurantD[i].Open, ",")
			closeDay := strings.Split(restaurantD[i].Close, ",")
			openStr := ""
			closeStr := ""
			openH := strings.Split(restaurantD[i].Open_Hour, ":")
			closeH := strings.Split(restaurantD[i].Close_Hour, ":")

			openHHour := openH[0]
			openHMinute := openH[1]

			closeHHour := closeH[0]
			closeHMinute := closeH[1]

			for j := 0; j < len(openDay); j++ {
				for k := 0; k < len(common.Daytoint); k++ {
					if openDay[j] == strconv.Itoa(common.Daytoint[k].No) {
						openStr += fmt.Sprintf("%v,", common.Daytoint[k].Day)
					}
				}
			}
			// fmt.Println("openSTR", openStr)
			for l := 0; l < len(closeDay); l++ {
				for m := 0; m < len(common.Daytoint); m++ {
					if closeDay[l] == strconv.Itoa(common.Daytoint[m].No) {
						closeStr += fmt.Sprintf("%v,", common.Daytoint[m].Day)
					}
				}
			}
			// fmt.Println("closeSTR", closeStr)

			restaurantD[i].Open = openStr
			restaurantD[i].Close = closeStr
			restaurantD[i].Open_Hour = openHHour + ":" + openHMinute
			restaurantD[i].Close_Hour = closeHHour + ":" + closeHMinute

		}

		return restaurantD, nil
	}

}

func (rr *RestaurantRepository) GetExistSeat(restauranId uint, date_time string) ([]entities.Transaction, int, error) {
	transactions := []entities.Transaction{}
	result := 0
	if err := rr.db.Where("restaurant_id=? AND date_time = ?", restauranId, date_time).Find(&transactions).Error; err != nil {
		return transactions, result, err
	} else {
		if err := rr.db.Model(&entities.Transaction{}).Select("sum(persons) as total").Where("date_time=?", date_time).Where("restaurant_id=?", restauranId).Find(&result).Error; err != nil {
			return transactions, result, err
		} else {
			return transactions, result, nil
		}
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

	openDay := strings.Split(updateRestaurantD.Open, ",")
	closeDay := strings.Split(updateRestaurantD.Close, ",")
	openInt := ""
	closeInt := ""
	for j := 0; j < len(openDay); j++ {
		for k := 0; k < len(common.Daytoint); k++ {
			if openDay[j] == common.Daytoint[k].Day {
				openInt += fmt.Sprintf("%v,", common.Daytoint[k].No)
			}
		}
	}

	for j := 0; j < len(closeDay); j++ {
		for k := 0; k < len(common.Daytoint); k++ {
			if closeDay[j] == common.Daytoint[k].Day {
				closeInt += fmt.Sprintf("%v,", common.Daytoint[k].No)
			}
		}
	}

	parsingint := entities.RestaurantDetail{
		ID:             restaurantId,
		Name:           updateRestaurantD.Name,
		Open:           openInt,
		Close:          closeInt,
		Open_Hour:      updateRestaurantD.Open_Hour,
		Close_Hour:     updateRestaurantD.Close_Hour,
		Price:          updateRestaurantD.Price,
		Latitude:       updateRestaurantD.Latitude,
		Longitude:      updateRestaurantD.Longitude,
		City:           updateRestaurantD.City,
		Address:        updateRestaurantD.Address,
		PhoneNumber:    updateRestaurantD.PhoneNumber,
		ProfilePicture: updateRestaurantD.ProfilePicture,
		Seats:          updateRestaurantD.Seats,
		Description:    updateRestaurantD.Description,
		Status:         "Waiting for approval",
	}

	parsingstring := entities.RestaurantDetail{
		ID:             restaurantId,
		Name:           updateRestaurantD.Name,
		Open:           updateRestaurantD.Open,
		Close:          updateRestaurantD.Close,
		Open_Hour:      updateRestaurantD.Open_Hour,
		Close_Hour:     updateRestaurantD.Close_Hour,
		Price:          updateRestaurantD.Price,
		Latitude:       updateRestaurantD.Latitude,
		Longitude:      updateRestaurantD.Longitude,
		City:           updateRestaurantD.City,
		Address:        updateRestaurantD.Address,
		PhoneNumber:    updateRestaurantD.PhoneNumber,
		ProfilePicture: updateRestaurantD.ProfilePicture,
		Seats:          updateRestaurantD.Seats,
		Description:    updateRestaurantD.Description,
		Status:         "Waiting for approval",
	}

	if err := rr.db.First(&restaurant, "id=?", restaurantId).Error; err != nil {
		return restaurantD, err
	}

	rr.db.First(&restaurantD, "id=?", restaurant.RestaurantDetailID)
	rr.db.Model(&restaurantD).Updates(parsingint)
	return parsingstring, nil

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
