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

	tr.db.Preload("User").First(&topupdata, &topup.ID)

	return topupdata, nil
}

func (tr *TopUpRepository) GetAllWaiting(userId uint) ([]entities.TopUp, error) {
	topupdata := []entities.TopUp{}

	if err := tr.db.Where("status = ?", "PENDING").Find(&topupdata, "user_id = ?", userId).Error; err != nil {
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

func (tr *TopUpRepository) Update(extId string, topUp entities.TopUp) (entities.TopUp, error) {
	newTopUp := entities.TopUp{}

	if err := tr.db.First(&newTopUp, "invoice_id= ?", extId).Error; err != nil {
		return newTopUp, err
	}

	tr.db.Model(&newTopUp).Updates(topUp)

	return newTopUp, nil
}

func (tr *TopUpRepository) GetByInvoice(extId string) (entities.TopUp, error) {
	invoice := entities.TopUp{}

	if err := tr.db.First(&invoice, "invoice_id= ?", extId).Error; err != nil {
		return invoice, err
	}

	return invoice, nil
}

func (tr *TopUpRepository) GetUser(userId int) (entities.User, error) {
	user := entities.User{}

	if err := tr.db.First(&user, "id= ?", userId).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (tr *TopUpRepository) UpdateUserBalance(userId int, user entities.User) (entities.User, error) {
	newUserBalance := entities.User{}

	if err := tr.db.First(&newUserBalance, "id= ?", userId).Error; err != nil {
		return user, err
	}

	tr.db.Model(&newUserBalance).Updates(user)

	return user, nil
}
