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

func TestRegisterRestaurantsRepo(t *testing.T) {
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
	t.Run("ERROR Register Restaurant", func(t *testing.T) {
		hash := sha256.Sum256([]byte("resto123"))
		password := fmt.Sprintf("%x", hash[:])
		var newRestaurant entities.Restaurant
		newRestaurant.Email = "restaurant1@outlook.my"
		newRestaurant.Password = password

		res, err := restaurantRepo.Register(newRestaurant)
		assert.Equal(t, res, res)
		assert.Error(t, err)
	})

}

func TestLoginRestaurantsRepo(t *testing.T) {
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

	t.Run("Login Restaurant", func(t *testing.T) {
		hash := sha256.Sum256([]byte("resto123"))
		password := fmt.Sprintf("%x", hash[:])
		var newRestaurant entities.Restaurant
		newRestaurant.Email = "restaurant1@outlook.my"
		newRestaurant.Password = password

		res, err := restaurantRepo.Login(newRestaurant.Email, newRestaurant.Password)
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

	t.Run("ERROR Login Restaurant", func(t *testing.T) {
		hash := sha256.Sum256([]byte("resto1234"))
		password := fmt.Sprintf("%x", hash[:])
		var newRestaurant entities.Restaurant
		newRestaurant.Email = "restaurant1@outlook.my"
		newRestaurant.Password = password

		res, err := restaurantRepo.Login(newRestaurant.Email, newRestaurant.Password)
		assert.Equal(t, res, res)
		assert.Error(t, err)
	})

}

// func TestGetWaitingRestaurantsRepo(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)

// 	restaurantRepo := NewRestaurantsRepo(db)

// 	t.Run("Register Restaurant", func(t *testing.T) {
// 		hash := sha256.Sum256([]byte("resto123"))
// 		password := fmt.Sprintf("%x", hash[:])
// 		var newRestaurant entities.Restaurant
// 		newRestaurant.Email = "restaurant1@outlook.my"
// 		newRestaurant.Password = password

// 		res, err := restaurantRepo.Register(newRestaurant)
// 		assert.Equal(t, res, res)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Update Restaurant Detail", func(t *testing.T) {
// 		var updateRestaurant entities.RestaurantDetail
// 		updateRestaurant.Name = "Restaurant Nasi Padang"
// 		updateRestaurant.Open = "Monday,Tuesday"
// 		updateRestaurant.Close = "Wednesday,Thursday,Friday,Saturday,Sunday"
// 		updateRestaurant.Open_Hour = "10:00"
// 		updateRestaurant.Close_Hour = "17:00"
// 		updateRestaurant.Price = 10000
// 		updateRestaurant.Latitude = 1
// 		updateRestaurant.Longitude = 1
// 		updateRestaurant.City = "Jakarta"
// 		updateRestaurant.Address = "Jl.Taman Daan Mogot 2,no.5"
// 		updateRestaurant.PhoneNumber = "0877"
// 		updateRestaurant.ProfilePicture = "https://"
// 		updateRestaurant.Seats = 200
// 		updateRestaurant.Description = "Khas Rempah Sumbar"
// 		res, err := restaurantRepo.UpdateDetail(uint(1), updateRestaurant)
// 		assert.Equal(t, res.ID, uint(1))
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Get Waiting Restaurant", func(t *testing.T) {
// 		res, err := restaurantRepo.GetsWaiting()
// 		assert.Equal(t, res, res)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Approve Restaurant", func(t *testing.T) {
// 		res, err := restaurantRepo.Approve(1, "OPEN")
// 		assert.Equal(t, res, res)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Get Waiting Restaurant", func(t *testing.T) {
// 		res, err := restaurantRepo.GetsWaiting()
// 		fmt.Println("res", res)
// 		fmt.Println("err", err)
// 		assert.Equal(t, res, res)
// 		assert.Error(t, err)
// 	})
// }

// func TestApproveRestaurantsRepo(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)

// 	restaurantRepo := NewRestaurantsRepo(db)

// 	t.Run("Register Restaurant", func(t *testing.T) {
// 		hash := sha256.Sum256([]byte("resto123"))
// 		password := fmt.Sprintf("%x", hash[:])
// 		var newRestaurant entities.Restaurant
// 		newRestaurant.Email = "restaurant1@outlook.my"
// 		newRestaurant.Password = password

// 		res, err := restaurantRepo.Register(newRestaurant)
// 		assert.Equal(t, res, res)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Approve Restaurant", func(t *testing.T) {
// 		res, err := restaurantRepo.Approve(uint(1), "OPEN")
// 		assert.Equal(t, res.ID, uint(1))
// 		assert.Nil(t, err)
// 	})

// 	t.Run("ERROR Approve Restaurant", func(t *testing.T) {
// 		res, err := restaurantRepo.Approve(2, "OPEN")
// 		assert.Equal(t, res, res)
// 		assert.Error(t, err)
// 	})

// }

// func TestGetRestaurantsRepo(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)

// 	restaurantRepo := NewRestaurantsRepo(db)

// 	t.Run("Register Restaurant", func(t *testing.T) {
// 		hash := sha256.Sum256([]byte("resto123"))
// 		password := fmt.Sprintf("%x", hash[:])
// 		var newRestaurant entities.Restaurant
// 		newRestaurant.Email = "restaurant1@outlook.my"
// 		newRestaurant.Password = password

// 		res, err := restaurantRepo.Register(newRestaurant)
// 		assert.Equal(t, res, res)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Get Restaurant", func(t *testing.T) {
// 		res, resD, err := restaurantRepo.Get(uint(1))
// 		assert.Equal(t, res.ID, uint(1))
// 		assert.Equal(t, resD, resD)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("ERROR Get Restaurant", func(t *testing.T) {
// 		res, resD, err := restaurantRepo.Get(uint(2))
// 		assert.Equal(t, res.ID, uint(0))
// 		assert.Equal(t, resD, resD)
// 		assert.Error(t, err)
// 	})

// }

// func TestGetsByOpenRestaurantsRepo(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)

// 	restaurantRepo := NewRestaurantsRepo(db)

// 	t.Run("Register Restaurant", func(t *testing.T) {
// 		hash := sha256.Sum256([]byte("resto123"))
// 		password := fmt.Sprintf("%x", hash[:])
// 		var newRestaurant entities.Restaurant
// 		newRestaurant.Email = "restaurant1@outlook.my"
// 		newRestaurant.Password = password

// 		res, err := restaurantRepo.Register(newRestaurant)
// 		assert.Equal(t, res, res)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Update Restaurant Detail", func(t *testing.T) {
// 		var updateRestaurant entities.RestaurantDetail
// 		updateRestaurant.Name = "Restaurant Nasi Padang"
// 		updateRestaurant.Open = "Monday,Tuesday"
// 		updateRestaurant.Close = "Wednesday,Thursday,Friday,Saturday,Sunday"
// 		updateRestaurant.Open_Hour = "10:00"
// 		updateRestaurant.Close_Hour = "17:00"
// 		updateRestaurant.Price = 10000
// 		updateRestaurant.Latitude = 1
// 		updateRestaurant.Longitude = 1
// 		updateRestaurant.City = "Jakarta"
// 		updateRestaurant.Address = "Jl.Taman Daan Mogot 2,no.5"
// 		updateRestaurant.PhoneNumber = "0877"
// 		updateRestaurant.ProfilePicture = "https://"
// 		updateRestaurant.Seats = 200
// 		updateRestaurant.Description = "Khas Rempah Sumbar"
// 		res, err := restaurantRepo.UpdateDetail(uint(1), updateRestaurant)
// 		assert.Equal(t, res.ID, uint(1))
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Approve Restaurant", func(t *testing.T) {
// 		res, err := restaurantRepo.Approve(1, "OPEN")
// 		assert.Equal(t, res, res)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("GetsByOpen Restaurant", func(t *testing.T) {
// 		res, err := restaurantRepo.GetsByOpen(1)
// 		assert.Equal(t, res, res)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("ERROR GetsByOpen Restaurant", func(t *testing.T) {
// 		res, err := restaurantRepo.GetsByOpen(4)
// 		assert.Equal(t, res, res)
// 		assert.Error(t, err)
// 	})

// }

// func TestGetsRestaurantsRepo(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)

// 	restaurantRepo := NewRestaurantsRepo(db)

// 	t.Run("ERROR Gets Restaurant", func(t *testing.T) {
// 		res, err := restaurantRepo.Gets()
// 		assert.Equal(t, res, res)
// 		assert.Error(t, err)
// 	})

// 	t.Run("Register Restaurant", func(t *testing.T) {
// 		hash := sha256.Sum256([]byte("resto123"))
// 		password := fmt.Sprintf("%x", hash[:])
// 		var newRestaurant entities.Restaurant
// 		newRestaurant.Email = "restaurant1@outlook.my"
// 		newRestaurant.Password = password

// 		res, err := restaurantRepo.Register(newRestaurant)
// 		assert.Equal(t, res, res)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Register Restaurant", func(t *testing.T) {
// 		hash := sha256.Sum256([]byte("resto123"))
// 		password := fmt.Sprintf("%x", hash[:])
// 		var newRestaurant entities.Restaurant
// 		newRestaurant.Email = "restaurant2@outlook.my"
// 		newRestaurant.Password = password

// 		res, err := restaurantRepo.Register(newRestaurant)
// 		assert.Equal(t, res, res)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Update Restaurant Detail", func(t *testing.T) {
// 		var updateRestaurant entities.RestaurantDetail
// 		updateRestaurant.Name = "Restaurant Nasi Padang"
// 		updateRestaurant.Open = "Monday,Tuesday"
// 		updateRestaurant.Close = "Wednesday,Thursday,Friday,Saturday,Sunday"
// 		updateRestaurant.Open_Hour = "10:00"
// 		updateRestaurant.Close_Hour = "17:00"
// 		updateRestaurant.Price = 10000
// 		updateRestaurant.Latitude = 1
// 		updateRestaurant.Longitude = 1
// 		updateRestaurant.City = "Jakarta"
// 		updateRestaurant.Address = "Jl.Taman Daan Mogot 2,no.5"
// 		updateRestaurant.PhoneNumber = "0877"
// 		updateRestaurant.ProfilePicture = "https://"
// 		updateRestaurant.Seats = 200
// 		updateRestaurant.Description = "Khas Rempah Sumbar"
// 		res, err := restaurantRepo.UpdateDetail(uint(1), updateRestaurant)
// 		assert.Equal(t, res.ID, uint(1))
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Approve Restaurant", func(t *testing.T) {
// 		res, err := restaurantRepo.Approve(1, "OPEN")
// 		assert.Equal(t, res, res)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Gets Restaurant", func(t *testing.T) {
// 		res, err := restaurantRepo.Gets()
// 		assert.Equal(t, res, res)
// 		assert.Nil(t, err)
// 	})

// }

// func TestGetExistSeatRestaurantsRepo(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)

// 	restaurantRepo := NewRestaurantsRepo(db)
// 	userRepo := users.NewUsersRepo(db)
// 	transactionRepo := transactions.NewTransactionRepo(db)

// 	t.Run("Register Restaurant", func(t *testing.T) {
// 		hash := sha256.Sum256([]byte("resto123"))
// 		password := fmt.Sprintf("%x", hash[:])
// 		var newRestaurant entities.Restaurant
// 		newRestaurant.Email = "restaurant1@outlook.my"
// 		newRestaurant.Password = password

// 		res, err := restaurantRepo.Register(newRestaurant)
// 		assert.Equal(t, res, res)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Register Restaurant", func(t *testing.T) {
// 		hash := sha256.Sum256([]byte("resto123"))
// 		password := fmt.Sprintf("%x", hash[:])
// 		var newRestaurant entities.Restaurant
// 		newRestaurant.Email = "restaurant2@outlook.my"
// 		newRestaurant.Password = password

// 		res, err := restaurantRepo.Register(newRestaurant)
// 		assert.Equal(t, res, res)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Update Restaurant Detail", func(t *testing.T) {
// 		var updateRestaurant entities.RestaurantDetail
// 		updateRestaurant.Name = "Restaurant Nasi Padang"
// 		updateRestaurant.Open = "Monday,Tuesday"
// 		updateRestaurant.Close = "Wednesday,Thursday,Friday,Saturday,Sunday"
// 		updateRestaurant.Open_Hour = "10:00"
// 		updateRestaurant.Close_Hour = "21:00"
// 		updateRestaurant.Price = 10000
// 		updateRestaurant.Latitude = 1
// 		updateRestaurant.Longitude = 1
// 		updateRestaurant.City = "Jakarta"
// 		updateRestaurant.Address = "Jl.Taman Daan Mogot 2,no.5"
// 		updateRestaurant.PhoneNumber = "0877"
// 		updateRestaurant.ProfilePicture = "https://"
// 		updateRestaurant.Seats = 200
// 		updateRestaurant.Description = "Khas Rempah Sumbar"
// 		res, err := restaurantRepo.UpdateDetail(uint(1), updateRestaurant)
// 		assert.Equal(t, res.ID, uint(1))
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Update Restaurant Detail", func(t *testing.T) {
// 		var updateRestaurant entities.RestaurantDetail
// 		updateRestaurant.Name = "Restaurant Nasi Padang"
// 		updateRestaurant.Open = "Monday"
// 		updateRestaurant.Close = "Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday"
// 		updateRestaurant.Open_Hour = "21:00"
// 		updateRestaurant.Close_Hour = "22:00"
// 		updateRestaurant.Price = 10000
// 		updateRestaurant.Latitude = 1
// 		updateRestaurant.Longitude = 1
// 		updateRestaurant.City = "Jakarta"
// 		updateRestaurant.Address = "Jl.Taman Daan Mogot 2,no.5"
// 		updateRestaurant.PhoneNumber = "0877"
// 		updateRestaurant.ProfilePicture = "https://"
// 		updateRestaurant.Seats = 200
// 		updateRestaurant.Description = "Khas Rempah Sumbar"
// 		res, err := restaurantRepo.UpdateDetail(uint(2), updateRestaurant)
// 		assert.Equal(t, res.ID, uint(2))
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Approve Restaurant", func(t *testing.T) {
// 		res, err := restaurantRepo.Approve(1, "OPEN")
// 		assert.Equal(t, res, res)
// 		assert.Nil(t, err)
// 	})
// 	t.Run("Approve Restaurant", func(t *testing.T) {
// 		res, err := restaurantRepo.Approve(2, "OPEN")
// 		assert.Equal(t, res, res)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Register User", func(t *testing.T) {
// 		hash := sha256.Sum256([]byte("herli123"))
// 		password := fmt.Sprintf("%x", hash[:])
// 		var newUser entities.User
// 		newUser.Email = "herlianto@outlook.my"
// 		newUser.Password = password

// 		res, err := userRepo.Register(newUser)
// 		assert.Equal(t, res, res)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Create Transactions", func(t *testing.T) {
// 		loc, _ := time.LoadLocation("Asia/Singapore")
// 		date_string := "2022-02-14 16:00"
// 		var dateTime, _ = time.ParseInLocation("2006-01-02 15:04", date_string, loc)
// 		fmt.Println("date_time", dateTime)

// 		var newTransaction entities.Transaction
// 		newTransaction.RestaurantID = 1
// 		newTransaction.UserID = 1
// 		newTransaction.DateTime = dateTime
// 		newTransaction.Persons = 1
// 		newTransaction.Total = 10000
// 		res, err := transactionRepo.Create(newTransaction)
// 		assert.Equal(t, res, res)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Create Transactions", func(t *testing.T) {
// 		loc, _ := time.LoadLocation("Asia/Singapore")
// 		date_string := "2022-02-14 16:00"
// 		var dateTime, _ = time.ParseInLocation("2006-01-02 15:04", date_string, loc)
// 		fmt.Println("date_time", dateTime)

// 		var newTransaction entities.Transaction
// 		newTransaction.RestaurantID = 2
// 		newTransaction.UserID = 1
// 		newTransaction.DateTime = dateTime
// 		newTransaction.Persons = 0
// 		newTransaction.Total = 10000
// 		res, err := transactionRepo.Create(newTransaction)
// 		assert.Equal(t, res, res)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("GetExistSeat Restaurant", func(t *testing.T) {
// 		res, total_seat, err := restaurantRepo.GetExistSeat(1, "2022-02-14 16:00")
// 		assert.Equal(t, res, res)
// 		assert.Equal(t, total_seat, 1)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("ERROR GetExistSeat Restaurant", func(t *testing.T) {
// 		res, total_seat, err := restaurantRepo.GetExistSeat(0, "2022-02-14 16:00")
// 		fmt.Println("res", res)
// 		fmt.Println("total_seat", total_seat)
// 		fmt.Println("err", err)
// 		assert.Equal(t, res, res)
// 		assert.Equal(t, total_seat, 0)
// 		assert.Error(t, err)
// 	})
// 	fmt.Println("=======")
// 	t.Run("=>>>>>ERROR GetExistSeat Restaurant", func(t *testing.T) {
// 		res, total_seat, err := restaurantRepo.GetExistSeat(2, "2022-02-14 10:00")
// 		fmt.Println("res resto2", res)
// 		fmt.Println("total_seat resto2", total_seat)
// 		fmt.Println("err resto2", err)
// 		assert.Equal(t, res, res)
// 		assert.Equal(t, total_seat, 0)
// 		assert.Error(t, err)
// 	})
// 	fmt.Println("===========")
// }

// func TestUpdateRestaurantsRepo(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)

// 	restaurantRepo := NewRestaurantsRepo(db)

// 	t.Run("Register Restaurant", func(t *testing.T) {
// 		hash := sha256.Sum256([]byte("resto123"))
// 		password := fmt.Sprintf("%x", hash[:])
// 		var newRestaurant entities.Restaurant
// 		newRestaurant.Email = "restaurant1@outlook.my"
// 		newRestaurant.Password = password

// 		res, err := restaurantRepo.Register(newRestaurant)
// 		assert.Equal(t, res, res)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Update Restaurant", func(t *testing.T) {
// 		var updateRestaurant entities.Restaurant
// 		updateRestaurant.Email = "restaurant1@outlook.my"
// 		updateRestaurant.Password = "resto123"
// 		res, err := restaurantRepo.Update(uint(1), updateRestaurant)
// 		assert.Equal(t, res.ID, uint(1))
// 		assert.Nil(t, err)
// 	})
// 	t.Run("ERROR Update Restaurant", func(t *testing.T) {
// 		var updateRestaurant entities.Restaurant
// 		updateRestaurant.Email = "restaurant1@outlook.my"
// 		updateRestaurant.Password = "resto123"
// 		res, err := restaurantRepo.Update(uint(2), updateRestaurant)
// 		assert.Equal(t, res.ID, uint(0))
// 		assert.Error(t, err)
// 	})

// }

// func TestDeleteRestaurantsRepo(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)

// 	restaurantRepo := NewRestaurantsRepo(db)

// 	t.Run("Register Restaurant", func(t *testing.T) {
// 		hash := sha256.Sum256([]byte("resto123"))
// 		password := fmt.Sprintf("%x", hash[:])
// 		var newRestaurant entities.Restaurant
// 		newRestaurant.Email = "restaurant1@outlook.my"
// 		newRestaurant.Password = password

// 		res, err := restaurantRepo.Register(newRestaurant)
// 		assert.Equal(t, res, res)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Delete Restaurant", func(t *testing.T) {
// 		res, err := restaurantRepo.Delete(uint(1))
// 		assert.Equal(t, res.ID, uint(1))
// 		assert.Nil(t, err)
// 	})

// 	t.Run("ERROR Delete Restaurant", func(t *testing.T) {
// 		res, err := restaurantRepo.Delete(uint(2))
// 		assert.Equal(t, res.ID, uint(0))
// 		assert.Error(t, err)
// 	})

// }

// func TestUpdateDetailRestaurantsRepo(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)

// 	restaurantRepo := NewRestaurantsRepo(db)

// 	t.Run("Register Restaurant", func(t *testing.T) {
// 		hash := sha256.Sum256([]byte("resto123"))
// 		password := fmt.Sprintf("%x", hash[:])
// 		var newRestaurant entities.Restaurant
// 		newRestaurant.Email = "restaurant1@outlook.my"
// 		newRestaurant.Password = password

// 		res, err := restaurantRepo.Register(newRestaurant)
// 		assert.Equal(t, res, res)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Update Restaurant Detail", func(t *testing.T) {
// 		var updateRestaurant entities.RestaurantDetail
// 		updateRestaurant.Name = "Restaurant Nasi Padang"
// 		updateRestaurant.Open = "Monday,Tuesday"
// 		updateRestaurant.Close = "Wednesday,Thursday,Friday,Saturday,Sunday"
// 		updateRestaurant.Open_Hour = "10:00"
// 		updateRestaurant.Close_Hour = "17:00"
// 		updateRestaurant.Price = 10000
// 		updateRestaurant.Latitude = 1
// 		updateRestaurant.Longitude = 1
// 		updateRestaurant.City = "Jakarta"
// 		updateRestaurant.Address = "Jl.Taman Daan Mogot 2,no.5"
// 		updateRestaurant.PhoneNumber = "0877"
// 		updateRestaurant.ProfilePicture = "https://"
// 		updateRestaurant.Seats = 200
// 		updateRestaurant.Description = "Khas Rempah Sumbar"
// 		res, err := restaurantRepo.UpdateDetail(uint(1), updateRestaurant)
// 		assert.Equal(t, res.ID, uint(1))
// 		assert.Nil(t, err)
// 	})

// 	t.Run("ERROR Update Restaurant Detail", func(t *testing.T) {
// 		var updateRestaurant entities.RestaurantDetail
// 		updateRestaurant.Name = "Restaurant Nasi Padang"
// 		updateRestaurant.Open = "Monday,Tuesday"
// 		updateRestaurant.Close = "Wednesday,Thursday,Friday,Saturday,Sunday"
// 		updateRestaurant.Open_Hour = "10:00"
// 		updateRestaurant.Close_Hour = "17:00"
// 		updateRestaurant.Price = 10000
// 		updateRestaurant.Latitude = 1
// 		updateRestaurant.Longitude = 1
// 		updateRestaurant.City = "Jakarta"
// 		updateRestaurant.Address = "Jl.Taman Daan Mogot 2,no.5"
// 		updateRestaurant.PhoneNumber = "0877"
// 		updateRestaurant.ProfilePicture = "https://"
// 		updateRestaurant.Seats = 200
// 		updateRestaurant.Description = "Khas Rempah Sumbar"
// 		res, err := restaurantRepo.UpdateDetail(uint(2), updateRestaurant)
// 		assert.Equal(t, res.ID, uint(0))
// 		assert.Error(t, err)
// 	})

// }
