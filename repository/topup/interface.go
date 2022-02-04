package topup

import "Restobook/entities"

type TopUpInterface interface {
	Create(topup entities.TopUp) (entities.TopUp, error)
	GetAllWaiting(userId uint) ([]entities.TopUp, error)
	GetAllPaid(userId uint) ([]entities.TopUp, error)
	Update(extId string, topUp entities.TopUp) (entities.TopUp, error)
	GetByInvoice(extId string) (entities.TopUp, error)
	GetUser(userId int) (entities.User, error)
	UpdateUserBalance(userId int, user entities.User) (entities.User, error)
}
