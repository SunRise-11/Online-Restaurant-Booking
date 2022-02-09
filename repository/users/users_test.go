package users

import (
	"Restobook/configs"
	"Restobook/entities"
	"Restobook/utils"
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterAdminRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	userRepo := NewUsersRepo(db)

	t.Run("Register Admin", func(t *testing.T) {
		hash := sha256.Sum256([]byte("herlianto123"))
		password := fmt.Sprintf("%x", hash[:])
		var newUser entities.User
		newUser.Name = "admin"
		newUser.Email = "admin@outlook.my"
		newUser.Password = password
		newUser.PhoneNumber = "0877"

		res, err := userRepo.RegisterAdmin(newUser)
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})
	t.Run("ERROR Register Admin", func(t *testing.T) {
		hash := sha256.Sum256([]byte("herlianto123"))
		password := fmt.Sprintf("%x", hash[:])
		var newUser entities.User
		newUser.Name = "admin"
		newUser.Email = "admin@outlook.my"
		newUser.Password = password
		newUser.PhoneNumber = "0877"

		res, err := userRepo.RegisterAdmin(newUser)
		assert.Equal(t, res, res)
		assert.Error(t, err)
	})
}

func TestRegisterUsersRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := NewUsersRepo(db)

	t.Run("Register User", func(t *testing.T) {
		hash := sha256.Sum256([]byte("herlianto123"))
		password := fmt.Sprintf("%x", hash[:])
		var newUser entities.User
		newUser.Name = "herlianto"
		newUser.Email = "herlianto@outlook.my"
		newUser.Password = password
		newUser.PhoneNumber = "0877"

		res, err := userRepo.Register(newUser)
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

	t.Run("ERROR Register User", func(t *testing.T) {
		hash := sha256.Sum256([]byte("herlianto123"))
		password := fmt.Sprintf("%x", hash[:])
		var newUser entities.User
		newUser.ID = 1
		newUser.Name = "herlianto"
		newUser.Email = "herlianto@outlook.my"
		newUser.Password = password
		newUser.PhoneNumber = "0877"

		res, err := userRepo.Register(newUser)
		assert.Equal(t, res, res)
		assert.Error(t, err)
	})

}

func TestLoginUsersRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := NewUsersRepo(db)

	t.Run("Register User", func(t *testing.T) {
		hash := sha256.Sum256([]byte("herlianto123"))
		password := fmt.Sprintf("%x", hash[:])
		var newUser entities.User
		newUser.Name = "herlianto"
		newUser.Email = "herlianto@outlook.my"
		newUser.Password = password
		newUser.PhoneNumber = "0877"

		res, err := userRepo.Register(newUser)
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

	t.Run("Login User", func(t *testing.T) {
		hash := sha256.Sum256([]byte("herlianto123"))
		password := fmt.Sprintf("%x", hash[:])
		var loginUser entities.User
		loginUser.Email = "herlianto@outlook.my"
		loginUser.Password = password

		res, err := userRepo.LoginUser(loginUser.Email, loginUser.Password)
		assert.Equal(t, res.ID, uint(1))
		assert.Nil(t, err)
	})

	t.Run("ERROR Login User", func(t *testing.T) {
		hash := sha256.Sum256([]byte("herlianto1234"))
		password := fmt.Sprintf("%x", hash[:])
		var loginUser entities.User
		loginUser.Email = "herlianto@outlook.my"
		loginUser.Password = password

		res, err := userRepo.LoginUser(loginUser.Email, loginUser.Password)
		assert.Equal(t, res.ID, uint(0))
		assert.Error(t, err)
	})
}

func TestDeleteUsersRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := NewUsersRepo(db)

	t.Run("Register User", func(t *testing.T) {
		hash := sha256.Sum256([]byte("herlianto123"))
		password := fmt.Sprintf("%x", hash[:])
		var newUser entities.User
		newUser.Name = "herlianto"
		newUser.Email = "herlianto@outlook.my"
		newUser.Password = password
		newUser.PhoneNumber = "0877"

		res, err := userRepo.Register(newUser)
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

	t.Run("Delete User", func(t *testing.T) {
		res, err := userRepo.Delete(1)
		assert.Equal(t, res.ID, uint(1))
		assert.Nil(t, err)
	})

	t.Run("ERROR Delete User", func(t *testing.T) {
		res, err := userRepo.Delete(1)
		assert.Equal(t, res.ID, uint(0))
		assert.Error(t, err)
	})
}

func TestUpdateUsersRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := NewUsersRepo(db)

	t.Run("Register User", func(t *testing.T) {
		hash := sha256.Sum256([]byte("herlianto123"))
		password := fmt.Sprintf("%x", hash[:])
		var newUser entities.User
		newUser.Name = "herlianto"
		newUser.Email = "herlianto@outlook.my"
		newUser.Password = password
		newUser.PhoneNumber = "0877"

		res, err := userRepo.Register(newUser)
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

	t.Run("Update User", func(t *testing.T) {
		hash := sha256.Sum256([]byte("herlianto123"))
		password := fmt.Sprintf("%x", hash[:])
		var updateUser entities.User
		updateUser.Name = "herlianto"
		updateUser.Email = "herlianto@outlook.my"
		updateUser.Password = password
		updateUser.PhoneNumber = "0877"
		updateUser.Balance = 100.000
		updateUser.Reputation = 100

		res, err := userRepo.Update(uint(1), updateUser)
		assert.Equal(t, res.ID, uint(1))
		assert.Nil(t, err)
	})

	t.Run("ERROR Update User", func(t *testing.T) {
		hash := sha256.Sum256([]byte("herlianto123"))
		password := fmt.Sprintf("%x", hash[:])
		var updateUser entities.User
		updateUser.Name = "herlianto"
		updateUser.Email = "herlianto@outlook.my"
		updateUser.Password = password
		updateUser.PhoneNumber = "0877"
		updateUser.Balance = 100.000
		updateUser.Reputation = 100

		res, err := userRepo.Update(uint(4), updateUser)
		assert.Equal(t, res.ID, uint(0))
		assert.Error(t, err)
	})
}

func TestGetUsersRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := NewUsersRepo(db)

	t.Run("Register User", func(t *testing.T) {
		hash := sha256.Sum256([]byte("herlianto123"))
		password := fmt.Sprintf("%x", hash[:])
		var newUser entities.User
		newUser.Name = "herlianto"
		newUser.Email = "herlianto@outlook.my"
		newUser.Password = password
		newUser.PhoneNumber = "0877"

		res, err := userRepo.Register(newUser)
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

	t.Run("Get User", func(t *testing.T) {
		res, err := userRepo.Get(uint(1))
		assert.Equal(t, res.ID, uint(1))
		assert.Nil(t, err)
	})

	t.Run("ERROR Get User", func(t *testing.T) {
		res, err := userRepo.Get(uint(2))
		assert.Equal(t, res.ID, uint(0))
		assert.Error(t, err)
	})
}
