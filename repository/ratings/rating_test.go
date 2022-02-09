package ratings

import (
	"Restobook/configs"
	"Restobook/entities"
	"Restobook/repository/restaurants"
	"Restobook/repository/transactions"
	"Restobook/repository/users"
	"Restobook/utils"
	"crypto/sha256"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateRating(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := users.NewUsersRepo(db)
	restaurantRepo := restaurants.NewRestaurantsRepo(db)
	transactionRepo := transactions.NewTransactionRepo(db)
	ratingRepo := NewRatingsRepo(db)

	//CREATE USER
	var userDummy entities.User
	hash := sha256.Sum256([]byte("resto123"))
	password := fmt.Sprintf("%x", hash[:])
	userDummy.Email = "restaurant1@outlook.my"
	userDummy.Password = password
	userRepo.Register(userDummy)

	//CREATE RESTAURANT
	var restaurantDummy entities.Restaurant
	hash = sha256.Sum256([]byte("resto123"))
	password = fmt.Sprintf("%x", hash[:])
	restaurantDummy.Email = "restaurant1@outlook.my"
	restaurantDummy.Password = password
	restaurantRepo.Register(restaurantDummy)

	//CREATE TRANSACTION
	now := time.Now()
	var transactionDummy entities.Transaction
	transactionDummy.RestaurantID = 1
	transactionDummy.UserID = 1
	transactionDummy.Status = "SUCCESS"
	transactionDummy.DateTime = now
	transactionRepo.Create(transactionDummy)

	t.Run("Create Rating 1", func(t *testing.T) {
		var newRating entities.Rating
		newRating.UserID = 1
		newRating.RestaurantDetailID = 1
		newRating.Comment = "Mantap"
		newRating.Rating = 5
		res, err := ratingRepo.Create(newRating)
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

	t.Run("FALSE Create Rating 1", func(t *testing.T) {

		var newRating entities.Rating
		newRating.UserID = 1
		newRating.RestaurantDetailID = 1
		newRating.Comment = "Mantap"
		newRating.Rating = 5

		res, err := ratingRepo.Create(newRating)
		assert.Equal(t, res, res)
		assert.Error(t, err)
	})
}

func TestCanGiveRating(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := users.NewUsersRepo(db)
	restaurantRepo := restaurants.NewRestaurantsRepo(db)
	transactionRepo := transactions.NewTransactionRepo(db)
	ratingRepo := NewRatingsRepo(db)

	//CREATE USER
	var userDummy entities.User
	hash := sha256.Sum256([]byte("resto123"))
	password := fmt.Sprintf("%x", hash[:])
	userDummy.Email = "restaurant1@outlook.my"
	userDummy.Password = password
	userRepo.Register(userDummy)

	//CREATE RESTAURANT
	var restaurantDummy entities.Restaurant
	hash = sha256.Sum256([]byte("resto123"))
	password = fmt.Sprintf("%x", hash[:])
	restaurantDummy.Email = "restaurant1@outlook.my"
	restaurantDummy.Password = password
	restaurantRepo.Register(restaurantDummy)

	//CREATE TRANSACTION
	now := time.Now()
	var transactionDummy entities.Transaction
	transactionDummy.RestaurantID = 1
	transactionDummy.UserID = 1
	transactionDummy.Status = "SUCCESS"
	transactionDummy.DateTime = now
	transactionRepo.Create(transactionDummy)

	t.Run("Can Give Rating", func(t *testing.T) {
		res, err := ratingRepo.IsCanGiveRating(1, 1)
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

	t.Run("FALSE Can Give Rating", func(t *testing.T) {
		res, err := ratingRepo.IsCanGiveRating(2, 2)
		assert.Equal(t, res, res)
		assert.Error(t, err)
	})
}

func TestUpdateRating(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := users.NewUsersRepo(db)
	restaurantRepo := restaurants.NewRestaurantsRepo(db)
	transactionRepo := transactions.NewTransactionRepo(db)
	ratingRepo := NewRatingsRepo(db)

	//CREATE USER
	var userDummy entities.User
	hash := sha256.Sum256([]byte("resto123"))
	password := fmt.Sprintf("%x", hash[:])
	userDummy.Email = "restaurant1@outlook.my"
	userDummy.Password = password
	userRepo.Register(userDummy)

	//CREATE RESTAURANT
	var restaurantDummy entities.Restaurant
	hash = sha256.Sum256([]byte("resto123"))
	password = fmt.Sprintf("%x", hash[:])
	restaurantDummy.Email = "restaurant1@outlook.my"
	restaurantDummy.Password = password
	restaurantRepo.Register(restaurantDummy)

	//CREATE TRANSACTION
	now := time.Now()
	var transactionDummy entities.Transaction
	transactionDummy.RestaurantID = 1
	transactionDummy.UserID = 1
	transactionDummy.Status = "SUCCESS"
	transactionDummy.DateTime = now
	transactionRepo.Create(transactionDummy)

	//CREATE RATING
	var ratingDummy entities.Rating
	ratingDummy.UserID = 1
	ratingDummy.RestaurantDetailID = 1
	ratingDummy.Comment = "Mantap Sekali"
	ratingDummy.Rating = 5
	ratingRepo.Create(ratingDummy)

	t.Run("Update Rating", func(t *testing.T) {
		var newRating entities.Rating
		newRating.UserID = 1
		newRating.RestaurantDetailID = 1
		newRating.Comment = "Mantap Sekali"
		newRating.Rating = 5
		res, err := ratingRepo.Update(newRating)
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

	t.Run("FALSE Update Rating", func(t *testing.T) {
		var newRating entities.Rating
		newRating.UserID = 10
		newRating.RestaurantDetailID = 110
		newRating.Comment = "Mantap Sekali"
		newRating.Rating = 5
		res, err := ratingRepo.Update(newRating)
		assert.Equal(t, res, res)
		assert.Error(t, err)
	})
}

func TestDeleteRating(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := users.NewUsersRepo(db)
	restaurantRepo := restaurants.NewRestaurantsRepo(db)
	transactionRepo := transactions.NewTransactionRepo(db)
	ratingRepo := NewRatingsRepo(db)

	//CREATE USER
	var userDummy entities.User
	hash := sha256.Sum256([]byte("resto123"))
	password := fmt.Sprintf("%x", hash[:])
	userDummy.Email = "restaurant1@outlook.my"
	userDummy.Password = password
	userRepo.Register(userDummy)

	//CREATE RESTAURANT
	var restaurantDummy entities.Restaurant
	hash = sha256.Sum256([]byte("resto123"))
	password = fmt.Sprintf("%x", hash[:])
	restaurantDummy.Email = "restaurant1@outlook.my"
	restaurantDummy.Password = password
	restaurantRepo.Register(restaurantDummy)

	//CREATE TRANSACTION
	now := time.Now()
	var transactionDummy entities.Transaction
	transactionDummy.RestaurantID = 1
	transactionDummy.UserID = 1
	transactionDummy.Status = "SUCCESS"
	transactionDummy.DateTime = now
	transactionRepo.Create(transactionDummy)

	//CREATE RATING
	var ratingDummy entities.Rating
	ratingDummy.UserID = 1
	ratingDummy.RestaurantDetailID = 1
	ratingDummy.Comment = "Mantap Sekali"
	ratingDummy.Rating = 5
	ratingRepo.Create(ratingDummy)

	t.Run("Delete Rating", func(t *testing.T) {
		res, err := ratingRepo.Delete(1, 1)
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

	t.Run("FALSE Delete Rating", func(t *testing.T) {
		res, err := ratingRepo.Delete(1, 1)
		assert.Equal(t, res, res)
		assert.Error(t, err)
	})
}
