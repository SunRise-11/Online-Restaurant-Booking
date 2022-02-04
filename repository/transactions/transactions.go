package transactions

import (
	"Restobook/entities"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (tr *TransactionRepository) Create(newTransaction entities.Transaction) (entities.Transaction, error) {
	err := tr.db.Save(&newTransaction).Error
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}
func (tr *TransactionRepository) GetAllWaiting(userId uint) ([]entities.Transaction, error) {
	transaction := []entities.Transaction{}

	if err := tr.db.Where("user_id=? and status=?", userId, "waiting for confirmation").Find(&transaction).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}
func (tr *TransactionRepository) GetHistory(userId uint) ([]entities.Transaction, error) {
	transaction := []entities.Transaction{}

	if err := tr.db.Where("user_id=?", userId).Where("status in ?", []string{"Success", "Fail", "Cancel", "Rejected"}).Find(&transaction).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}
func (tr *TransactionRepository) GetAllAppointed(userId uint) ([]entities.Transaction, error) {
	transaction := []entities.Transaction{}

	if err := tr.db.Where("user_id=? and status=? ", userId, "Accepted").Find(&transaction).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}
func (tr *TransactionRepository) GetTransactionById(id uint) (entities.Transaction, error) {
	transaction := entities.Transaction{}
	if err := tr.db.Preload("User").First(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}
func (tr *TransactionRepository) ShowAllTransaction(restaurantId uint) ([]entities.Transaction, error) {
	transaction := []entities.Transaction{}

	if err := tr.db.Where("restaurant_id=? and status=? ", restaurantId, "waiting for confirmation").Find(&transaction).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}
func (tr *TransactionRepository) GetBalanceAndPriceResto(userId, restaurantId uint) (BalanceAndPriceResto, error) {
	user := entities.User{}
	resto := entities.RestaurantDetail{}
	if err := tr.db.Select("price", "seats").Where("id=?", restaurantId).First(&resto).Error; err != nil {
		return BalanceAndPriceResto{PriceResto: resto.Price, Seats: resto.Seats}, err
	}
	if err := tr.db.Select("balance").Where("id=?", userId).First(&user).Error; err != nil {
		return BalanceAndPriceResto{Balance: user.Balance}, err
	}

	return BalanceAndPriceResto{Balance: user.Balance, PriceResto: resto.Price, Seats: resto.Seats}, nil
}

func (tr *TransactionRepository) UpdateUserBalance(userId uint, balance int) (entities.User, error) {
	user := entities.User{}
	updateUser := make(map[string]interface{})
	if err := tr.db.First(&user, "id=?", userId).Error; err != nil {
		return user, err
	}
	updateUser["balance"] = balance
	tr.db.Model(&user).Updates(&updateUser)
	return user, nil

}

func (tr *TransactionRepository) UpdateTransactionStatus(newTransaction entities.Transaction) (entities.Transaction, error) {
	transaction := entities.Transaction{}
	if err := tr.db.First(&transaction, "id=?", newTransaction.ID).Error; err != nil {
		return transaction, err
	}
	tr.db.Model(&transaction).Updates(newTransaction)
	return transaction, nil

}
