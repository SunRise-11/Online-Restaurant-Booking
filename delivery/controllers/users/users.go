package users

import (
	"Restobook/delivery/common"
	"Restobook/entities"
	"Restobook/repository/users"
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UsersController struct {
	Repo users.UsersInterface
}

func NewUsersControllers(usrep users.UsersInterface) *UsersController {
	return &UsersController{Repo: usrep}
}

func (uscon UsersController) RegisterUserCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		newUserReq := UserRequestFormat{}
		if err := c.Bind(&newUserReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		hash := sha256.Sum256([]byte(newUserReq.Password))
		stringPassword := fmt.Sprintf("%x", hash[:])
		newUser := entities.User{
			Email:    newUserReq.Email,
			Password: stringPassword,
			Name:     newUserReq.Name,
		}
		res, err := uscon.Repo.Register(newUser)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		data := UserResponse{
			ID:    res.ID,
			Name:  res.Name,
			Email: res.Email,
		}
		response := UserResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    data,
		}

		return c.JSON(http.StatusOK, response)
	}
}
