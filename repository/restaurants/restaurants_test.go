package restaurants

import (
	"Restobook/configs"
	"Restobook/entities"
	"Restobook/utils"
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestaurantsRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	restaurantRepo := NewRestaurantsRepo(db)

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

	t.Run("Show RestaurantID 1", func(t *testing.T) {
		res, resD, err := restaurantRepo.Get(1)
		assert.Equal(t, res.ID, uint(1))
		assert.Equal(t, resD.ID, uint(1))
		assert.Nil(t, err)
	})

	t.Run("Login RestaurantID 1", func(t *testing.T) {
		hash := sha256.Sum256([]byte("resto123"))
		password := fmt.Sprintf("%x", hash[:])
		var loginUser entities.User
		loginUser.Email = "restaurant1@outlook.my"
		loginUser.Password = password

		res, err := restaurantRepo.LoginRestaurant(loginUser.Email, loginUser.Password)
		assert.Equal(t, res.ID, uint(1))
		assert.Nil(t, err)
	})

	t.Run("Update RestaurantID 1", func(t *testing.T) {
		hash := sha256.Sum256([]byte("resto123"))
		password := fmt.Sprintf("%x", hash[:])
		var updateRestaurant entities.Restaurant
		updateRestaurant.Email = "restaurant1@outlook.my"
		updateRestaurant.Password = password

		res, err := restaurantRepo.Update(uint(1), updateRestaurant)
		assert.Equal(t, res.ID, uint(1))
		assert.Nil(t, err)
	})

	t.Run("Update Detail RestaurantID 1", func(t *testing.T) {
		var updateRestaurant entities.RestaurantDetail
		updateRestaurant.Name = "Restaurant Nasi Padang"
		updateRestaurant.Open = "Monday"
		updateRestaurant.Close = "Friday"
		updateRestaurant.OperationalHour = "10:00 - 17:00"
		updateRestaurant.Price = 10000
		updateRestaurant.Latitude = 1
		updateRestaurant.Longitude = 1
		updateRestaurant.City = "Jakarta"
		updateRestaurant.Address = "Jl.Taman Daan Mogot 2,no.5"
		updateRestaurant.PhoneNumber = "0877"
		updateRestaurant.ProfilePicture = "https://"
		updateRestaurant.Seats = 200
		updateRestaurant.Description = "Khas Rempah Sumbar"
		res, err := restaurantRepo.UpdateDetail(uint(1), updateRestaurant)
		assert.Equal(t, res.ID, uint(1))
		assert.Nil(t, err)
	})

	t.Run("All Waiting RestaurantID 1", func(t *testing.T) {
		res, err := restaurantRepo.GetsWaiting()
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

	t.Run("Approve RestaurantID 1", func(t *testing.T) {
		res, err := restaurantRepo.Approve(uint(1), "OPEN")
		assert.Equal(t, res.ID, uint(1))
		assert.Nil(t, err)
	})

	t.Run("Show Open by Day", func(t *testing.T) {
		res, err := restaurantRepo.GetsByOpen("Monday", "10:00")
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

	t.Run("Show Open", func(t *testing.T) {
		res, err := restaurantRepo.Gets()
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

	t.Run("Delete RestaurantID 1", func(t *testing.T) {
		res, err := restaurantRepo.Delete(1)
		assert.Equal(t, res.ID, uint(1))
		assert.Nil(t, err)
	})
}

func TestFalseRestaurantsRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	restaurantRepo := NewRestaurantsRepo(db)

	t.Run("Register Restaurant 2", func(t *testing.T) {
		hash := sha256.Sum256([]byte("resto123"))
		password := fmt.Sprintf("%x", hash[:])
		var newRestaurant entities.Restaurant
		newRestaurant.Email = "restaurant2@outlook.my"
		newRestaurant.Password = password

		res, err := restaurantRepo.Register(newRestaurant)
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

	t.Run("FALSE Register Restaurant 2", func(t *testing.T) {
		hash := sha256.Sum256([]byte("resto123"))
		password := fmt.Sprintf("%x", hash[:])
		var newRestaurant entities.Restaurant
		newRestaurant.Email = "restaurant2@outlook.my"
		newRestaurant.Password = password

		res, err := restaurantRepo.Register(newRestaurant)
		assert.Equal(t, res, res)
		assert.Error(t, err)
	})

	t.Run("Register Restaurant 3", func(t *testing.T) {
		hash := sha256.Sum256([]byte("resto123"))
		password := fmt.Sprintf("%x", hash[:])
		var newRestaurant entities.Restaurant
		newRestaurant.Email = "restaurant3@outlook.my"
		newRestaurant.Password = password

		res, err := restaurantRepo.Register(newRestaurant)
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

	t.Run("FALSE Show RestaurantID 4", func(t *testing.T) {
		res, resD, err := restaurantRepo.Get(4)
		assert.Equal(t, res.ID, uint(0))
		assert.Equal(t, resD.ID, uint(0))
		assert.Error(t, err)
	})

	t.Run("FALSE Login Restaurant 4", func(t *testing.T) {
		hash := sha256.Sum256([]byte("herlianto1234"))
		password := fmt.Sprintf("%x", hash[:])
		var loginUser entities.User
		loginUser.Email = "restaurant4@outlook.my"
		loginUser.Password = password

		res, err := restaurantRepo.LoginRestaurant(loginUser.Email, loginUser.Password)
		assert.Equal(t, res.ID, uint(0))
		assert.Error(t, err)
	})

	t.Run("FALSE Update RestaurantD 3", func(t *testing.T) {
		hash := sha256.Sum256([]byte("herlianto123"))
		password := fmt.Sprintf("%x", hash[:])
		var updateRestaurant entities.Restaurant
		updateRestaurant.Email = "herlianto@outlook.my"
		updateRestaurant.Password = password

		res, err := restaurantRepo.Update(uint(2), updateRestaurant)
		assert.Equal(t, res.ID, uint(0))
		assert.Error(t, err)
	})

	t.Run("Update Detail RestaurantID 3", func(t *testing.T) {
		var updateRestaurant entities.RestaurantDetail
		updateRestaurant.Name = "Restaurant Nasi Padang"
		updateRestaurant.Open = "Monday"
		updateRestaurant.Close = "Friday"
		updateRestaurant.OperationalHour = "10:00 - 17:00"
		updateRestaurant.Price = 10000
		updateRestaurant.Latitude = 1
		updateRestaurant.Longitude = 1
		updateRestaurant.City = "Jakarta"
		updateRestaurant.Address = "Jl.Taman Daan Mogot 2,no.5"
		updateRestaurant.PhoneNumber = "0877"
		updateRestaurant.ProfilePicture = "https://"
		updateRestaurant.Seats = 200
		updateRestaurant.Description = "Khas Rempah Sumbar"
		res, err := restaurantRepo.UpdateDetail(uint(3), updateRestaurant)
		assert.Equal(t, res.ID, uint(3))
		assert.Nil(t, err)
	})

	t.Run("FALSE Update Detail RestaurantID 4", func(t *testing.T) {
		var updateRestaurant entities.RestaurantDetail
		updateRestaurant.Name = "Restaurant Nasi Padang"

		res, err := restaurantRepo.UpdateDetail(uint(4), updateRestaurant)
		assert.Equal(t, res.ID, uint(0))
		assert.Error(t, err)
	})

	t.Run("FALSE Approve RestaurantID 4", func(t *testing.T) {
		res, err := restaurantRepo.Approve(uint(4), "OPEN")
		assert.Equal(t, res.ID, uint(0))
		assert.Error(t, err)
	})

	t.Run("FALSE Delete RestaurantID 4", func(t *testing.T) {
		res, err := restaurantRepo.Delete(4)
		assert.Equal(t, res.ID, uint(0))
		assert.Error(t, err)
	})

	db.Migrator().DropTable(&entities.RestaurantDetail{})
	db.Migrator().DropTable(&entities.Restaurant{})

	t.Run("FALSE Show Waiting", func(t *testing.T) {
		res, err := restaurantRepo.GetsWaiting()
		assert.Equal(t, res, res)
		assert.Error(t, err)
	})

	t.Run("FALSE Show open by day", func(t *testing.T) {
		res, err := restaurantRepo.GetsByOpen("as", "as")
		assert.Equal(t, res, res)
		assert.Error(t, err)
	})

	t.Run("FALSE Show all open", func(t *testing.T) {
		res, err := restaurantRepo.Gets()
		assert.Equal(t, res, res)
		assert.Error(t, err)
	})

	db.AutoMigrate(entities.RestaurantDetail{})
	db.AutoMigrate(entities.Restaurant{})
}
