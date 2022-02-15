package restaurants

import (
	"Restobook/delivery/helpers"
	"Restobook/entities"
	"errors"
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

	if err := rr.db.Save(&newRestaurant).Error; err != nil || newRestaurant.ID == 0 {
		return newRestaurant, errors.New("FAILED REGISTER")
	} else {
		return newRestaurant, nil
	}

}

func (rr *RestaurantRepository) Login(email, password string) (entities.Restaurant, error) {
	var restaurant entities.Restaurant

	if err := rr.db.Where("Email = ? AND Password=?", email, password).First(&restaurant).Error; err != nil || restaurant.ID == 0 {
		return restaurant, errors.New("FAILED LOGIN")
	} else {
		return restaurant, nil
	}

}

func (rr *RestaurantRepository) Update(restaurantId uint, updateRestaurant entities.Restaurant) (entities.Restaurant, error) {
	restaurant := entities.Restaurant{}

	if err := rr.db.First(&restaurant, "id=?", restaurantId).Error; err != nil || restaurant.ID == 0 {
		return restaurant, errors.New("FAILED UPDATE")
	} else {
		rr.db.Model(&restaurant).Updates(updateRestaurant)
		return restaurant, nil
	}

}

func (rr *RestaurantRepository) Get(restaurantId uint) (entities.Restaurant, entities.RestaurantDetail, error) {
	restaurant := entities.Restaurant{}
	restaurantD := entities.RestaurantDetail{}

	if err := rr.db.Preload("RestaurantDetail.Rating").First(&restaurant, restaurantId).Error; err != nil || restaurant.ID == 0 {

		return restaurant, restaurantD, errors.New("FAILED GET")

	} else {

		rr.db.Preload("Rating").First(&restaurantD, restaurant.RestaurantDetailID)
		openDay := strings.Split(restaurantD.Open, ",")
		closeDay := strings.Split(restaurantD.Close, ",")

		openStr, closeStr, _ := helpers.NumberToDayConverter(openDay, closeDay)

		restaurantD.Open = openStr
		restaurantD.Close = closeStr

		return restaurant, restaurantD, nil
	}

}

func (rr *RestaurantRepository) CreateDetail(restaurantId uint, updateRestaurantD entities.RestaurantDetail) (entities.RestaurantDetail, error) {
	restaurant := entities.Restaurant{}
	restaurantD := entities.RestaurantDetail{}

	openDay := strings.Split(updateRestaurantD.Open, ",")
	closeDay := strings.Split(updateRestaurantD.Close, ",")

	openInt, closeInt, _ := helpers.DaytoNumberConverter(openDay, closeDay)

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

	if err := rr.db.First(&restaurant, "id=?", restaurantId).Error; err != nil || restaurant.ID == 0 {
		return restaurantD, errors.New("FAILED CREATE DETAIL")
	} else {
		rr.db.First(&restaurantD, "id=?", restaurant.RestaurantDetailID)
		rr.db.Model(&restaurantD).Updates(parsingint)
		return parsingstring, nil
	}

}

func (rr *RestaurantRepository) UpdateDetail(restaurantId uint, updateRestaurantD entities.RestaurantDetail) (entities.RestaurantDetail, error) {
	restaurant := entities.Restaurant{}
	restaurantD := entities.RestaurantDetail{}

	openDay := strings.Split(updateRestaurantD.Open, ",")
	closeDay := strings.Split(updateRestaurantD.Close, ",")

	openInt, closeInt, _ := helpers.DaytoNumberConverter(openDay, closeDay)

	if err := rr.db.First(&restaurant, "id=?", restaurantId).Error; err != nil || restaurant.ID == 0 {
		return restaurantD, errors.New("FAILED UPDATE DETAIL")
	} else {
		rr.db.First(&restaurantD, "id=?", restaurant.RestaurantDetailID)

		parsingint := entities.RestaurantDetail{
			ID:             restaurantId,
			Name:           restaurantD.Name,
			Open:           openInt,
			Close:          closeInt,
			Open_Hour:      updateRestaurantD.Open_Hour,
			Close_Hour:     updateRestaurantD.Close_Hour,
			Price:          updateRestaurantD.Price,
			Latitude:       restaurantD.Latitude,
			Longitude:      restaurantD.Longitude,
			City:           restaurantD.City,
			Address:        restaurantD.Address,
			PhoneNumber:    updateRestaurantD.PhoneNumber,
			ProfilePicture: updateRestaurantD.ProfilePicture,
			Seats:          updateRestaurantD.Seats,
			Description:    updateRestaurantD.Description,
			Status:         "Waiting for approval",
		}

		parsingstring := entities.RestaurantDetail{
			ID:             restaurantId,
			Name:           restaurantD.Name,
			Open:           updateRestaurantD.Open,
			Close:          updateRestaurantD.Close,
			Open_Hour:      updateRestaurantD.Open_Hour,
			Close_Hour:     updateRestaurantD.Close_Hour,
			Price:          updateRestaurantD.Price,
			Latitude:       restaurantD.Latitude,
			Longitude:      restaurantD.Longitude,
			City:           restaurantD.City,
			Address:        restaurantD.Address,
			PhoneNumber:    updateRestaurantD.PhoneNumber,
			ProfilePicture: updateRestaurantD.ProfilePicture,
			Seats:          updateRestaurantD.Seats,
			Description:    updateRestaurantD.Description,
			Status:         "Waiting for approval",
		}

		rr.db.Model(&restaurantD).Updates(parsingint)
		return parsingstring, nil
	}

}

func (rr *RestaurantRepository) GetsWaiting() ([]entities.RestaurantDetail, error) {
	restaurantD := []entities.RestaurantDetail{}

	if err := rr.db.Where("status=?", "Waiting for approval").Find(&restaurantD).Error; err != nil || len(restaurantD) == 0 {
		return restaurantD, errors.New("FAILED GETS WAITING")
	} else {

		for i := 0; i < len(restaurantD); i++ {
			openDay := strings.Split(restaurantD[i].Open, ",")
			closeDay := strings.Split(restaurantD[i].Close, ",")

			openStr, closeStr, _ := helpers.NumberToDayConverter(openDay, closeDay)

			restaurantD[i].Open = openStr
			restaurantD[i].Close = closeStr
		}

		return restaurantD, nil
	}
}

func (rr *RestaurantRepository) Approve(restaurantId uint, status string) (entities.RestaurantDetail, error) {
	restaurantD := entities.RestaurantDetail{}

	if err := rr.db.Preload("Restaurant").First(&restaurantD, "id=?", restaurantId).Error; err != nil || restaurantD.ID == 0 {
		return restaurantD, errors.New("FAILED APPROVE")
	} else {
		updateStatus := entities.RestaurantDetail{
			Status: status,
		}
		rr.db.Model(&restaurantD).Updates(updateStatus)
		return restaurantD, nil
	}

}

func (rr *RestaurantRepository) Gets() ([]entities.RestaurantDetail, error) {
	restaurantD := []entities.RestaurantDetail{}
	if err := rr.db.Preload("Rating").Where("status=?", "OPEN").Find(&restaurantD).Error; err != nil || len(restaurantD) == 0 {
		return restaurantD, errors.New("FAILED GETS")
	} else {

		for i := 0; i < len(restaurantD); i++ {
			openDay := strings.Split(restaurantD[i].Open, ",")
			closeDay := strings.Split(restaurantD[i].Close, ",")

			openH := strings.Split(restaurantD[i].Open_Hour, ":")
			closeH := strings.Split(restaurantD[i].Close_Hour, ":")

			openHHour := openH[0]
			openHMinute := openH[1]

			closeHHour := closeH[0]
			closeHMinute := closeH[1]

			openStr, closeStr, _ := helpers.NumberToDayConverter(openDay, closeDay)

			restaurantD[i].Open = openStr
			restaurantD[i].Close = closeStr
			restaurantD[i].Open_Hour = openHHour + ":" + openHMinute
			restaurantD[i].Close_Hour = closeHHour + ":" + closeHMinute

		}

		return restaurantD, nil
	}

}

func (rr *RestaurantRepository) GetsByOpen(open int) ([]entities.RestaurantDetail, error) {
	restaurantD := []entities.RestaurantDetail{}
	openstr := strconv.Itoa(open)

	if err := rr.db.Preload("Rating").Where("status=? AND open LIKE ?", "OPEN", "%"+openstr+"%").Find(&restaurantD).Error; err != nil || len(restaurantD) == 0 {
		return restaurantD, errors.New("FAILED GETS BY OPEN")
	} else {

		for i := 0; i < len(restaurantD); i++ {
			openDay := strings.Split(restaurantD[i].Open, ",")
			closeDay := strings.Split(restaurantD[i].Close, ",")

			openStr, closeStr, _ := helpers.NumberToDayConverter(openDay, closeDay)

			restaurantD[i].Open = openStr
			restaurantD[i].Close = closeStr
		}

		return restaurantD, nil
	}

}

func (rr *RestaurantRepository) GetExistSeat(restauranId uint, date_time string) ([]entities.Transaction, int, error) {
	transactions := []entities.Transaction{}
	result := 0
	if err := rr.db.Model(&entities.Transaction{}).Select("sum(persons) as total").Where("date_time=?", date_time).Where("restaurant_id=?", restauranId).Where("status=?", "Accepted").Find(&result).Error; err != nil {
		return transactions, result, errors.New("FAILED GET EXIST SEAT")
	} else {
		return transactions, result, nil
	}
}

func (rr *RestaurantRepository) Delete(restaurantId uint) (entities.Restaurant, error) {
	restaurant := entities.Restaurant{}

	if err := rr.db.First(&restaurant, "id=?", restaurantId).Delete(&restaurant).Error; err != nil || restaurant.ID == 0 {
		return restaurant, errors.New("FAILED DELETE")
	} else {
		return restaurant, nil
	}

}

func (rr *RestaurantRepository) Export(restaurantId uint, date string) ([]entities.Transaction, error) {
	transactions := []entities.Transaction{}

	if err := rr.db.Preload("Restaurant.RestaurantDetail").Where("restaurant_id=?", restaurantId).Where("date_time LIKE ?", "%"+date+"%").Find(&transactions).Error; err != nil || len(transactions) == 0 {
		return transactions, errors.New("FAILED EXPORT PDF")
	} else {
		return transactions, nil
	}

}
