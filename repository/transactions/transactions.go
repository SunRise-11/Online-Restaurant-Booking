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

	if err := tr.db.Find(&transaction).Where("user_id=? and status=waiting for confirmation", userId).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}
func (tr *TransactionRepository) GetHistory(userId uint) ([]entities.Transaction, error) {
	transaction := []entities.Transaction{}

	if err := tr.db.Find(&transaction).Where("user_id=? and status in ('Success','Fail','Cancel','Rejected'", userId).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}
func (tr *TransactionRepository) GetAllAppointed(userId uint) ([]entities.Transaction, error) {
	transaction := []entities.Transaction{}

	if err := tr.db.Find(&transaction).Where("user_id=? and status='Accepted' )", userId).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}
