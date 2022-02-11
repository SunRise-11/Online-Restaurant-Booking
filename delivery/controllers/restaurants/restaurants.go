package restaurants

import (
	"Restobook/delivery/common"
	"Restobook/delivery/controllers/auth"
	"Restobook/delivery/helpers"
	"Restobook/entities"
	"Restobook/repository/restaurants"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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

		file, _ := c.FormFile("profile_picture")

		var imgurlink Imgurlink

		if src, err := file.Open(); err != nil {
			fmt.Println("===> ERROR FILE OPEN", err)

		} else {

			defer src.Close()
			if dst, err := os.Create(fmt.Sprintf("./IMAGES/Restaurants/%v", file.Filename)); err != nil {
				fmt.Println("===> ERROR FILE CREATE")

			} else {

				defer dst.Close()
				if _, err := io.Copy(dst, src); err != nil {
					fmt.Println("===> ERROR FILE COPY")

				} else {

					filebytes, _ := ioutil.ReadFile(fmt.Sprintf("./IMAGES/Restaurants/%v", file.Filename))
					a := helpers.ImgurUpload(filebytes)
					json.Unmarshal(a, &imgurlink)
				}
			}
		}

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
			ProfilePicture: imgurlink.Data.Link,
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
			score := []float64{}
			var rating float64
			var values float64
			for _, value := range restaurant.Rating {
				score = append(score, float64(value.Rating))
			}

			if len(score) < 1 {
				rating = 0
			} else {
				for _, value := range score {
					values += value
				}
				rating = values / float64(len(score))
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

		if res, err := rescon.Repo.Gets(); err != nil || len(res) == 0 {

			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())

		} else {

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

			if res, err := rescon.Repo.GetsByOpen(daytoint); err != nil {
				return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
			} else {

				newRestaurantD := []RestaurantDetailResponseFormat{}

				for i := 0; i < len(res); i++ {

					splitOpenHour := strings.Split(res[i].Open_Hour, ":")
					allOpenHour, _ := strconv.Atoi(splitOpenHour[0] + splitOpenHour[1])

					splitCloseHour := strings.Split(res[i].Close_Hour, ":")
					allCloseHour, _ := strconv.Atoi(splitCloseHour[0] + splitCloseHour[1])

					if timeallsInt >= allOpenHour && timeallsInt <= allCloseHour {

						date_time_parse_noutc := date_time_split[0] + " " + date_time_split[1]

						score := []float64{}
						var rating float64
						var values float64

						for a := 0; a < len(res[i].Rating); a++ {
							score = append(score, float64(res[i].Rating[a].Rating))
						}
						if len(score) < 1 {
							rating = 0
						} else {
							for _, value := range score {
								values += value
							}
							rating = values / float64(len(score))
						}

						if _, total_seat, err := rescon.Repo.GetExistSeat(res[i].ID, date_time_parse_noutc); err != nil {

							newRestaurantD = append(newRestaurantD, RestaurantDetailResponseFormat{
								ID:             res[i].ID,
								Status:         res[i].Status,
								ProfilePicture: res[i].ProfilePicture,
								Name:           res[i].Name,
								Description:    res[i].Description,
								Rating:         rating,
								Open:           res[i].Open,
								Close:          res[i].Close,
								Open_Hour:      res[i].Open_Hour,
								Close_Hour:     res[i].Close_Hour,
								Address:        res[i].Address,
								City:           res[i].City,
								PhoneNumber:    res[i].PhoneNumber,
								Latitude:       res[i].Latitude,
								Longitude:      res[i].Longitude,
								Seats:          res[i].Seats,
								Price:          res[i].Price,
							})

							response := RestaurantsResponseFormat{
								Code:    http.StatusOK,
								Message: "Successful Operation",
								Data:    newRestaurantD,
							}

							return c.JSON(http.StatusOK, response)

						} else {

							res[i].Seats = res[i].Seats - total_seat
							newRestaurantD = append(newRestaurantD, RestaurantDetailResponseFormat{
								ID:             res[i].ID,
								Status:         res[i].Status,
								ProfilePicture: res[i].ProfilePicture,
								Name:           res[i].Name,
								Description:    res[i].Description,
								Rating:         rating,
								Open:           res[i].Open,
								Close:          res[i].Close,
								Open_Hour:      res[i].Open_Hour,
								Close_Hour:     res[i].Close_Hour,
								Address:        res[i].Address,
								City:           res[i].City,
								PhoneNumber:    res[i].PhoneNumber,
								Latitude:       res[i].Latitude,
								Longitude:      res[i].Longitude,
								Seats:          res[i].Seats,
								Price:          res[i].Price,
							})
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

		day := c.QueryParam("day")
		month := c.QueryParam("month")
		year := c.QueryParam("year")

		finalday := ""
		if day != "" && month != "" && year != "" {
			finalday = fmt.Sprintf("%v-%v-%v", year, month, day)
		} else if day == "" && month != "" && year != "" {
			finalday = fmt.Sprintf("%v-%v", year, month)
		} else if day == "" && month == "" && year != "" {
			finalday = fmt.Sprintf("%v", year)
		}

		if res, err := rescon.Repo.Export(uint(restoID), finalday); err != nil || len(res) == 0 {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		} else {
			var successOrder, successSeat,
				failOrder, failSeat,
				cancelOrder, cancelSeat,
				rejectedOrder, rejectedSeat,
				successTotal, failTotal, cancelTotal int

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
				Date:    finalday,
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

			helpers.CreatePDFReport(
				res[0].Restaurant.RestaurantDetail.Name,
				res[0].Restaurant.RestaurantDetail.Address,
				finalday,
				[]int{successOrder, successSeat, successTotal},
				[]int{failOrder, failSeat, failTotal},
				[]int{cancelOrder, cancelSeat, cancelTotal},
				[]int{totalOrder, totalSeat, grandTotal},
				[]int{rejectedOrder, rejectedSeat})

			return c.JSON(http.StatusOK, response)
		}
	}
}

func (rescon RestaurantsController) ImgurCallBack() echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("c", c)
		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}
}
