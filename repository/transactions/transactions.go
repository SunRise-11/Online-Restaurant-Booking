package transactions

import (
	"Restobook/entities"
	"errors"

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
	if err != nil || newTransaction.ID == 0 {
		return newTransaction, errors.New("FAILED CREATE TRANSACTION")
	}

	return newTransaction, nil
}
func (tr *TransactionRepository) GetAllWaiting(userId uint) ([]entities.Transaction, error) {
	transaction := []entities.Transaction{}

	if err := tr.db.Where("user_id=? and status=?", userId, "waiting for confirmation").Find(&transaction).Error; err != nil || len(transaction) == 0 {
		return transaction, errors.New("FAILED GET DATA")
	}

	return transaction, nil
}
func (tr *TransactionRepository) GetAllWaitingForResto(restaurantId uint) ([]entities.Transaction, error) {
	transaction := []entities.Transaction{}

	if err := tr.db.Where("restaurant_id=? and status=?", restaurantId, "waiting for confirmation").Find(&transaction).Error; err != nil || len(transaction) == 0 {
		return transaction, errors.New("FAILED GET DATA")
	}

	return transaction, nil
}
func (tr *TransactionRepository) GetAllAcceptedForResto(restaurantId uint) ([]entities.Transaction, error) {
	transaction := []entities.Transaction{}

	if err := tr.db.Where("restaurant_id=? and status=?", restaurantId, "Accepted").Find(&transaction).Error; err != nil || len(transaction) == 0 {
		return transaction, errors.New("FAILED GET DATA")
	}
	return transaction, nil
}
func (tr *TransactionRepository) GetHistory(userId uint) ([]entities.Transaction, error) {
	transaction := []entities.Transaction{}

	if err := tr.db.Where("user_id=?", userId).Where("status in ?", []string{"Success", "Fail", "Cancel", "Rejected", "Dismissed"}).Find(&transaction).Error; err != nil || len(transaction) == 0 {
		return transaction, errors.New("FAILED GET DATA")
	}
	return transaction, nil
}
func (tr *TransactionRepository) GetAllAppointed(userId uint) ([]entities.Transaction, error) {
	transaction := []entities.Transaction{}

	if err := tr.db.Where("user_id=? and status=? ", userId, "Accepted").Find(&transaction).Error; err != nil || len(transaction) == 0 {
		return transaction, errors.New("FAILED GET DATA")
	}
	return transaction, nil
}
func (tr *TransactionRepository) GetTransactionById(id, userId uint) (entities.Transaction, error) {
	transaction := entities.Transaction{}
	if err := tr.db.Preload("User").Where("id=? and user_id=?", id, userId).First(&transaction).Error; err != nil || transaction.ID == 0 {
		return transaction, errors.New("FAILED GET DATA")
	}

	return transaction, nil
}
func (tr *TransactionRepository) GetTransactionUserByStatus(id, restaurant_id uint, status string) (entities.Transaction, error) {
	transaction := entities.Transaction{}
	if err := tr.db.Preload("User").Preload("Restaurant.RestaurantDetail").Where("id=? and restaurant_id=? and status=?", id, restaurant_id, status).First(&transaction).Error; err != nil || transaction.ID == 0 {
		return transaction, errors.New("FAILED GET DATA")
	}

	return transaction, nil
}

func (tr *TransactionRepository) GetBalance(userId uint) (entities.User, error) {
	user := entities.User{}

	if err := tr.db.Select("balance").Where("id=?", userId).First(&user).Error; err != nil {
		return user, errors.New("FAILED GET DATA")
	}

	return user, nil
}
func (tr *TransactionRepository) GetRestoDetail(restaurantId uint) (entities.RestaurantDetail, error) {
	resto := entities.RestaurantDetail{}
	if err := tr.db.Select("price", "seats", "open", "open_hour", "close_hour", "status").Where("id=?", restaurantId).First(&resto).Error; err != nil {
		return resto, errors.New("FAILED GET DATA")
	}
	return resto, nil
}

func (tr *TransactionRepository) UpdateUserBalance(userId uint, balance int) (entities.User, error) {
	user := entities.User{}
	updateUser := make(map[string]interface{})
	if err := tr.db.First(&user, "id=?", userId).Error; err != nil || user.ID == 0 {
		return user, errors.New("FAILED UPDATE DATA")
	}
	updateUser["balance"] = balance
	tr.db.Model(&user).Updates(&updateUser)
	return user, nil

}
func (tr *TransactionRepository) GetReputationUser(userId uint) (entities.User, error) {
	user := entities.User{}
	if err := tr.db.Select("reputation").Where("id=?", userId).First(&user).Error; err != nil {
		return user, errors.New("FAILED GET DATA")
	}

	return user, nil

}
func (tr *TransactionRepository) UpdateUserReputation(userId uint, reputation int) (entities.User, error) {
	user := entities.User{}
	updateUser := make(map[string]interface{})
	if err := tr.db.First(&user, "id=?", userId).Error; err != nil || user.ID == 0 {
		return user, errors.New("FAILED UPDATE DATA")
	}
	updateUser["reputation"] = reputation
	tr.db.Model(&user).Updates(&updateUser)
	return user, nil

}
func (tr *TransactionRepository) UpdateTransactionStatus(newTransaction entities.Transaction) (entities.Transaction, error) {
	transaction := entities.Transaction{}
	if err := tr.db.First(&transaction, "id=?", newTransaction.ID).Error; err != nil || transaction.ID == 0 {
		return transaction, errors.New("FAILED UPDATE DATA")
	}
	tr.db.Model(&transaction).Updates(newTransaction)
	return transaction, nil

}

func (tr *TransactionRepository) GetTotalSeat(restaurantId uint, dateTime string) (int, error) {
	var result int
	err := tr.db.Model(&entities.Transaction{}).Select("sum(persons) as total").Where("date_time=?", dateTime).Where("restaurant_id=?", restaurantId).Find(&result).Error
	if err != nil {
		return result, err
	}
	return result, nil
}
func (tr *TransactionRepository) CheckSameHour(restaurantId, userId uint, dateTime string) (bool, error) {
	transaction := entities.Transaction{}
	if err := tr.db.First(&transaction, "user_id=? and restaurant_id=? and date_time=?", userId, restaurantId, dateTime).Error; err != nil || transaction.ID == 0 {
		return false, errors.New("FAILED GET DATA")
	}
	return true, nil
}
