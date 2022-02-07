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

func TestRatingsRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := users.NewUsersRepo(db)
	restaurantRepo := restaurants.NewRestaurantsRepo(db)
	transactionRepo := transactions.NewTransactionRepo(db)
	ratingRepo := NewRatingsRepo(db)

	t.Run("Register User 1", func(t *testing.T) {
		hash := sha256.Sum256([]byte("resto123"))
		password := fmt.Sprintf("%x", hash[:])
		var newUser entities.User
		newUser.Email = "restaurant1@outlook.my"
		newUser.Password = password

		res, err := userRepo.Register(newUser)
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

	t.Run("Register Restaurant 1", func(t *testing.T) {
		hash := sha256.Sum256([]byte("resto123"))
		password := fmt.Sprintf("%x", hash[:])
		var newRestaurant entities.Restaurant
		newRestaurant.Email = "restaurant1@outlook.my"
		newRestaurant.Password = password

		res, err := restaurantRepo.Register(newRestaurant)
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

	t.Run("Create Transactions 1", func(t *testing.T) {
		now := time.Now()
		var newTransaction entities.Transaction
		newTransaction.RestaurantID = 1
		newTransaction.UserID = 1
		newTransaction.Status = "SUCCESS"
		newTransaction.DateTime = now

		res, err := transactionRepo.Create(newTransaction)
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

	t.Run("Create Rating 1", func(t *testing.T) {
		var newRating entities.Rating
		newRating.UserID = 1
		newRating.RestaurantID = 1
		newRating.Comment = "Mantap"
		newRating.Rating = 5
		res, err := ratingRepo.Create(newRating)
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

	t.Run("Can Give Rating", func(t *testing.T) {
		res, err := ratingRepo.IsCanGiveRating(1, 1)
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

	t.Run("Update Rating", func(t *testing.T) {
		var newRating entities.Rating
		newRating.UserID = 1
		newRating.RestaurantID = 1
		newRating.Comment = "Mantap Sekali"
		newRating.Rating = 5
		res, err := ratingRepo.Update(newRating)
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

	t.Run("Delete Rating", func(t *testing.T) {
		res, err := ratingRepo.Delete(1, 1)
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

}

func TestFalseRatingsRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	ratingRepo := NewRatingsRepo(db)

	t.Run("FALSE Create Rating 1", func(t *testing.T) {

		var newRating entities.Rating
		newRating.UserID = 1
		newRating.RestaurantID = 1
		newRating.Comment = "Mantap"
		newRating.Rating = 5

		res, err := ratingRepo.Create(newRating)
		assert.Equal(t, res, res)
		assert.Error(t, err)
	})

	t.Run("FALSE Can Give Rating", func(t *testing.T) {
		res, err := ratingRepo.IsCanGiveRating(1, 1)
		assert.Equal(t, res, res)
		assert.Error(t, err)
	})

	t.Run("FALSE Update Rating", func(t *testing.T) {
		var newRating entities.Rating
		newRating.UserID = 1
		newRating.RestaurantID = 1
		newRating.Comment = "Mantap Sekali"
		newRating.Rating = 5
		res, err := ratingRepo.Update(newRating)
		assert.Equal(t, res, res)
		assert.Error(t, err)
	})

	t.Run("FALSE Delete Rating", func(t *testing.T) {
		res, err := ratingRepo.Delete(1, 1)
		assert.Equal(t, res, res)
		assert.Error(t, err)
	})

}
