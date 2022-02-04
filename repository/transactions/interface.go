package transactions

import "Restobook/entities"

type TransactionsInterface interface {
	Create(newTransaction entities.Transaction) (entities.Transaction, error)
	GetAllWaiting(userId uint) ([]entities.Transaction, error)
	GetHistory(userId uint) ([]entities.Transaction, error)
	GetAllAppointed(userId uint) ([]entities.Transaction, error)
}
