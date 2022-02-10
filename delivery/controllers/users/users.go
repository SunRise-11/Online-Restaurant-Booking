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
		if err := c.Bind(&newUserReq); err != nil || newUserReq.Email == "" || newUserReq.Password == "" {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		hash := sha256.Sum256([]byte(newUserReq.Password))
		stringPassword := fmt.Sprintf("%x", hash[:])
		newUser := entities.User{
			Email:       newUserReq.Email,
			Password:    stringPassword,
			Name:        newUserReq.Name,
			PhoneNumber: newUserReq.Phone_Number,
		}
		if res, err := uscon.Repo.Register(newUser); err != nil || res.ID == 0 {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		} else {
			responses := UserResponse{
				ID:         res.ID,
				Name:       res.Name,
				Email:      res.Email,
				Phone:      res.PhoneNumber,
				Balance:    res.Balance,
				Reputation: res.Reputation,
			}
			response := UserResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Data:    responses,
			}

			return c.JSON(http.StatusOK, response)
		}

	}
}

func (uscon UsersController) LoginAuthCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		loginFormat := LoginRequestFormat{}
		if err := c.Bind(&loginFormat); err != nil || loginFormat.Email == "" || loginFormat.Password == "" {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		hash := sha256.Sum256([]byte(loginFormat.Password))
		stringPassword := fmt.Sprintf("%x", hash[:])
		if res, err := uscon.Repo.LoginUser(loginFormat.Email, stringPassword); err != nil || res.ID == 0 {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {
			token, _ := auth.CreateTokenAuthUser(res.ID)

			return c.JSON(http.StatusOK, LoginResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Token:   token,
			})
		}

	}
}

func (uscon UsersController) GetUserCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))
		if res, err := uscon.Repo.Get(uint(userID)); err != nil || res.ID == 0 {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {
			responses := UserResponse{
				ID:         res.ID,
				Name:       res.Name,
				Email:      res.Email,
				Phone:      res.PhoneNumber,
				Balance:    res.Balance,
				Reputation: res.Reputation,
			}

			response := UserResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Data:    responses,
			}

			return c.JSON(http.StatusOK, response)
		}

	}
}

func (uscon UsersController) UpdateUserCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))
		updateUserReq := UserRequestFormat{}
		if err := c.Bind(&updateUserReq); err != nil || updateUserReq.Email == "" || updateUserReq.Password == "" {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		hash := sha256.Sum256([]byte(updateUserReq.Password))
		stringPassword := fmt.Sprintf("%x", hash[:])
		updateUser := entities.User{
			Name:        updateUserReq.Name,
			Email:       updateUserReq.Email,
			Password:    stringPassword,
			PhoneNumber: updateUserReq.Phone_Number,
		}
		if res, err := uscon.Repo.Update(uint(userID), updateUser); err != nil || res.ID == 0 {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {
			responses := UserResponse{
				ID:         res.ID,
				Name:       res.Name,
				Email:      res.Email,
				Balance:    res.Balance,
				Reputation: res.Reputation,
			}
			response := UserResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Data:    responses,
			}

			return c.JSON(http.StatusOK, response)
		}

	}
}

func (uscon UsersController) DeleteUserCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))
		if res, err := uscon.Repo.Delete(uint(userID)); err != nil || res.ID == 0 {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {
			responses := UserResponse{
				ID:    res.ID,
				Name:  res.Name,
				Email: res.Email,
				Phone: res.PhoneNumber,
			}
			response := UserResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Data:    responses,
			}

			return c.JSON(http.StatusOK, response)
		}

	}
}
