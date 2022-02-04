package restaurants

import (
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

var Daytoint = []struct {
	day string
	no  int
}{
	{"Monday", 0},
	{"Tuesday", 1},
	{"Wednesday", 2},
	{"Thursday", 3},
	{"Friday", 4},
	{"Saturday", 5},
	{"Sunday", 6},
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

	if err := rr.db.Not("status=?", "DISABLED").Not("status=?", "CLOSED").Find(&restaurantD).Error; err != nil {
		return restaurantD, err
	} else {

		// fmt.Println("===> semua resto", restaurantD)

		// openArray := make(map[string]int)
		// angkaOpen := 0
		// for i := 0; i < len(restaurantD); i++ {
		// 	openDay := strings.Split(restaurantD[i].Open, ",")
		// 	for j := 0; j < len(openDay); j++ {
		// 		for k := 0; k < len(daytoint); k++ {
		// 			if openDay[j] == daytoint[k].day {
		// 				openArray[openDay[j]] = daytoint[k].no
		// 			}
		// 		}
		// 	}
		// }
		// for k := 0; k < len(daytoint); k++ {
		// 	if open == daytoint[k].day {
		// 		angkaOpen = daytoint[k].no
		// 	}
		// }
		// fmt.Println("open", openArray)
		// fmt.Println("show open", angkaOpen)

		return restaurantD, nil
	}

}

func (rr *RestaurantRepository) Gets() ([]entities.RestaurantDetail, error) {
	restaurantD := []entities.RestaurantDetail{}
	newRestaurantD := []entities.RestaurantDetail{}

	if err := rr.db.Not("status=?", "DISABLED").Not("status=?", "CLOSED").Not("status=?", "Waiting for approval").Find(&restaurantD).Error; err != nil {
		return restaurantD, err
	} else {
		fmt.Println("===> Semua resto yang open", restaurantD)
		for i := 0; i < len(restaurantD); i++ {
			openDay := strings.Split(restaurantD[i].Open, ",")
			closeDay := strings.Split(restaurantD[i].Close, ",")
			openStr := ""
			closeStr := ""

			fmt.Println("open", openDay)
			fmt.Println("close", closeDay)
			for j := 0; j < len(openDay); j++ {
				for k := 0; k < len(Daytoint); k++ {
					if openDay[j] == strconv.Itoa(Daytoint[k].no) {
						openStr += fmt.Sprintf("%v,", Daytoint[k].day)
					}
				}
			}
			fmt.Println("openSTR", openStr)
			for l := 0; l < len(closeDay); l++ {
				for m := 0; m < len(Daytoint); m++ {
					if closeDay[l] == strconv.Itoa(Daytoint[m].no) {
						closeStr += fmt.Sprintf("%v,", Daytoint[m].day)
					}
				}
			}
			fmt.Println("closeSTR", closeStr)

			restaurantD[i].Open = openStr
			restaurantD[i].Close = closeStr

			fmt.Println("===> intip", newRestaurantD)

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

	openDay := strings.Split(updateRestaurantD.Open, ",")
	closeDay := strings.Split(updateRestaurantD.Close, ",")
	openInt := ""
	closeInt := ""
	for j := 0; j < len(openDay); j++ {
		for k := 0; k < len(Daytoint); k++ {
			if openDay[j] == Daytoint[k].day {
				openInt += fmt.Sprintf("%v,", Daytoint[k].no)
			}
		}
	}

	for j := 0; j < len(closeDay); j++ {
		for k := 0; k < len(Daytoint); k++ {
			if closeDay[j] == Daytoint[k].day {
				closeInt += fmt.Sprintf("%v,", Daytoint[k].no)
			}
		}
	}

	parsingint := entities.RestaurantDetail{
		Name:            updateRestaurantD.Name,
		Open:            openInt,
		Close:           closeInt,
		OperationalHour: updateRestaurantD.OperationalHour,
		Price:           updateRestaurantD.Price,
		Latitude:        updateRestaurantD.Latitude,
		Longitude:       updateRestaurantD.Longitude,
		City:            updateRestaurantD.City,
		Address:         updateRestaurantD.Address,
		PhoneNumber:     updateRestaurantD.PhoneNumber,
		ProfilePicture:  updateRestaurantD.ProfilePicture,
		Seats:           updateRestaurantD.Seats,
		Description:     updateRestaurantD.Description,
		Status:          "Waiting for approval",
	}

	if err := rr.db.First(&restaurant, "id=?", restaurantId).Error; err != nil {
		return restaurantD, err
	}

	rr.db.First(&restaurantD, "id=?", restaurant.RestaurantDetailID)
	rr.db.Model(&restaurantD).Updates(parsingint)
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
