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

func (ur *UserRepository) RegisterAdmin(newAdmin entities.User) (entities.User, error) {
	admin := entities.User{
		ID:       1,
		Name:     newAdmin.Name,
		Email:    newAdmin.Email,
		Password: newAdmin.Password,
	}

	if err := ur.db.Save(&admin).Error; err != nil {
		return admin, err
	} else {
		return admin, nil
	}

}

func (ur *UserRepository) Register(newUser entities.User) (entities.User, error) {
	if err := ur.db.Save(&newUser).Error; err != nil {
		return newUser, err
	} else {
		return newUser, nil
	}

}

func (ur *UserRepository) LoginUser(email, password string) (entities.User, error) {
	var user entities.User

	if err := ur.db.Where("Email = ? AND Password=?", email, password).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
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
