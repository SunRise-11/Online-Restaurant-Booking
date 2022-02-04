package rating

import (
	"Restobook/delivery/common"
	"Restobook/entities"
	"Restobook/repository/rating"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type RatingController struct {
	Repo rating.RatingInterface
}

func NewRatingController(ratrep rating.RatingInterface) *RatingController {
	return &RatingController{Repo: ratrep}
}

func (rc RatingController) Create(c echo.Context) error {
	var ratingRequest PostRatingRequest

	if err := c.Bind(&ratingRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	uid := c.Get("user").(*jwt.Token)
	claims := uid.Claims.(jwt.MapClaims)
	userID := int(claims["userid"].(float64))

	isCanGiveRating, _ := rc.Repo.IsCanGiveRating(userID, ratingRequest.RestaurantID)
	if !isCanGiveRating {
		return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
	}

	data := entities.Rating{
		RestaurantID: uint(ratingRequest.RestaurantID),
		UserID:       uint(userID),
		Rating:       ratingRequest.Rating,
		Comment:      ratingRequest.Comment,
	}

	ratingData, err := rc.Repo.Create(data)
	if err != nil {
		ratingData, err = rc.Repo.Update(data)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
	}

	response := RatingResponse{
		RestaurantID: int(ratingData.RestaurantID),
		UserID:       int(ratingData.UserID),
		Username:     ratingData.User.Name,
		Rating:       ratingData.Rating,
		Comment:      ratingData.Comment,
	}

	return c.JSON(http.StatusOK, response)
}

func (rc RatingController) Update(c echo.Context) error {
	var ratingRequest UpdateRatingRequest

	if err := c.Bind(&ratingRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	restaurantId, err := strconv.Atoi(c.Param("restaurantId"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	uid := c.Get("user").(*jwt.Token)
	claims := uid.Claims.(jwt.MapClaims)
	userID := int(claims["userid"].(float64))

	data := entities.Rating{
		RestaurantID: uint(restaurantId),
		UserID:       uint(userID),
		Rating:       ratingRequest.Rating,
		Comment:      ratingRequest.Comment,
	}

	ratingData, err := rc.Repo.Update(data)
	if err != nil {
		return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
	}

	response := RatingResponse{
		RestaurantID: int(ratingData.RestaurantID),
		UserID:       int(ratingData.UserID),
		Username:     ratingData.User.Name,
		Rating:       ratingData.Rating,
		Comment:      ratingData.Comment,
	}

	return c.JSON(http.StatusOK, response)
}

func (rc RatingController) Delete(c echo.Context) error {

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
