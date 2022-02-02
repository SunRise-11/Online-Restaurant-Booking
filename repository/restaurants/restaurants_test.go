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

	t.Run("Register Restaurant", func(t *testing.T) {
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
		res, err := restaurantRepo.Get(1)
		assert.Equal(t, res.ID, uint(1))
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

	t.Run("Register Restaurant", func(t *testing.T) {
		hash := sha256.Sum256([]byte("resto123"))
		password := fmt.Sprintf("%x", hash[:])
		var newRestaurant entities.Restaurant
		newRestaurant.Email = "restaurant1@outlook.my"
		newRestaurant.Password = password

		res, err := restaurantRepo.Register(newRestaurant)
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

	t.Run("FALSE Register Restaurant", func(t *testing.T) {
		hash := sha256.Sum256([]byte("resto123"))
		password := fmt.Sprintf("%x", hash[:])
		var newRestaurant entities.Restaurant
		newRestaurant.Email = "restaurant1@outlook.my"
		newRestaurant.Password = password

		res, err := restaurantRepo.Register(newRestaurant)
		assert.Equal(t, res, res)
		assert.Error(t, err)
	})

	t.Run("FALSE Show RestaurantID 2", func(t *testing.T) {
		res, err := restaurantRepo.Get(2)
		assert.Equal(t, res.ID, uint(0))
		assert.Error(t, err)
	})

	t.Run("FALSE Login Restaurant", func(t *testing.T) {
		hash := sha256.Sum256([]byte("herlianto1234"))
		password := fmt.Sprintf("%x", hash[:])
		var loginUser entities.User
		loginUser.Email = "herlianto@outlook.my"
		loginUser.Password = password

		res, err := restaurantRepo.LoginRestaurant(loginUser.Email, loginUser.Password)
		assert.Equal(t, res.ID, uint(0))
		assert.Error(t, err)
	})

	t.Run("FALSE Update RestaurantID 2", func(t *testing.T) {
		hash := sha256.Sum256([]byte("herlianto123"))
		password := fmt.Sprintf("%x", hash[:])
		var updateRestaurant entities.Restaurant
		updateRestaurant.Email = "herlianto@outlook.my"
		updateRestaurant.Password = password

		res, err := restaurantRepo.Update(uint(2), updateRestaurant)
		assert.Equal(t, res.ID, uint(0))
		assert.Error(t, err)
	})

	t.Run("FALSE Delete RestaurantID 2", func(t *testing.T) {
		res, err := restaurantRepo.Delete(2)
		assert.Equal(t, res.ID, uint(0))
		assert.Error(t, err)
	})

}
