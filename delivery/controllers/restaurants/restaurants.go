package restaurants

import (
	"Restobook/delivery/common"
	"Restobook/delivery/controllers/auth"
	"Restobook/entities"
	"Restobook/repository/restaurants"
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RestaurantsController struct {
	Repo restaurants.RestaurantsInterface
}

func NewRestaurantsControllers(resrep restaurants.RestaurantsInterface) *RestaurantsController {
	return &RestaurantsController{Repo: resrep}
}

func (rescon RestaurantsController) RegisterUserCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		newUserReq := RestaurantRequestFormat{}
		if err := c.Bind(&newUserReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		hash := sha256.Sum256([]byte(newUserReq.Password))
		stringPassword := fmt.Sprintf("%x", hash[:])
		newResto := entities.Restaurant{
			Email:    newUserReq.Email,
			Password: stringPassword,
		}
		if res, err := rescon.Repo.Register(newResto); err != nil || res.ID == 0 {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		} else {
			data := RestaurantResponse{
				ID:    res.ID,
				Email: res.Email,
			}
			response := RestaurantResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Data:    data,
			}

			return c.JSON(http.StatusOK, response)
		}

	}
}

func (rescon RestaurantsController) LoginAuthCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		loginFormat := LoginRequestFormat{}
		if err := c.Bind(&loginFormat); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		hash := sha256.Sum256([]byte(loginFormat.Password))
		stringPassword := fmt.Sprintf("%x", hash[:])
		if res, err := rescon.Repo.LoginRestaurant(loginFormat.Email, stringPassword); err != nil || res.Email == "" || res.ID == 0 {
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
