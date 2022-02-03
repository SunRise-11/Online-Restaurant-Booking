package users

import (
	"Restobook/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) LoginUser(email, password string) (entities.User, error) {
	var user entities.User

	if err := ur.db.Where("Email = ? AND Password = ?", email, password).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
func (ur *UserRepository) Register(newUser entities.User) (entities.User, error) {
	err := ur.db.Save(&newUser).Error

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
func (ur *UserRepository) Delete(userId uint) (entities.User, error) {
	user := entities.User{}

	if err := ur.db.First(&user, "id=?", userId).Error; err != nil {
		return user, err
	}

	ur.db.Delete(&user)

	return user, nil
}
func (ur *UserRepository) Update(userId uint, newUser entities.User) (entities.User, error) {
	user := entities.User{}

	if err := ur.db.First(&user, "id=?", userId).Error; err != nil {
		return user, err
	}

	ur.db.Model(&user).Updates(newUser)

	return user, nil
}
func (ur *UserRepository) Get(userId uint) (entities.User, error) {
	user := entities.User{}

	if err := ur.db.First(&user, userId).Error; err != nil {
		return user, err
	}

	return user, nil
}
