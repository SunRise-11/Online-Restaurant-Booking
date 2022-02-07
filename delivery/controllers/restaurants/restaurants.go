package restaurants

import (
	"Restobook/delivery/common"
	"Restobook/delivery/controllers/auth"
	"Restobook/entities"
	"Restobook/repository/restaurants"
	"crypto/sha256"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

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

func (rescon RestaurantsController) GetsWaiting() echo.HandlerFunc {
	return func(c echo.Context) error {

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restoID := claims["restoid"]
		userID := claims["userid"]

		if userID != nil || restoID != nil {
			return c.JSON(http.StatusNotAcceptable, common.NewStatusNotAcceptable())
		} else {
			if res, err := rescon.Repo.GetsWaiting(); err != nil || len(res) == 0 {
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
}

func (rescon RestaurantsController) Approve() echo.HandlerFunc {
	return func(c echo.Context) error {

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restoID := claims["restoid"]
		userID := claims["userid"]

		approveRestaurant := ApproveRestaurantDRequestFormat{}
		if err := c.Bind(&approveRestaurant); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		if userID != nil || restoID != nil {
			return c.JSON(http.StatusNotAcceptable, common.NewStatusNotAcceptable())
		} else {
			if res, err := rescon.Repo.Approve(approveRestaurant.ID, approveRestaurant.Status); err != nil || res.ID == 0 {
				return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
			} else {

				responses := RestaurantDApproveResponse{
					ID:          res.ID,
					Name:        res.Name,
					PhoneNumber: res.PhoneNumber,
					Status:      res.Status,
				}

				response := RestaurantResponseFormat{
					Code:    http.StatusOK,
					Message: "Successful Operation",
					Data:    responses,
				}

				return c.JSON(http.StatusOK, response)
			}
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

func (rescon RestaurantsController) GetsByOpen() echo.HandlerFunc {
	return func(c echo.Context) error {

		date_time := c.QueryParam("date_time")
		fmt.Println("======================================")
		fmt.Println("DATE_TIME", date_time)

		if res, err := rescon.Repo.Gets(); err != nil || len(res) == 0 {
			fmt.Println("======================================")
			fmt.Println("=>ERROR Gets", res)
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {
			fmt.Println("======================================")
			fmt.Println("=>SUCCESS Gets", res)
			date_time_parse, _ := time.Parse("2006-01-02 15:04:05", date_time)
			date_time_split := strings.Split(date_time, " ")

			day := date_time_parse.Weekday().String()
			daytoint := 0
			for i := 0; i < len(common.Daytoint); i++ {
				if day == common.Daytoint[i].Day {
					daytoint = common.Daytoint[i].No
				}
			}
			time := date_time_split[1]
			timeall := strings.Split(time, ":")
			timealls := timeall[0] + timeall[1]
			timeallsInt, _ := strconv.Atoi(timealls)
			fmt.Println("======================================")
			fmt.Println("=>FIND date_time_parse", date_time_parse)
			fmt.Println("=>FIND DayOpen in", day)
			fmt.Println("=>FIND TimeOpen in", timeallsInt)
			fmt.Println("=>FIND dayyoint", daytoint)
			if res, err := rescon.Repo.GetsByOpen(daytoint); err != nil {
				fmt.Println("======================================")
				fmt.Println("=>ERROR GetsByOpen", res)
				return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
			} else {
				newRestaurantD := []entities.RestaurantDetail{}
				for i := 0; i < len(res); i++ {
					fmt.Println("======================================")
					fmt.Println("=>SUCCESS GetsByOpen", res[i])
					splitOH := strings.Split(res[i].Open_Hour, ":")
					allOH, _ := strconv.Atoi(splitOH[0] + splitOH[1])

					splitCH := strings.Split(res[i].Close_Hour, ":")
					allCH, _ := strconv.Atoi(splitCH[0] + splitCH[1])

					if timeallsInt >= allOH && timeallsInt <= allCH {
						date_time_parse_noutc := date_time_split[0] + " " + date_time_split[1]
						fmt.Println("======================================")
						fmt.Println("=>FIND date_time_parse_noutc", date_time_parse_noutc)
						if _, total_seat, err := rescon.Repo.GetExistSeat(res[i].ID, date_time_parse_noutc); err != nil {
							fmt.Println("======================================")
							fmt.Println("=>ERROR GetExistSeat because no transaction with restaurantID", res[i].ID)
							newRestaurantD = append(newRestaurantD, res[i])
							response := RestaurantResponseFormat{
								Code:    http.StatusOK,
								Message: "Successful Operation",
								Data:    newRestaurantD,
							}

							return c.JSON(http.StatusOK, response)
						} else {
							fmt.Println("======================================")
							fmt.Println("=>SUCCESS GetExistSeat", total_seat)
							res[i].Seats = res[i].Seats - total_seat
							newRestaurantD = append(newRestaurantD, res[i])
						}

					}
				}
				response := RestaurantResponseFormat{
					Code:    http.StatusOK,
					Message: "Successful Operation",
					Data:    newRestaurantD,
				}

				return c.JSON(http.StatusOK, response)
			}

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
			Open:           createRestoDReq.Open,
			Close:          createRestoDReq.Close,
			Open_Hour:      createRestoDReq.Open_Hour,
			Close_Hour:     createRestoDReq.Close_Hour,
			Price:          createRestoDReq.Price,
			Latitude:       createRestoDReq.Latitude,
			Longitude:      createRestoDReq.Longitude,
			City:           createRestoDReq.City,
			Address:        createRestoDReq.Address,
			PhoneNumber:    createRestoDReq.PhoneNumber,
			ProfilePicture: createRestoDReq.ProfilePicture,
			Seats:          createRestoDReq.Seats,
			Description:    createRestoDReq.Description,
			Status:         "Waiting for approval",
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
			Open:           updateRestoDReq.Open,
			Close:          updateRestoDReq.Close,
			Open_Hour:      updateRestoDReq.Open_Hour,
			Close_Hour:     updateRestoDReq.Close_Hour,
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
		restoID := claims["restoid"]
		userID := claims["userid"]

		if userID != nil || restoID != nil {
			return c.JSON(http.StatusNotAcceptable, common.NewStatusNotAcceptable())
		} else {
			delRestaurant := DeleteRestauranRequestFormat{}
			if err := c.Bind(&delRestaurant); err != nil {
				return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
			}

			if _, err := rescon.Repo.Delete(uint(delRestaurant.ID)); err != nil {
				return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
			} else {

				return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
			}
		}
	}
}
