package ratings

import (
	"Restobook/delivery/common"
	"Restobook/entities"
	"Restobook/repository/ratings"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type RatingsController struct {
	Repo ratings.RatingsInterface
}

func NewRatingController(ratrep ratings.RatingsInterface) *RatingsController {
	return &RatingsController{Repo: ratrep}
}

func (rc RatingsController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var ratingRequest PostRatingRequest

		if err := c.Bind(&ratingRequest); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))

		isCanGiveRating, _ := rc.Repo.IsCanGiveRating(userID, ratingRequest.RestaurantDetailID)
		if !isCanGiveRating {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

		data := entities.Rating{
			RestaurantDetailID: uint(ratingRequest.RestaurantDetailID),
			UserID:             uint(userID),
			Rating:             ratingRequest.Rating,
			Comment:            ratingRequest.Comment,
		}

		ratingData, err := rc.Repo.Create(data)
		if err != nil {
			ratingData, err = rc.Repo.Update(data)
		}

		response := RatingResponse{
			RestaurantDetailID: int(ratingData.RestaurantDetailID),
			UserID:             int(ratingData.UserID),
			Username:           ratingData.User.Name,
			Rating:             ratingData.Rating,
			Comment:            ratingData.Comment,
		}

		finalResponse := RatingResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    response,
		}
		return c.JSON(http.StatusOK, finalResponse)
	}
}

func (rc RatingsController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var ratingRequest UpdateRatingRequest

		if err := c.Bind(&ratingRequest); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		restaurantDetailId, err := strconv.Atoi(c.Param("restaurantId"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))

		data := entities.Rating{
			RestaurantDetailID: uint(restaurantDetailId),
			UserID:             uint(userID),
			Rating:             ratingRequest.Rating,
			Comment:            ratingRequest.Comment,
		}

		ratingData, err := rc.Repo.Update(data)
		if err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

		response := RatingResponse{
			RestaurantDetailID: int(ratingData.RestaurantDetailID),
			UserID:             int(ratingData.UserID),
			Username:           ratingData.User.Name,
			Rating:             ratingData.Rating,
			Comment:            ratingData.Comment,
		}

		finalResponse := RatingResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    response,
		}
		return c.JSON(http.StatusOK, finalResponse)
	}
}

func (rc RatingsController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {

		restaurantId, err := strconv.Atoi(c.Param("restaurantId"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))

		_, err = rc.Repo.Delete(userID, restaurantId)
		if err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}
}
