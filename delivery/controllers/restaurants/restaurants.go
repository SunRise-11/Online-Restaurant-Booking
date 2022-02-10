package restaurants

import (
	"Restobook/delivery/common"
	"Restobook/delivery/controllers/auth"
	"Restobook/delivery/helpers"
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
		newUserReq := RegisterRequestFormat{}
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
			data := RestaurantResponseFormat{
				ID:    res.ID,
				Email: res.Email,
			}
			response := RestaurantsResponseFormat{
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
		if res, err := rescon.Repo.Login(loginFormat.Email, stringPassword); err != nil || res.Email == "" || res.ID == 0 {
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

func (rescon RestaurantsController) UpdateMyRestoCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restoID := int(claims["restoid"].(float64))

		updateRestoReq := RegisterRequestFormat{}
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
			data := RestaurantResponseFormat{
				ID:    res.ID,
				Email: res.Email,
			}
			response := RestaurantsResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Data:    data,
			}

			return c.JSON(http.StatusOK, response)
		}

	}
}

func (rescon RestaurantsController) GetMyRestoCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restoID := int(claims["restoid"].(float64))

		if res, resD, err := rescon.Repo.Get(uint(restoID)); err != nil || res.ID == 0 {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {

			restaurantDetail := RestaurantDetailResponseFormat{
				ID:             resD.ID,
				Status:         resD.Status,
				ProfilePicture: resD.ProfilePicture,
				Name:           resD.Name,
				Description:    resD.Description,
				Open:           resD.Open,
				Close:          resD.Close,
				Open_Hour:      resD.Open_Hour,
				Close_Hour:     resD.Close_Hour,
				Address:        resD.Address,
				City:           resD.City,
				Latitude:       resD.Latitude,
				Longitude:      resD.Longitude,
				PhoneNumber:    resD.PhoneNumber,
				Seats:          resD.Seats,
				Price:          resD.Price,
			}

			response := RestaurantsResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Data:    restaurantDetail,
			}

			return c.JSON(http.StatusOK, response)
		}

	}
}

func (rescon RestaurantsController) CreateDetailRestoCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restoID := int(claims["restoid"].(float64))

		createRestoDReq := CreateRestaurantDetailRequestFormat{}
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
		}

		if res, err := rescon.Repo.CreateDetail(uint(restoID), createRestoD); err != nil || res.ID == 0 {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {

			restaurantDetail := RestaurantDetailResponseFormat{
				ID:             res.ID,
				Status:         res.Status,
				ProfilePicture: res.ProfilePicture,
				Name:           res.Name,
				Description:    res.Description,
				Open:           res.Open,
				Close:          res.Close,
				Open_Hour:      res.Open_Hour,
				Close_Hour:     res.Close_Hour,
				Address:        res.Address,
				City:           res.City,
				Latitude:       res.Latitude,
				Longitude:      res.Longitude,
				PhoneNumber:    res.PhoneNumber,
				Seats:          res.Seats,
				Price:          res.Price,
			}

			response := RestaurantsResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Data:    restaurantDetail,
			}

			return c.JSON(http.StatusOK, response)
		}

	}
}

func (rescon RestaurantsController) UpdateDetailRestoCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restoID := int(claims["restoid"].(float64))

		updateRestoDReq := UpdateRestaurantDetailRequestFormat{}
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

			restaurantDetail := RestaurantDetailResponseFormat{
				ID:             res.ID,
				Status:         res.Status,
				ProfilePicture: res.ProfilePicture,
				Name:           res.Name,
				Description:    res.Description,
				Open:           res.Open,
				Close:          res.Close,
				Open_Hour:      res.Open_Hour,
				Close_Hour:     res.Close_Hour,
				Address:        res.Address,
				City:           res.City,
				Latitude:       res.Latitude,
				Longitude:      res.Longitude,
				PhoneNumber:    res.PhoneNumber,
				Seats:          res.Seats,
				Price:          res.Price,
			}
			response := RestaurantsResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Data:    restaurantDetail,
			}

			return c.JSON(http.StatusOK, response)
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

				response := RestaurantsResponseFormat{
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

		approveRestaurant := ApproveRequestFormat{}
		if err := c.Bind(&approveRestaurant); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		if userID != nil || restoID != nil {
			return c.JSON(http.StatusNotAcceptable, common.NewStatusNotAcceptable())
		} else {
			if res, err := rescon.Repo.Approve(approveRestaurant.ID, approveRestaurant.Status); err != nil || res.ID == 0 {
				return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
			} else {

				responses := RestaurantResponseFormat{
					ID:          res.ID,
					Name:        res.Name,
					Email:       res.Restaurant[0].Email,
					PhoneNumber: res.PhoneNumber,
					Status:      res.Status,
				}

				response := RestaurantsResponseFormat{
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
		res, err := rescon.Repo.Gets()
		if err != nil || len(res) == 0 {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

		responses := []RestaurantDetailResponseFormat{}

		for _, restaurant := range res {
			score := []int{}
			rating := 0
			values := 0
			for _, value := range restaurant.Rating {
				score = append(score, value.Rating)
			}

			if len(score) < 1 {
				rating = 0
			} else {
				for _, value := range score {
					values += value
				}
				rating = values / len(score)
			}
			responses = append(responses, RestaurantDetailResponseFormat{
				ID:             restaurant.ID,
				Name:           restaurant.Name,
				Open:           restaurant.Open,
				Close:          restaurant.Close,
				Open_Hour:      restaurant.Open_Hour,
				Close_Hour:     restaurant.Close_Hour,
				Price:          restaurant.Price,
				Latitude:       restaurant.Latitude,
				Longitude:      restaurant.Longitude,
				City:           restaurant.City,
				Address:        restaurant.Address,
				PhoneNumber:    restaurant.PhoneNumber,
				ProfilePicture: restaurant.ProfilePicture,
				Seats:          restaurant.Seats,
				Description:    restaurant.Description,
				Status:         restaurant.Status,
				Rating:         rating,
			})
		}
		response := RestaurantsResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    responses,
		}
		return c.JSON(http.StatusOK, response)

	}
}

func (rescon RestaurantsController) GetsByOpen() echo.HandlerFunc {
	return func(c echo.Context) error {

		date_time := c.QueryParam("date_time")
		// fmt.Println("======================================")
		// fmt.Println("DATE_TIME", date_time)

		if res, err := rescon.Repo.Gets(); err != nil || len(res) == 0 {
			// fmt.Println("======================================")
			// fmt.Println("=>ERROR Gets", err, "<=>", res)
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {
			// fmt.Println("======================================")
			// fmt.Println("=>SUCCESS Gets", res)
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
			// fmt.Println("======================================")
			// fmt.Println("=>FIND date_time_parse", date_time_parse)
			// fmt.Println("=>FIND DayOpen in", day)
			// fmt.Println("=>FIND TimeOpen in", timeallsInt)
			// fmt.Println("=>FIND dayyoint", daytoint)
			if res, err := rescon.Repo.GetsByOpen(daytoint); err != nil {
				// fmt.Println("======================================")
				// fmt.Println("=>ERROR GetsByOpen", err, "<=>", res)
				return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
			} else {
				newRestaurantD := []entities.RestaurantDetail{}
				for i := 0; i < len(res); i++ {
					// fmt.Println("======================================")
					// fmt.Println("=>SUCCESS GetsByOpen", res[i])
					splitOH := strings.Split(res[i].Open_Hour, ":")
					allOH, _ := strconv.Atoi(splitOH[0] + splitOH[1])

					splitCH := strings.Split(res[i].Close_Hour, ":")
					allCH, _ := strconv.Atoi(splitCH[0] + splitCH[1])

					if timeallsInt >= allOH && timeallsInt <= allCH {
						date_time_parse_noutc := date_time_split[0] + " " + date_time_split[1]
						// fmt.Println("======================================")
						// fmt.Println("=>FIND date_time_parse_noutc", date_time_parse_noutc)
						if _, total_seat, err := rescon.Repo.GetExistSeat(res[i].ID, date_time_parse_noutc); err != nil {
							// fmt.Println("======================================")
							// fmt.Println("=>ERROR GetExistSeat because no transaction with restaurantID", res[i].ID, "<=>", err, "<=>", total_seat)
							newRestaurantD = append(newRestaurantD, res[i])
							response := RestaurantsResponseFormat{
								Code:    http.StatusOK,
								Message: "Successful Operation",
								Data:    newRestaurantD,
							}
							return c.JSON(http.StatusOK, response)
						} else {
							// fmt.Println("======================================")
							// fmt.Println("=>SUCCESS GetExistSeat", total_seat)
							res[i].Seats = res[i].Seats - total_seat
							newRestaurantD = append(newRestaurantD, res[i])
						}

					}
				}
				response := RestaurantsResponseFormat{
					Code:    http.StatusOK,
					Message: "Successful Operation",
					Data:    newRestaurantD,
				}

				return c.JSON(http.StatusOK, response)
			}

		}
	}
}

func (rescon RestaurantsController) DeleteRestoCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restoID := claims["restoid"]
		userID := claims["userid"]

		if userID != nil || restoID != nil {
			return c.JSON(http.StatusNotAcceptable, common.NewStatusNotAcceptable())
		} else {
			delRestaurant := DeleteRequestFormat{}
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

func (rescon RestaurantsController) ExportPDF() echo.HandlerFunc {
	return func(c echo.Context) error {

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restoID := int(claims["restoid"].(float64))

		date_time := c.QueryParam("date_time")
		date_time_parse, _ := time.Parse("2006-01-02 15:04:05", date_time)
		date_time_split := strings.Split(date_time_parse.String(), " ")

		fmt.Println("date_time_split", date_time_split)

		if res, err := rescon.Repo.Export(uint(restoID), date_time_split[0]); err != nil || len(res) == 0 {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		} else {

			var successOrder, successSeat, failOrder, failSeat, cancelOrder, cancelSeat, rejectedOrder, rejectedSeat int
			var successTotal, failTotal, cancelTotal int

			for i := 0; i < len(res); i++ {
				if res[i].Status == "Success" {
					successOrder += 1
					successSeat += res[i].Persons
					successTotal += res[i].Total
				} else if res[i].Status == "Fail" {
					failOrder += 1
					failSeat += res[i].Persons
					failTotal += res[i].Total
				} else if res[i].Status == "Cancel" {
					cancelOrder += 1
					cancelSeat += res[i].Persons
					cancelTotal += 20000
				} else if res[i].Status == "Rejected" {
					rejectedOrder += 1
					rejectedSeat += res[i].Persons
				}
			}

			totalOrder := successOrder + failOrder + cancelOrder
			totalSeat := successSeat + failSeat + cancelSeat
			grandTotal := successTotal + failTotal + cancelTotal

			resOrder := ExportPDF_Order_Response{
				Number_of_success_orders:  fmt.Sprintf("%v Orders", successOrder),
				Number_of_fail_orders:     fmt.Sprintf("%v Orders", failOrder),
				Number_of_cancel_orders:   fmt.Sprintf("%v Orders", cancelOrder),
				Total_orders:              fmt.Sprintf("%v Orders", totalOrder),
				Number_of_rejected_orders: fmt.Sprintf("%v Orders", rejectedOrder),
			}

			resSeats := ExportPDF_Seats_Response{
				Number_of_success_seats:  fmt.Sprintf("%v seats", successSeat),
				Number_of_fail_seats:     fmt.Sprintf("%v seats", failSeat),
				Number_of_cancel_seats:   fmt.Sprintf("%v seats", cancelSeat),
				Total_seats:              fmt.Sprintf("%v seats", totalSeat),
				Number_of_rejected_seats: fmt.Sprintf("%v seats", rejectedSeat),
			}

			resTotal := ExportPDF_Total_Response{
				Number_of_success_total: fmt.Sprintf("Rp.%v", successTotal),
				Number_of_fail_total:    fmt.Sprintf("Rp.%v", failTotal),
				Number_of_cancel_total:  fmt.Sprintf("Rp.%v", cancelTotal),
				Total:                   fmt.Sprintf("Rp.%v", grandTotal),
			}

			responses := ExportPDFResponseFormat{
				Date:    date_time_split[0],
				Name:    res[0].Restaurant.RestaurantDetail.Name,
				Address: res[0].Restaurant.RestaurantDetail.Address,
				Orders:  resOrder,
				Seats:   resSeats,
				Total:   resTotal,
			}

			response := RestaurantsResponseFormat{
				Code:    200,
				Message: "Successful Operation",
				Data:    responses,
			}

			helpers.CreatePDFReport(res[0].Restaurant.RestaurantDetail.Name, res[0].Restaurant.RestaurantDetail.Address, date_time_split[0])

			return c.JSON(http.StatusOK, response)
		}

	}
}
