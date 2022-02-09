package transactions

import (
	"Restobook/configs"
	"Restobook/entities"
	"Restobook/repository/restaurants"
	"Restobook/repository/users"
	"Restobook/utils"
	"crypto/sha256"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Create_Transaction(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	loc, _ := time.LoadLocation("Asia/Jakarta")

	userRepo := users.NewUsersRepo(db)
	restaurantRepo := restaurants.NewRestaurantsRepo(db)
	transactionRepo := NewTransactionRepo(db)

	var newUser entities.User
	hash := sha256.Sum256([]byte("user123"))
	password := fmt.Sprintf("%x", hash[:])
	newUser.Email = "user1@outlook.my"
	newUser.Password = password
	userRepo.Register(newUser)

	var newRestaurant entities.Restaurant
	hashResto := sha256.Sum256([]byte("resto123"))
	passwordResto := fmt.Sprintf("%x", hashResto[:])
	newRestaurant.Email = "resto1@outlook.my"
	newRestaurant.Password = passwordResto
	restaurantRepo.Register(newRestaurant)
	t.Run("Create Transaction", func(t *testing.T) {
		var dateTime, _ = time.ParseInLocation("2006-01-02 15:04", "2022-02-09 10:00", loc)
		var newTransaction entities.Transaction
		newTransaction.UserID = 1
		newTransaction.RestaurantID = 1
		newTransaction.DateTime = dateTime
		_, err := transactionRepo.Create(newTransaction)
		assert.Nil(t, err)

	})
	t.Run("Error Create Transaction", func(t *testing.T) {
		var dateTime, _ = time.ParseInLocation("2006-01-02 15:04", "2022-02-09 10:00", loc)
		var newTransaction entities.Transaction
		newTransaction.RestaurantID = 1
		newTransaction.DateTime = dateTime
		_, err := transactionRepo.Create(newTransaction)
		assert.Error(t, err)

	})
}
func Test_Get_All_Waiting_Status_For_User(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	loc, _ := time.LoadLocation("Asia/Jakarta")

	userRepo := users.NewUsersRepo(db)
	restaurantRepo := restaurants.NewRestaurantsRepo(db)
	transactionRepo := NewTransactionRepo(db)

	var newUser entities.User
	hash := sha256.Sum256([]byte("user123"))
	password := fmt.Sprintf("%x", hash[:])
	newUser.Email = "user1@outlook.my"
	newUser.Password = password
	userRepo.Register(newUser)

	var newRestaurant entities.Restaurant
	hashResto := sha256.Sum256([]byte("resto123"))
	passwordResto := fmt.Sprintf("%x", hashResto[:])
	newRestaurant.Email = "resto1@outlook.my"
	newRestaurant.Password = passwordResto
	restaurantRepo.Register(newRestaurant)
	t.Run("Create Transaction", func(t *testing.T) {
		var dateTime, _ = time.ParseInLocation("2006-01-02 15:04", "2022-02-09 10:00", loc)
		var newTransaction entities.Transaction
		newTransaction.UserID = 1
		newTransaction.RestaurantID = 1
		newTransaction.DateTime = dateTime
		_, err := transactionRepo.Create(newTransaction)
		assert.Nil(t, err)

	})
	t.Run("Get All Waiting Status For User", func(t *testing.T) {
		_, err := transactionRepo.GetAllWaiting(1)
		assert.Nil(t, err)

	})
	t.Run("Get All Waiting Status For User", func(t *testing.T) {
		_, err := transactionRepo.GetAllWaiting(100)
		assert.Error(t, err)

	})

}

func Test_Get_All_Waiting_Status_For_Resto(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	loc, _ := time.LoadLocation("Asia/Jakarta")

	userRepo := users.NewUsersRepo(db)
	restaurantRepo := restaurants.NewRestaurantsRepo(db)
	transactionRepo := NewTransactionRepo(db)

	var newUser entities.User
	hash := sha256.Sum256([]byte("user123"))
	password := fmt.Sprintf("%x", hash[:])
	newUser.Email = "user1@outlook.my"
	newUser.Password = password
	userRepo.Register(newUser)

	var newRestaurant entities.Restaurant
	hashResto := sha256.Sum256([]byte("resto123"))
	passwordResto := fmt.Sprintf("%x", hashResto[:])
	newRestaurant.Email = "resto1@outlook.my"
	newRestaurant.Password = passwordResto
	restaurantRepo.Register(newRestaurant)
	t.Run("Create Transaction", func(t *testing.T) {
		var dateTime, _ = time.ParseInLocation("2006-01-02 15:04", "2022-02-09 10:00", loc)
		var newTransaction entities.Transaction
		newTransaction.UserID = 1
		newTransaction.RestaurantID = 1
		newTransaction.DateTime = dateTime
		_, err := transactionRepo.Create(newTransaction)
		assert.Nil(t, err)

	})
	t.Run("Get All Waiting Status For Resto", func(t *testing.T) {
		_, err := transactionRepo.GetAllWaitingForResto(1)
		assert.Nil(t, err)

	})
	t.Run("Get All Waiting Status For Resto", func(t *testing.T) {
		_, err := transactionRepo.GetAllWaitingForResto(100)
		assert.Error(t, err)

	})

}
func Test_Get_All_Accepted_Status_For_Resto(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	loc, _ := time.LoadLocation("Asia/Jakarta")

	userRepo := users.NewUsersRepo(db)
	restaurantRepo := restaurants.NewRestaurantsRepo(db)
	transactionRepo := NewTransactionRepo(db)

	var newUser entities.User
	hash := sha256.Sum256([]byte("user123"))
	password := fmt.Sprintf("%x", hash[:])
	newUser.Email = "user1@outlook.my"
	newUser.Password = password
	userRepo.Register(newUser)

	var newRestaurant entities.Restaurant
	hashResto := sha256.Sum256([]byte("resto123"))
	passwordResto := fmt.Sprintf("%x", hashResto[:])
	newRestaurant.Email = "resto1@outlook.my"
	newRestaurant.Password = passwordResto
	restaurantRepo.Register(newRestaurant)
	t.Run("Create Transaction", func(t *testing.T) {
		var dateTime, _ = time.ParseInLocation("2006-01-02 15:04", "2022-02-09 10:00", loc)
		var newTransaction entities.Transaction
		newTransaction.UserID = 1
		newTransaction.RestaurantID = 1
		newTransaction.DateTime = dateTime
		newTransaction.Status = "Accepted"
		_, err := transactionRepo.Create(newTransaction)
		assert.Nil(t, err)

	})
	t.Run("Get All Waiting Status For Resto", func(t *testing.T) {
		_, err := transactionRepo.GetAllAcceptedForResto(1)
		assert.Nil(t, err)

	})
	t.Run("Get All Waiting Status For Resto", func(t *testing.T) {
		_, err := transactionRepo.GetAllAcceptedForResto(100)
		assert.Error(t, err)

	})

}
func Test_Get_All_History_Transaction_For_User(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	loc, _ := time.LoadLocation("Asia/Jakarta")

	userRepo := users.NewUsersRepo(db)
	restaurantRepo := restaurants.NewRestaurantsRepo(db)
	transactionRepo := NewTransactionRepo(db)

	var newUser entities.User
	hash := sha256.Sum256([]byte("user123"))
	password := fmt.Sprintf("%x", hash[:])
	newUser.Email = "user1@outlook.my"
	newUser.Password = password
	userRepo.Register(newUser)

	var newRestaurant entities.Restaurant
	hashResto := sha256.Sum256([]byte("resto123"))
	passwordResto := fmt.Sprintf("%x", hashResto[:])
	newRestaurant.Email = "resto1@outlook.my"
	newRestaurant.Password = passwordResto
	restaurantRepo.Register(newRestaurant)
	t.Run("Create Transaction", func(t *testing.T) {
		var dateTime, _ = time.ParseInLocation("2006-01-02 15:04", "2022-02-09 10:00", loc)
		var newTransaction entities.Transaction
		newTransaction.UserID = 1
		newTransaction.RestaurantID = 1
		newTransaction.DateTime = dateTime
		newTransaction.Status = "Success"
		_, err := transactionRepo.Create(newTransaction)
		assert.Nil(t, err)

	})
	t.Run("Get All History Transaction", func(t *testing.T) {
		_, err := transactionRepo.GetHistory(1)
		assert.Nil(t, err)

	})
	t.Run("Error Get All History Tranasction", func(t *testing.T) {
		_, err := transactionRepo.GetHistory(100)
		assert.Error(t, err)

	})

}
func Test_Get_All_Accepted_Status_For_User(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	loc, _ := time.LoadLocation("Asia/Jakarta")

	userRepo := users.NewUsersRepo(db)
	restaurantRepo := restaurants.NewRestaurantsRepo(db)
	transactionRepo := NewTransactionRepo(db)

	var newUser entities.User
	hash := sha256.Sum256([]byte("user123"))
	password := fmt.Sprintf("%x", hash[:])
	newUser.Email = "user1@outlook.my"
	newUser.Password = password
	userRepo.Register(newUser)

	var newRestaurant entities.Restaurant
	hashResto := sha256.Sum256([]byte("resto123"))
	passwordResto := fmt.Sprintf("%x", hashResto[:])
	newRestaurant.Email = "resto1@outlook.my"
	newRestaurant.Password = passwordResto
	restaurantRepo.Register(newRestaurant)
	t.Run("Create Transaction", func(t *testing.T) {
		var dateTime, _ = time.ParseInLocation("2006-01-02 15:04", "2022-02-09 10:00", loc)
		var newTransaction entities.Transaction
		newTransaction.UserID = 1
		newTransaction.RestaurantID = 1
		newTransaction.DateTime = dateTime
		newTransaction.Status = "Accepted"
		_, err := transactionRepo.Create(newTransaction)
		assert.Nil(t, err)

	})
	t.Run("Get All Accepted Status For User", func(t *testing.T) {
		_, err := transactionRepo.GetAllAppointed(1)
		assert.Nil(t, err)

	})
	t.Run("Error Get All Accepted Status For User", func(t *testing.T) {
		_, err := transactionRepo.GetAllAppointed(100)
		assert.Error(t, err)

	})

}
func Test_Get_Transaction_By_Id(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	loc, _ := time.LoadLocation("Asia/Jakarta")

	userRepo := users.NewUsersRepo(db)
	restaurantRepo := restaurants.NewRestaurantsRepo(db)
	transactionRepo := NewTransactionRepo(db)

	var newUser entities.User
	hash := sha256.Sum256([]byte("user123"))
	password := fmt.Sprintf("%x", hash[:])
	newUser.Email = "user1@outlook.my"
	newUser.Password = password
	userRepo.Register(newUser)

	var newRestaurant entities.Restaurant
	hashResto := sha256.Sum256([]byte("resto123"))
	passwordResto := fmt.Sprintf("%x", hashResto[:])
	newRestaurant.Email = "resto1@outlook.my"
	newRestaurant.Password = passwordResto
	restaurantRepo.Register(newRestaurant)
	t.Run("Create Transaction", func(t *testing.T) {
		var dateTime, _ = time.ParseInLocation("2006-01-02 15:04", "2022-02-09 10:00", loc)
		var newTransaction entities.Transaction
		newTransaction.UserID = 1
		newTransaction.RestaurantID = 1
		newTransaction.DateTime = dateTime
		_, err := transactionRepo.Create(newTransaction)
		assert.Nil(t, err)

	})
	t.Run("Get Transaction By Id", func(t *testing.T) {
		_, err := transactionRepo.GetTransactionById(1, 1, "waiting for confirmation")
		assert.Nil(t, err)

	})
	t.Run("Error Get Transaction By Id", func(t *testing.T) {
		_, err := transactionRepo.GetTransactionById(100, 100, "Sucess")
		assert.Error(t, err)

	})

}
func Test_Get_Transaction_User_By_Status(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	loc, _ := time.LoadLocation("Asia/Jakarta")

	userRepo := users.NewUsersRepo(db)
	restaurantRepo := restaurants.NewRestaurantsRepo(db)
	transactionRepo := NewTransactionRepo(db)

	var newUser entities.User
	hash := sha256.Sum256([]byte("user123"))
	password := fmt.Sprintf("%x", hash[:])
	newUser.Email = "user1@outlook.my"
	newUser.Password = password
	userRepo.Register(newUser)

	var newRestaurant entities.Restaurant
	hashResto := sha256.Sum256([]byte("resto123"))
	passwordResto := fmt.Sprintf("%x", hashResto[:])
	newRestaurant.Email = "resto1@outlook.my"
	newRestaurant.Password = passwordResto
	restaurantRepo.Register(newRestaurant)
	t.Run("Create Transaction", func(t *testing.T) {
		var dateTime, _ = time.ParseInLocation("2006-01-02 15:04", "2022-02-09 10:00", loc)
		var newTransaction entities.Transaction
		newTransaction.UserID = 1
		newTransaction.RestaurantID = 1
		newTransaction.DateTime = dateTime
		_, err := transactionRepo.Create(newTransaction)
		assert.Nil(t, err)

	})
	t.Run("Get Transaction User By Status", func(t *testing.T) {
		_, err := transactionRepo.GetTransactionUserByStatus(1, 1, "waiting for confirmation")
		assert.Nil(t, err)

	})
	t.Run("Error Get Transaction User By Status", func(t *testing.T) {
		_, err := transactionRepo.GetTransactionUserByStatus(100, 100, "Success")
		assert.Error(t, err)

	})

}
func Test_Get_User_Balance(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := users.NewUsersRepo(db)
	transactionRepo := NewTransactionRepo(db)
	var newUser entities.User
	hash := sha256.Sum256([]byte("user123"))
	password := fmt.Sprintf("%x", hash[:])
	newUser.Email = "user1@outlook.my"
	newUser.Password = password
	newUser.Balance = 10000
	userRepo.Register(newUser)
	t.Run("Get User Balance", func(t *testing.T) {

		_, err := transactionRepo.GetBalance(1)
		assert.Nil(t, err)

	})
	t.Run("Error Get User Balance", func(t *testing.T) {
		_, err := transactionRepo.GetBalance(100)
		assert.Error(t, err)

	})
}
func Test_Get_Resto_Detail(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	restaurantRepo := restaurants.NewRestaurantsRepo(db)
	transactionRepo := NewTransactionRepo(db)
	var newRestaurant entities.Restaurant
	hashResto := sha256.Sum256([]byte("resto123"))
	passwordResto := fmt.Sprintf("%x", hashResto[:])
	newRestaurant.Email = "resto1@outlook.my"
	newRestaurant.Password = passwordResto
	restaurantRepo.Register(newRestaurant)
	var restaurantDetail entities.RestaurantDetail
	restaurantDetail.Price = 10000
	restaurantDetail.Seats = 10
	restaurantDetail.Open = "1,2"
	restaurantDetail.Open_Hour = "08:00"
	restaurantDetail.Close_Hour = "14:00"
	restaurantDetail.Status = "OPEN"
	restaurantRepo.UpdateDetail(1, restaurantDetail)
	t.Run("Get Resto Detail", func(t *testing.T) {

		_, err := transactionRepo.GetRestoDetail(1)
		assert.Nil(t, err)

	})
	t.Run("Error Get Resto Detail", func(t *testing.T) {
		_, err := transactionRepo.GetRestoDetail(100)
		assert.Error(t, err)

	})
}
func Test_Update_User_Balance(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := users.NewUsersRepo(db)
	transactionRepo := NewTransactionRepo(db)
	var newUser entities.User
	hash := sha256.Sum256([]byte("user123"))
	password := fmt.Sprintf("%x", hash[:])
	newUser.Email = "user1@outlook.my"
	newUser.Password = password
	newUser.Balance = 10000
	userRepo.Register(newUser)

	t.Run("Update User Balance", func(t *testing.T) {
		_, err := transactionRepo.UpdateUserBalance(1, 500000)
		assert.Nil(t, err)

	})
	t.Run("Error Update User Balance", func(t *testing.T) {
		_, err := transactionRepo.UpdateUserBalance(100, 100)
		assert.Error(t, err)

	})
}
func Test_Get_Reputation_User(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := users.NewUsersRepo(db)
	transactionRepo := NewTransactionRepo(db)
	var newUser entities.User
	hash := sha256.Sum256([]byte("user123"))
	password := fmt.Sprintf("%x", hash[:])
	newUser.Email = "user1@outlook.my"
	newUser.Password = password
	newUser.Balance = 10000
	newUser.Reputation = 80
	userRepo.Register(newUser)
	t.Run("Get User Reputation", func(t *testing.T) {

		_, err := transactionRepo.GetReputationUser(1)
		assert.Nil(t, err)

	})
	t.Run("Error Get User Reputation", func(t *testing.T) {
		_, err := transactionRepo.GetReputationUser(100)
		assert.Error(t, err)

	})
}
func Test_Update_User_Reputation(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := users.NewUsersRepo(db)
	transactionRepo := NewTransactionRepo(db)
	var newUser entities.User
	hash := sha256.Sum256([]byte("user123"))
	password := fmt.Sprintf("%x", hash[:])
	newUser.Email = "user1@outlook.my"
	newUser.Password = password
	newUser.Balance = 10000
	userRepo.Register(newUser)

	t.Run("Update User Reputation", func(t *testing.T) {
		_, err := transactionRepo.UpdateUserReputation(1, 100)
		assert.Nil(t, err)

	})
	t.Run("Error Update User Reputation", func(t *testing.T) {
		_, err := transactionRepo.UpdateUserReputation(100, 100)
		assert.Error(t, err)

	})
}
func Test_Update_Transaction_Status(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	loc, _ := time.LoadLocation("Asia/Jakarta")

	userRepo := users.NewUsersRepo(db)
	restaurantRepo := restaurants.NewRestaurantsRepo(db)
	transactionRepo := NewTransactionRepo(db)

	var newUser entities.User
	hash := sha256.Sum256([]byte("user123"))
	password := fmt.Sprintf("%x", hash[:])
	newUser.Email = "user1@outlook.my"
	newUser.Password = password
	userRepo.Register(newUser)

	var newRestaurant entities.Restaurant
	hashResto := sha256.Sum256([]byte("resto123"))
	passwordResto := fmt.Sprintf("%x", hashResto[:])
	newRestaurant.Email = "resto1@outlook.my"
	newRestaurant.Password = passwordResto
	restaurantRepo.Register(newRestaurant)
	t.Run("Create Transaction", func(t *testing.T) {
		var dateTime, _ = time.ParseInLocation("2006-01-02 15:04", "2022-02-09 10:00", loc)
		var newTransaction entities.Transaction
		newTransaction.UserID = 1
		newTransaction.RestaurantID = 1
		newTransaction.DateTime = dateTime
		_, err := transactionRepo.Create(newTransaction)
		assert.Nil(t, err)

	})
	t.Run("Update Transaction Status", func(t *testing.T) {
		var updateTransaction entities.Transaction
		updateTransaction.ID = 1
		updateTransaction.Status = "Accepted"
		_, err := transactionRepo.UpdateTransactionStatus(updateTransaction)
		assert.Nil(t, err)

	})
	t.Run("Error Update Transaction Status", func(t *testing.T) {
		var updateTransaction entities.Transaction
		updateTransaction.ID = 100
		updateTransaction.Status = "Accepted"
		_, err := transactionRepo.UpdateTransactionStatus(updateTransaction)
		assert.Error(t, err)

	})

}
func Test_Get_Total_Seat(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	loc, _ := time.LoadLocation("Asia/Jakarta")

	userRepo := users.NewUsersRepo(db)
	restaurantRepo := restaurants.NewRestaurantsRepo(db)
	transactionRepo := NewTransactionRepo(db)

	var newUser entities.User
	hash := sha256.Sum256([]byte("user123"))
	password := fmt.Sprintf("%x", hash[:])
	newUser.Email = "user1@outlook.my"
	newUser.Password = password
	userRepo.Register(newUser)

	var newRestaurant entities.Restaurant
	hashResto := sha256.Sum256([]byte("resto123"))
	passwordResto := fmt.Sprintf("%x", hashResto[:])
	newRestaurant.Email = "resto1@outlook.my"
	newRestaurant.Password = passwordResto
	restaurantRepo.Register(newRestaurant)
	t.Run("Create Transaction", func(t *testing.T) {
		var dateTime, _ = time.ParseInLocation("2006-01-02 15:04", "2022-02-09 10:00", loc)
		var newTransaction entities.Transaction
		newTransaction.UserID = 1
		newTransaction.RestaurantID = 1
		newTransaction.DateTime = dateTime
		_, err := transactionRepo.Create(newTransaction)
		assert.Nil(t, err)

	})
	t.Run("Get Total Seat", func(t *testing.T) {
		_, err := transactionRepo.GetTotalSeat(1, "2022-02-09 10:00")
		assert.Nil(t, err)

	})
	t.Run("Error Get Total Seat", func(t *testing.T) {
		_, err := transactionRepo.GetTotalSeat(100, "2022-02-09 10:00")
		assert.Error(t, err)

	})

}
func Test_Check_Same_Hour(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	loc, _ := time.LoadLocation("Asia/Jakarta")

	userRepo := users.NewUsersRepo(db)
	restaurantRepo := restaurants.NewRestaurantsRepo(db)
	transactionRepo := NewTransactionRepo(db)

	var newUser entities.User
	hash := sha256.Sum256([]byte("user123"))
	password := fmt.Sprintf("%x", hash[:])
	newUser.Email = "user1@outlook.my"
	newUser.Password = password
	userRepo.Register(newUser)

	var newRestaurant entities.Restaurant
	hashResto := sha256.Sum256([]byte("resto123"))
	passwordResto := fmt.Sprintf("%x", hashResto[:])
	newRestaurant.Email = "resto1@outlook.my"
	newRestaurant.Password = passwordResto
	restaurantRepo.Register(newRestaurant)
	t.Run("Create Transaction", func(t *testing.T) {
		var dateTime, _ = time.ParseInLocation("2006-01-02 15:04", "2022-02-09 10:00", loc)
		var newTransaction entities.Transaction
		newTransaction.UserID = 1
		newTransaction.RestaurantID = 1
		newTransaction.DateTime = dateTime
		_, err := transactionRepo.Create(newTransaction)
		assert.Nil(t, err)

	})
	t.Run("Check Same Hour", func(t *testing.T) {
		_, err := transactionRepo.CheckSameHour(1, 1, "2022-02-09 10:00")
		assert.Nil(t, err)

	})
	t.Run("Error Check Same Hour", func(t *testing.T) {
		_, err := transactionRepo.CheckSameHour(100, 100, "2022-02-09 10:00")
		assert.Error(t, err)

	})

}
