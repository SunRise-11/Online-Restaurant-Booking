package topup

import (
	"Restobook/configs"
	"Restobook/entities"
	"Restobook/repository/users"
	"Restobook/utils"
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTopup(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := users.NewUsersRepo(db)
	topupRepo := NewTopUpRepo(db)

	//Register User
	var newUser entities.User
	hash := sha256.Sum256([]byte("resto123"))
	password := fmt.Sprintf("%x", hash[:])
	newUser.Email = "restaurant1@outlook.my"
	newUser.Password = password
	userRepo.Register(newUser)

	t.Run("Create Top Up Success", func(t *testing.T) {
		newTopup := entities.TopUp{
			ID:         1,
			UserID:     1,
			InvoiceID:  "XENDIT INVOICE ID",
			PaymentUrl: "XENDIT PAYMENT URL",
			Total:      10000,
			Status:     "PENDING",
		}

		res, err := topupRepo.Create(newTopup)
		assert.Nil(t, err)
		assert.Equal(t, uint(1), res.ID)
	})

	db.Migrator().DropTable(&entities.TopUp{})

	t.Run("Create Top Up Error", func(t *testing.T) {
		newTopup := entities.TopUp{
			ID:         1,
			UserID:     1,
			InvoiceID:  "XENDIT INVOICE ID",
			PaymentUrl: "XENDIT PAYMENT URL",
			Total:      10000,
			Status:     "PENDING",
		}

		_, err := topupRepo.Create(newTopup)
		assert.NotNil(t, err)
	})
}

func TestGetAllWaitingTopup(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := users.NewUsersRepo(db)
	topupRepo := NewTopUpRepo(db)

	//Register User
	var newUser entities.User
	hash := sha256.Sum256([]byte("resto123"))
	password := fmt.Sprintf("%x", hash[:])
	newUser.Email = "restaurant1@outlook.my"
	newUser.Password = password
	userRepo.Register(newUser)

	//Create Topup
	newTopup := entities.TopUp{
		ID:         1,
		UserID:     1,
		InvoiceID:  "XENDIT INVOICE ID",
		PaymentUrl: "XENDIT PAYMENT URL",
		Total:      10000,
		Status:     "PENDING",
	}
	topupRepo.Create(newTopup)

	t.Run("Get All Waiting Success", func(t *testing.T) {
		res, err := topupRepo.GetAllWaiting(1)
		assert.Nil(t, err)
		assert.Equal(t, uint(1), res[0].UserID)
	})

	db.Migrator().DropTable(&entities.TopUp{})

	t.Run("Get All Waiting Error", func(t *testing.T) {
		_, err := topupRepo.GetAllWaiting(1)
		assert.NotNil(t, err)
	})
}

func TestGetAllPaidTopup(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := users.NewUsersRepo(db)
	topupRepo := NewTopUpRepo(db)

	//Register User
	var newUser entities.User
	hash := sha256.Sum256([]byte("resto123"))
	password := fmt.Sprintf("%x", hash[:])
	newUser.Email = "restaurant1@outlook.my"
	newUser.Password = password
	userRepo.Register(newUser)

	//Create Topup
	newTopup := entities.TopUp{
		ID:         1,
		UserID:     1,
		InvoiceID:  "XENDIT INVOICE ID",
		PaymentUrl: "XENDIT PAYMENT URL",
		Total:      10000,
		Status:     "PAID",
	}
	topupRepo.Create(newTopup)

	t.Run("Get All Paid Success", func(t *testing.T) {
		res, err := topupRepo.GetAllPaid(1)
		assert.Nil(t, err)
		assert.Equal(t, uint(1), res[0].UserID)
	})

	db.Migrator().DropTable(&entities.TopUp{})

	t.Run("Get All Paid Error", func(t *testing.T) {
		_, err := topupRepo.GetAllPaid(1)
		assert.NotNil(t, err)
	})
}

func TestUpdateTopup(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := users.NewUsersRepo(db)
	topupRepo := NewTopUpRepo(db)

	//Register User
	var newUser entities.User
	hash := sha256.Sum256([]byte("resto123"))
	password := fmt.Sprintf("%x", hash[:])
	newUser.Email = "restaurant1@outlook.my"
	newUser.Password = password
	userRepo.Register(newUser)

	//Create Topup
	newTopup := entities.TopUp{
		ID:         1,
		UserID:     1,
		InvoiceID:  "XENDIT INVOICE ID",
		PaymentUrl: "XENDIT PAYMENT URL",
		Total:      10000,
		Status:     "PENDING",
	}
	topupRepo.Create(newTopup)

	t.Run("Update Topup Success", func(t *testing.T) {
		updateTopup := entities.TopUp{
			ID:         1,
			UserID:     1,
			InvoiceID:  "XENDIT INVOICE ID",
			PaymentUrl: "XENDIT PAYMENT URL",
			Total:      10000,
			Status:     "PAID",
		}
		res, err := topupRepo.Update("XENDIT INVOICE ID", updateTopup)
		assert.Nil(t, err)
		assert.Equal(t, "PAID", res.Status)
	})

	db.Migrator().DropTable(&entities.TopUp{})

	t.Run("Get All Paid Error", func(t *testing.T) {
		updateTopup := entities.TopUp{}
		_, err := topupRepo.Update("XENDIT INVOICE ID", updateTopup)
		assert.NotNil(t, err)
	})
}

func TestGetByInvoiceTopup(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := users.NewUsersRepo(db)
	topupRepo := NewTopUpRepo(db)

	//Register User
	var newUser entities.User
	hash := sha256.Sum256([]byte("resto123"))
	password := fmt.Sprintf("%x", hash[:])
	newUser.Email = "restaurant1@outlook.my"
	newUser.Password = password
	userRepo.Register(newUser)

	//Create Topup
	newTopup := entities.TopUp{
		ID:         1,
		UserID:     1,
		InvoiceID:  "XENDIT INVOICE ID",
		PaymentUrl: "XENDIT PAYMENT URL",
		Total:      10000,
		Status:     "PENDING",
	}
	topupRepo.Create(newTopup)

	t.Run("Get By Invoice Success", func(t *testing.T) {
		res, err := topupRepo.GetByInvoice("XENDIT INVOICE ID")
		assert.Nil(t, err)
		assert.Equal(t, "PENDING", res.Status)
	})

	db.Migrator().DropTable(&entities.TopUp{})

	t.Run("Get By Invoice Error", func(t *testing.T) {
		_, err := topupRepo.GetByInvoice("XENDIT INVOICE ID")
		assert.NotNil(t, err)
	})
}

func TestGetUserTopup(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := users.NewUsersRepo(db)
	topupRepo := NewTopUpRepo(db)

	//Register User
	var newUser entities.User
	hash := sha256.Sum256([]byte("resto123"))
	password := fmt.Sprintf("%x", hash[:])
	newUser.Email = "restaurant1@outlook.my"
	newUser.Password = password
	userRepo.Register(newUser)

	t.Run("Get User Success", func(t *testing.T) {
		res, err := topupRepo.GetUser(1)
		assert.Nil(t, err)
		assert.Equal(t, uint(1), res.ID)
	})

	db.Migrator().DropTable(&entities.User{})

	t.Run("Get User Error", func(t *testing.T) {
		_, err := topupRepo.GetUser(1)
		assert.NotNil(t, err)
	})
}

func TestUpdateUserBalanceTopup(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := users.NewUsersRepo(db)
	topupRepo := NewTopUpRepo(db)

	//Register User
	var newUser entities.User
	hash := sha256.Sum256([]byte("resto123"))
	password := fmt.Sprintf("%x", hash[:])
	newUser.Email = "restaurant1@outlook.my"
	newUser.Password = password
	userRepo.Register(newUser)

	t.Run("Update User Balance Success", func(t *testing.T) {
		var updateUser entities.User
		newUser.Balance = 10000
		_, err := topupRepo.UpdateUserBalance(1, updateUser)
		assert.Nil(t, err)
	})

	db.Migrator().DropTable(&entities.User{})

	t.Run("Update User Balance Error", func(t *testing.T) {
		var updateUser entities.User
		newUser.Balance = 10000
		_, err := topupRepo.UpdateUserBalance(1, updateUser)
		assert.NotNil(t, err)
	})
}
