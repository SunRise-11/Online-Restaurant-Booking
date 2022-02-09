package transactions

import "Restobook/entities"

type TransactionsInterface interface {
	Create(newTransaction entities.Transaction) (entities.Transaction, error)
	GetAllWaiting(userId uint) ([]entities.Transaction, error)
	GetAllWaitingForResto(restaurantId uint) ([]entities.Transaction, error)
	GetAllAcceptedForResto(restaurantId uint) ([]entities.Transaction, error)
	GetHistory(userId uint) ([]entities.Transaction, error)
	GetAllAppointed(userId uint) ([]entities.Transaction, error)
	GetBalance(userId uint) (entities.User, error)
	GetRestoDetail(restaurantId uint) (entities.RestaurantDetail, error)
	UpdateUserBalance(userId uint, balance int) (entities.User, error)
	UpdateUserReputation(userId uint, reputation int) (entities.User, error)
	UpdateTransactionStatus(newTransaction entities.Transaction) (entities.Transaction, error)
	GetTransactionById(id, userId uint, status string) (entities.Transaction, error)
	GetTotalSeat(restaurantId uint, dateTime string) (int, error)
	CheckSameHour(restaurantId, userId uint, dateTime string) (bool, error)
	GetReputationUser(userId uint) (entities.User, error)
	GetTransactionUserByStatus(id, restaurant_id uint, status string) (entities.Transaction, error)
}
