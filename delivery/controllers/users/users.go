package users

import (
	"Restobook/delivery/common"
	"Restobook/delivery/controllers/auth"
	"Restobook/entities"
	"Restobook/repository/users"
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
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
		if res, err := uscon.Repo.Register(newUser); err != nil || res.ID == 0 {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		} else {
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
}

func (uscon UsersController) LoginAuthCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		loginFormat := LoginRequestFormat{}
		if err := c.Bind(&loginFormat); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		hash := sha256.Sum256([]byte(loginFormat.Password))
		stringPassword := fmt.Sprintf("%x", hash[:])
		if res, err := uscon.Repo.LoginUser(loginFormat.Email, stringPassword); err != nil || res.Email == "" || res.ID == 0 {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {
			token, _ := auth.CreateTokenAuth(res.ID)

			return c.JSON(http.StatusOK, LoginResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Token:   token,
			})
		}

	}
}

func (uscon UsersController) GetUserByIdCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		id := int(claims["userid"].(float64))
		if res, err := uscon.Repo.Get(id); err != nil || res.ID == 0 {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {
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
}

func (uscon UsersController) UpdateUserCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		id := int(claims["userid"].(float64))
		updateUserReq := UserRequestFormat{}
		if err := c.Bind(&updateUserReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		hash := sha256.Sum256([]byte(updateUserReq.Password))
		stringPassword := fmt.Sprintf("%x", hash[:])
		updateUser := entities.User{
			Email: updateUserReq.Email,
			Name:  updateUserReq.Name,
		}
		if updateUserReq.Password != "" {
			updateUser.Password = stringPassword
		}
		if res, err := uscon.Repo.Update(updateUser, id); err != nil || res.ID == 0 {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {
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
}

func (uscon UsersController) DeleteUserCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		id := int(claims["userid"].(float64))
		if res, err := uscon.Repo.Delete(id); err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {
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
}
