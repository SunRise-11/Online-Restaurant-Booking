package restaurants

import (
	"Restobook/delivery/common"
	"Restobook/delivery/controllers/auth"
	"Restobook/entities"
	"Restobook/repository/restaurants"
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type RestaurantsController struct {
	Repo restaurants.RestaurantsInterface
}

func NewRestaurantsControllers(resrep restaurants.RestaurantsInterface) *RestaurantsController {
	return &RestaurantsController{Repo: resrep}
}

func (rescon RestaurantsController) RegisterRestoCtrl() echo.HandlerFunc {

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

func (rescon RestaurantsController) LoginRestoCtrl() echo.HandlerFunc {
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
			token, _ := auth.CreateTokenAuthRestaurant(res.ID)

			return c.JSON(http.StatusOK, LoginResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Token:   token,
			})
		}

	}
}

func (rescon RestaurantsController) Gets() echo.HandlerFunc {
	return func(c echo.Context) error {

		if res, err := rescon.Repo.Gets(); err != nil || len(res) == 0 {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {

			response := RestaurantResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Data:    res,
			}

			return c.JSON(http.StatusOK, response)
		}

	}
}

func (rescon RestaurantsController) GetRestoByIdCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restoID := int(claims["restoid"].(float64))

		if res, resD, err := rescon.Repo.Get(uint(restoID)); err != nil || res.ID == 0 {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {
			data := RestaurantResponse{
				ID:    res.ID,
				Email: res.Email,
				Data:  resD,
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

func (rescon RestaurantsController) UpdateRestoByIdCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restoID := int(claims["restoid"].(float64))

		updateRestoReq := RestaurantRequestFormat{}
		if err := c.Bind(&updateRestoReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		hash := sha256.Sum256([]byte(updateRestoReq.Password))
		stringPassword := fmt.Sprintf("%x", hash[:])
		updateResto := entities.Restaurant{
			Email:    updateRestoReq.Email,
			Password: stringPassword,
		}

		if res, err := rescon.Repo.Update(uint(restoID), updateResto); err != nil || res.ID == 0 {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
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

func (rescon RestaurantsController) CreateDetailRestoByIdCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restoID := int(claims["restoid"].(float64))

		createRestoDReq := CreateRestaurantDRequestFormat{}
		if err := c.Bind(&createRestoDReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		createRestoD := entities.RestaurantDetail{
			Name:           createRestoDReq.Name,
			Open:           createRestoDReq.Open.String(),
			Close:          createRestoDReq.Close.String(),
			Price:          createRestoDReq.Price,
			Latitude:       createRestoDReq.Latitude,
			Longitude:      createRestoDReq.Longitude,
			City:           createRestoDReq.City,
			Address:        createRestoDReq.Address,
			PhoneNumber:    createRestoDReq.PhoneNumber,
			ProfilePicture: createRestoDReq.ProfilePicture,
			Seats:          createRestoDReq.Seats,
			Description:    createRestoDReq.Description,
		}

		if res, err := rescon.Repo.UpdateDetail(uint(restoID), createRestoD); err != nil || res.ID == 0 {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {
			response := RestaurantResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Data:    res,
			}

			return c.JSON(http.StatusOK, response)
		}

	}
}

func (rescon RestaurantsController) UpdateDetailRestoByIdCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restoID := int(claims["restoid"].(float64))

		updateRestoDReq := UpdateRestaurantDRequestFormat{}
		if err := c.Bind(&updateRestoDReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		updateRestoD := entities.RestaurantDetail{
			Open:           updateRestoDReq.Open.String(),
			Close:          updateRestoDReq.Close.String(),
			Price:          updateRestoDReq.Price,
			PhoneNumber:    updateRestoDReq.PhoneNumber,
			ProfilePicture: updateRestoDReq.ProfilePicture,
			Seats:          updateRestoDReq.Seats,
			Description:    updateRestoDReq.Description,
		}

		if res, err := rescon.Repo.UpdateDetail(uint(restoID), updateRestoD); err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {
			response := RestaurantResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Data:    res,
			}

			return c.JSON(http.StatusOK, response)
		}

	}
}

func (rescon RestaurantsController) DeleteRestaurantCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		adminID := int(claims["admin"].(float64))

		if adminID != 1 {
			return c.JSON(http.StatusNotAcceptable, common.NewStatusNotAcceptable())
		} else {
			delRestaurant := DeleteRestauranRequestFormat{}
			if err := c.Bind(&delRestaurant); err != nil {
				return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
			}

			if res, err := rescon.Repo.Delete(uint(delRestaurant.ID)); err != nil {
				return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
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
}
