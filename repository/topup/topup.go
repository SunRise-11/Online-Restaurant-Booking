package topup

import (
	"Restobook/entities"

	"gorm.io/gorm"
)

type TopUpRepository struct {
	db *gorm.DB
}

func NewTopUpRepo(db *gorm.DB) *TopUpRepository {
	return &TopUpRepository{db: db}
}

func (tr *TopUpRepository) Create(topup entities.TopUp) (entities.TopUp, error) {
	if err := tr.db.Create(&topup).Error; err != nil {
		return topup, err
	}

	topupdata := entities.TopUp{}

	if err := tr.db.Preload("User").First(&topupdata, &topup.ID).Error; err != nil {
		return topup, err
	}

	return topupdata, nil
}

func (tr *TopUpRepository) GetAllWaiting(userId uint) ([]entities.TopUp, error) {
	topupdata := []entities.TopUp{}

	if err := tr.db.Where("status = ?", "waiting for payment").Find(&topupdata, "user_id = ?", userId).Error; err != nil {
		return topupdata, err
	}

	return topupdata, nil
}

func (tr *TopUpRepository) GetAllPaid(userId uint) ([]entities.TopUp, error) {
	topupdata := []entities.TopUp{}

	if err := tr.db.Where("status = ?", "PAID").Find(&topupdata, "user_id = ?", userId).Error; err != nil {
		return topupdata, err
	}

	return topupdata, nil
}

func (tr *TopUpRepository) Update(invId string, topUp entities.TopUp) (entities.TopUp, error) {
	newTopUp := entities.TopUp{}

	if err := tr.db.First(&newTopUp, "invoice_id = ?", invId).Error; err != nil {
		return newTopUp, err
	}

	tr.db.Model(&newTopUp).Updates(topUp)

	return newTopUp, nil
}
