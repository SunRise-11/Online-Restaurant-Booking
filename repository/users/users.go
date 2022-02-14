package users

import (
	"Restobook/entities"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) RegisterAdmin(newAdmin entities.User) (entities.User, error) {
	if err := ur.db.Save(&newAdmin).Error; err != nil {
		return newAdmin, errors.New("FAILED REGISTER ADMIN")
	} else {
		return newAdmin, nil
	}

}

func (ur *UserRepository) Register(newUser entities.User) (entities.User, error) {
	if err := ur.db.Save(&newUser).Error; err != nil || newUser.ID == 0 {
		return newUser, errors.New("FAILED REGISTER USER")
	} else {
		return newUser, nil
	}

}

func (ur *UserRepository) LoginUser(email, password string) (entities.User, error) {
	var user entities.User

	if err := ur.db.Where("Email = ? AND Password = ?", email, password).First(&user).Error; err != nil || user.ID == 0 {
		return user, errors.New("FAILED LOGIN")
	} else {
		return user, nil
	}

}

func (ur *UserRepository) Get(userId uint) (entities.User, error) {
	user := entities.User{}

	if err := ur.db.First(&user, userId).Error; err != nil || user.ID == 0 {
		return user, errors.New("FAILED GET")
	} else {
		return user, nil
	}

}

func (ur *UserRepository) Update(userId uint, updateUser entities.User) (entities.User, error) {
	user := entities.User{}

	if err := ur.db.First(&user, "id=?", userId).Error; err != nil || user.ID == 0 {
		return user, errors.New("FAILED UPDATE")
	} else {
		ur.db.Model(&user).Updates(updateUser)
		return user, nil
	}

}

func (ur *UserRepository) Delete(userId uint) (entities.User, error) {
	user := entities.User{}

	if err := ur.db.First(&user, "id=?", userId).Error; err != nil || user.ID == 0 {
		return user, errors.New("FAILED DELETE")
	} else {
		ur.db.Delete(&user)
		return user, nil
	}

}
