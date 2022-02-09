package auth

import (
	"Restobook/delivery/common"
	"Restobook/entities"
	"Restobook/repository/users"
	"crypto/sha256"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateTokenAuthUser(id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userid"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(common.JWT_SECRET_KEY))
}

func CreateTokenAuthRestaurant(id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["restoid"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(common.JWT_SECRET_KEY))
}

func CreateTokenAuthAdmin(id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["admin"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(common.JWT_SECRET_KEY))
}

type AdminController struct {
	Repo users.UsersInterface
}

func NewAdminControllers(usrep users.UsersInterface) *AdminController {
	return &AdminController{Repo: usrep}
}

func (admcon AdminController) RegisterAdminCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		newAdminReq := RegisterRequestFormat{}

		if err := c.Bind(&newAdminReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		hash := sha256.Sum256([]byte(newAdminReq.Password))
		stringPassword := fmt.Sprintf("%x", hash[:])
		newUser := entities.User{
			Name:     newAdminReq.Name,
			Email:    newAdminReq.Email,
			Password: stringPassword,
		}
		if res, err := admcon.Repo.RegisterAdmin(newUser); err != nil || res.ID == 0 {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		} else {
			data := AdminResponse{
				ID:    res.ID,
				Name:  res.Name,
				Email: res.Email,
			}
			response := AdminResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Data:    data,
			}

			return c.JSON(http.StatusOK, response)
		}
	}

}

func (admcon AdminController) LoginAdminCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		loginFormat := LoginRequestFormat{}
		if err := c.Bind(&loginFormat); err != nil || loginFormat.Email == "" || loginFormat.Password == "" {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		hash := sha256.Sum256([]byte(loginFormat.Password))
		stringPassword := fmt.Sprintf("%x", hash[:])
		if res, err := admcon.Repo.LoginUser(loginFormat.Email, stringPassword); err != nil || res.ID == 0 {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {
			token, _ := CreateTokenAuthAdmin(res.ID)

			return c.JSON(http.StatusOK, LoginResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Token:   token,
			})
		}
	}

}
