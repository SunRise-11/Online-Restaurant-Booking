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

func TestUsersRepo(t *testing.T) {
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

	t.Run("Show UserID 2", func(t *testing.T) {
		res, err := userRepo.Get(2)
		assert.Equal(t, res.ID, uint(2))
		assert.Nil(t, err)
	})

	t.Run("Login UserID 2", func(t *testing.T) {
		hash := sha256.Sum256([]byte("herlianto123"))
		password := fmt.Sprintf("%x", hash[:])
		var loginUser entities.User
		loginUser.Email = "herlianto@outlook.my"
		loginUser.Password = password

		res, err := userRepo.LoginUser(loginUser.Email, loginUser.Password)
		assert.Equal(t, res.ID, uint(2))
		assert.Nil(t, err)
	})

	t.Run("Update UserID 2", func(t *testing.T) {
		hash := sha256.Sum256([]byte("herlianto123"))
		password := fmt.Sprintf("%x", hash[:])
		var updateUser entities.User
		updateUser.Name = "herlianto"
		updateUser.Email = "herlianto@outlook.my"
		updateUser.Password = password
		updateUser.PhoneNumber = "0877"
		updateUser.Balance = 100.000
		updateUser.Reputation = 100

		res, err := userRepo.Update(uint(2), updateUser)
		assert.Equal(t, res.ID, uint(2))
		assert.Nil(t, err)
	})

	t.Run("Delete UserID 2", func(t *testing.T) {
		res, err := userRepo.Delete(2)
		assert.Equal(t, res.ID, uint(2))
		assert.Nil(t, err)
	})
}

func TestFalseUsersRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := NewUsersRepo(db)

	t.Run("Register User", func(t *testing.T) {
		hash := sha256.Sum256([]byte("herlianto123"))
		password := fmt.Sprintf("%x", hash[:])
		var newUser entities.User
		newUser.ID = 3
		newUser.Name = "herlianto"
		newUser.Email = "herlianto@outlook.my"
		newUser.Password = password
		newUser.PhoneNumber = "0877"

		res, err := userRepo.Register(newUser)
		assert.Equal(t, res, res)
		assert.Nil(t, err)
	})

	t.Run("FALSE Register User", func(t *testing.T) {
		hash := sha256.Sum256([]byte("herlianto123"))
		password := fmt.Sprintf("%x", hash[:])
		var newUser entities.User
		newUser.ID = 3
		newUser.Name = "herlianto"
		newUser.Email = "herlianto@outlook.my"
		newUser.Password = password
		newUser.PhoneNumber = "0877"

		res, err := userRepo.Register(newUser)
		assert.Equal(t, res.ID, uint(3))
		assert.Error(t, err)
	})

	t.Run("FALSE Show UserID 4", func(t *testing.T) {
		res, err := userRepo.Get(4)
		assert.Equal(t, res.ID, uint(0))
		assert.Error(t, err)
	})

	t.Run("FALSE Login User", func(t *testing.T) {
		hash := sha256.Sum256([]byte("herlianto1234"))
		password := fmt.Sprintf("%x", hash[:])
		var loginUser entities.User
		loginUser.Email = "herlianto@outlook.my"
		loginUser.Password = password

		res, err := userRepo.LoginUser(loginUser.Email, loginUser.Password)
		assert.Equal(t, res.ID, uint(0))
		assert.Error(t, err)
	})

	t.Run("FALSE Update UserID 4", func(t *testing.T) {
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

	t.Run("FALSE Delete UserID 4", func(t *testing.T) {
		res, err := userRepo.Delete(4)
		assert.Equal(t, res.ID, uint(0))
		assert.Error(t, err)
	})
}
