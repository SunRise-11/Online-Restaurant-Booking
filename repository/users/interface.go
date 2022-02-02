package users

import "Restobook/entities"

type UsersInterface interface {
	Gets() ([]entities.User, error)
	LoginUser(email string, password string) (entities.User, error)
	Register(newUser entities.User) (entities.User, error)
	Delete(userId int) (entities.User, error)
	Update(newUser entities.User, userId int) (entities.User, error)
	Get(userId int) (entities.User, error)
}
