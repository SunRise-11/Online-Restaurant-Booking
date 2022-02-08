package transactions

import "Restobook/entities"

type TransactionsInterface interface {
	Create(newTransaction entities.Transaction) (entities.Transaction, error)
	GetAllWaiting(userId uint) ([]entities.Transaction, error)
	GetAllWaitingForResto(restaurantId uint) ([]entities.Transaction, error)
	GetHistory(userId uint) ([]entities.Transaction, error)
	GetAllAppointed(userId uint) ([]entities.Transaction, error)
	GetBalance(userId uint) (entities.User, error)
	GetRestoDetail(restaurantId uint) (entities.RestaurantDetail, error)
	UpdateUserBalance(userId uint, balance int) (entities.User, error)
	UpdateTransactionStatus(newTransaction entities.Transaction) (entities.Transaction, error)
	ShowAllTransaction(restaurantId uint) ([]entities.Transaction, error)
	GetTransactionById(id uint) (entities.Transaction, error)
	GetTotalSeat(restaurantId uint, dateTime string) (int, error)
	CheckSameHour(restaurantId, userId uint, dateTime string) (bool, error)
}

type BalanceAndPriceResto struct {
	Balance    int
	Seats      int
	PriceResto int
}
