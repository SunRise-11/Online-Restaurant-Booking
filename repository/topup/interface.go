package topup

import "Restobook/entities"

type TopUpInterface interface {
	Create(topup entities.TopUp) (entities.TopUp, error)
	GetAllWaiting(userId uint) ([]entities.TopUp, error)
	GetAllPaid(userId uint) ([]entities.TopUp, error)
	Update(invId string, topUp entities.TopUp) (entities.TopUp, error)
	GetByInvoice(invId string) (entities.TopUp, error)
}
