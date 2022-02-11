package transactions

import (
	"Restobook/delivery/common"
	"Restobook/entities"
	"Restobook/repository/transactions"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type TransactionsController struct {
	Repo transactions.TransactionsInterface
}

func NewTransactionsControllers(transrep transactions.TransactionsInterface) *TransactionsController {
	return &TransactionsController{Repo: transrep}
}

func (transcon TransactionsController) CreateTransactionCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		var balance, day, total int
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))
		newTransactionReq := TransactionRequestFormat{}
		if err := c.Bind(&newTransactionReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		loc, _ := time.LoadLocation("Asia/Singapore")
		var dateTime, _ = time.ParseInLocation("2006-01-02 15:04", newTransactionReq.DateTime, loc)
		for i := 0; i < len(common.Daytoint); i++ {
			if dateTime.Weekday().String() == common.Daytoint[i].Day {
				day = common.Daytoint[i].No
			}
		}
		balanceUser, err := transcon.Repo.GetBalance(uint(userID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		restoDetail, err := transcon.Repo.GetRestoDetail(newTransactionReq.RestaurantID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		openHour := restoDetail.Open_Hour
		closeHour := restoDetail.Close_Hour
		if !strings.Contains(restoDetail.Open, fmt.Sprint(day)) || restoDetail.Status != "OPEN" {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "This Restaurant is Not Open Today",
			})
		}
		minHoursFormat := newTransactionReq.DateTime[0:11] + openHour[:]
		maxHoursFormat := newTransactionReq.DateTime[0:11] + closeHour[:]
		minHours, _ := time.ParseInLocation("2006-01-02 15:04", minHoursFormat, loc)
		maxHours, _ := time.ParseInLocation("2006-01-02 15:04", maxHoursFormat, loc)
		if dateTime.Before(minHours) {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "Sorry This Restaurant is Not Open Yet",
			})
		}
		if dateTime.After(maxHours) || dateTime.Equal(maxHours) {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "Sorry This Restaurant already closed",
			})
		}

		total = newTransactionReq.Persons * restoDetail.Price
		balance = balanceUser.Balance - total
		if balance < 0 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "Your Money is Not Enough For Booking This Restaurant",
			})
		}

		seat, _ := transcon.Repo.GetTotalSeat(newTransactionReq.RestaurantID, newTransactionReq.DateTime)
		isExist, _ := transcon.Repo.CheckSameHour(newTransactionReq.RestaurantID, uint(userID), newTransactionReq.DateTime)
		if isExist {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "You Already Booked at This Hour",
			})
		}
		seatAvailable := restoDetail.Seats - seat
		if newTransactionReq.Persons > seatAvailable {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "Just " + fmt.Sprint(seatAvailable) + " Seats Available at This Hour",
			})
		}
		if _, err := transcon.Repo.UpdateUserBalance(uint(userID), balance); err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		newTransaction := entities.Transaction{
			UserID:       uint(userID),
			RestaurantID: newTransactionReq.RestaurantID,
			DateTime:     dateTime,
			Persons:      newTransactionReq.Persons,
			Total:        total,
		}
		res, err := transcon.Repo.Create(newTransaction)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		data := TransactionResponse{
			ID:           res.ID,
			UserID:       res.UserID,
			RestaurantID: res.RestaurantID,
			DateTime:     res.DateTime,
			Person:       res.Persons,
			Total:        res.Total,
		}
		response := TransactionResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    data,
		}
		return c.JSON(http.StatusOK, response)

	}
}

func (transcon TransactionsController) GetAllWaitingCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))
		transactions, err := transcon.Repo.GetAllWaiting(uint(userID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		data := []TransactionResponse{}
		for _, transaction := range transactions {
			data = append(
				data, TransactionResponse{
					ID:           transaction.ID,
					UserID:       transaction.UserID,
					RestaurantID: transaction.RestaurantID,
					Person:       transaction.Persons,
					DateTime:     transaction.DateTime,
					Total:        transaction.Total,
				},
			)
		}
		response := TransactionResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    data,
		}
		return c.JSON(http.StatusOK, response)

	}
}
func (transcon TransactionsController) GetAllWaitingForRestoCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restoID := int(claims["restoid"].(float64))
		transactions, err := transcon.Repo.GetAllWaitingForResto(uint(restoID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		data := []TransactionResponse{}
		for _, transaction := range transactions {
			data = append(
				data, TransactionResponse{
					ID:           transaction.ID,
					UserID:       transaction.UserID,
					RestaurantID: transaction.RestaurantID,
					Person:       transaction.Persons,
					DateTime:     transaction.DateTime,
					Total:        transaction.Total,
				},
			)
		}
		response := TransactionResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    data,
		}
		return c.JSON(http.StatusOK, response)

	}
}
func (transcon TransactionsController) GetAllAcceptedForRestoCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restoID := int(claims["restoid"].(float64))
		transactions, err := transcon.Repo.GetAllAcceptedForResto(uint(restoID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		data := []TransactionResponse{}
		for _, transaction := range transactions {
			data = append(
				data, TransactionResponse{
					ID:           transaction.ID,
					UserID:       transaction.UserID,
					RestaurantID: transaction.RestaurantID,
					Person:       transaction.Persons,
					DateTime:     transaction.DateTime,
					Total:        transaction.Total,
				},
			)
		}
		response := TransactionResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    data,
		}
		return c.JSON(http.StatusOK, response)

	}
}
func (transcon TransactionsController) GetHistoryCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))
		transactions, err := transcon.Repo.GetHistory(uint(userID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		data := []TransactionHistoryResponse{}
		for _, transaction := range transactions {
			data = append(
				data, TransactionHistoryResponse{
					ID:           transaction.ID,
					UserID:       transaction.UserID,
					RestaurantID: transaction.RestaurantID,
					Person:       transaction.Persons,
					DateTime:     transaction.DateTime,
					Total:        transaction.Total,
					Status:       transaction.Status,
				},
			)
		}
		response := TransactionResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    data,
		}
		return c.JSON(http.StatusOK, response)

	}
}
func (transcon TransactionsController) GetAllAcceptedCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))
		transactions, err := transcon.Repo.GetAllAppointed(uint(userID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		data := []TransactionResponse{}
		for _, transaction := range transactions {
			data = append(
				data, TransactionResponse{
					ID:           transaction.ID,
					UserID:       transaction.UserID,
					RestaurantID: transaction.RestaurantID,
					Person:       transaction.Persons,
					DateTime:     transaction.DateTime,
					Total:        transaction.Total,
				},
			)
		}
		response := TransactionResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    data,
		}
		return c.JSON(http.StatusOK, response)

	}
}

func (transcon TransactionsController) AcceptTransactionCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restoID := int(claims["restoid"].(float64))
		newTransactionReq := TransactionRequestFormat{}
		if err := c.Bind(&newTransactionReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		newTransaction := entities.Transaction{
			ID:     newTransactionReq.ID,
			Status: newTransactionReq.Status,
		}
		_, err := transcon.Repo.GetTransactionUserByStatus(newTransactionReq.ID, uint(restoID), "waiting for confirmation")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		res, err := transcon.Repo.UpdateTransactionStatus(newTransaction)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		data := TransactionResponse{
			ID:           res.ID,
			UserID:       res.UserID,
			RestaurantID: res.RestaurantID,
			DateTime:     res.DateTime,
			Person:       res.Persons,
			Total:        res.Total,
		}
		response := TransactionResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    data,
		}
		return c.JSON(http.StatusOK, response)
	}

}
func (transcon TransactionsController) RejectTransactionCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newBalance int
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restoID := int(claims["restoid"].(float64))
		newTransactionReq := TransactionRequestFormat{}
		if err := c.Bind(&newTransactionReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		res, err := transcon.Repo.GetTransactionUserByStatus(newTransactionReq.ID, uint(restoID), "waiting for confirmation")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())

		}

		newBalance = res.User.Balance + res.Total

		if _, err := transcon.Repo.UpdateUserBalance(res.UserID, newBalance); err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		newTransaction := entities.Transaction{
			ID:     newTransactionReq.ID,
			Status: newTransactionReq.Status,
		}

		if res, err := transcon.Repo.UpdateTransactionStatus(newTransaction); err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		} else {
			data := TransactionResponse{
				ID:           res.ID,
				UserID:       res.UserID,
				RestaurantID: res.RestaurantID,
				DateTime:     res.DateTime,
				Person:       res.Persons,
				Total:        res.Total,
			}
			response := TransactionResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Data:    data,
			}
			return c.JSON(http.StatusOK, response)
		}
	}
}
func (transcon TransactionsController) SuccessTransactionCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restoID := int(claims["restoid"].(float64))
		newTransactionReq := TransactionRequestFormat{}
		if err := c.Bind(&newTransactionReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		transaction, err := transcon.Repo.GetTransactionUserByStatus(newTransactionReq.ID, uint(restoID), "Accepted")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		newTransaction := entities.Transaction{
			ID:     newTransactionReq.ID,
			Status: newTransactionReq.Status,
		}
		totalReputation := transaction.User.Reputation + 5
		if totalReputation > 100 {
			totalReputation = 100
		}

		if _, err := transcon.Repo.UpdateUserReputation(newTransactionReq.UserID, totalReputation); err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		res, err := transcon.Repo.UpdateTransactionStatus(newTransaction)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		data := TransactionResponse{
			ID:           res.ID,
			UserID:       res.UserID,
			RestaurantID: res.RestaurantID,
			DateTime:     res.DateTime,
			Person:       res.Persons,
			Total:        res.Total,
		}
		response := TransactionResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    data,
		}
		return c.JSON(http.StatusOK, response)

	}
}
func (transcon TransactionsController) FailTransactionCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restoID := int(claims["restoid"].(float64))
		newTransactionReq := TransactionRequestFormat{}
		if err := c.Bind(&newTransactionReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		transaction, err := transcon.Repo.GetTransactionUserByStatus(newTransactionReq.ID, uint(restoID), "Accepted")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		newTransaction := entities.Transaction{
			ID:     newTransactionReq.ID,
			Status: newTransactionReq.Status,
		}
		totalReputation := transaction.User.Reputation - 3
		if totalReputation < 0 {
			totalReputation = 0
		}

		if _, err := transcon.Repo.UpdateUserReputation(newTransactionReq.UserID, totalReputation); err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		res, err := transcon.Repo.UpdateTransactionStatus(newTransaction)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		data := TransactionResponse{
			ID:           res.ID,
			UserID:       res.UserID,
			RestaurantID: res.RestaurantID,
			DateTime:     res.DateTime,
			Person:       res.Persons,
			Total:        res.Total,
		}
		response := TransactionResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    data,
		}
		return c.JSON(http.StatusOK, response)

	}
}
func (transcon TransactionsController) CancelTransactionCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newBalance int
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userId := int(claims["userid"].(float64))
		newTransactionReq := TransactionRequestFormat{}
		if err := c.Bind(&newTransactionReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		transaction, err := transcon.Repo.GetTransactionById(newTransactionReq.ID, uint(userId))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		if !(transaction.Status == "Accepted" || transaction.Status == "waiting for confirmation") {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "You Cant Cancel This Transaction at This Time",
			})
		}
		newTransaction := entities.Transaction{
			ID:     newTransactionReq.ID,
			Status: "Dismissed",
		}
		if transaction.Status == "Accepted" {
			newTransaction.Status = newTransactionReq.Status
			totalReputation := transaction.User.Reputation - 3
			if totalReputation < 0 {
				totalReputation = 0
			}
			newBalance = transaction.User.Balance - 20000
			if newBalance < 0 {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"code":    http.StatusInternalServerError,
					"message": "Your Money is Not Enough For Cancel This Transaction",
				})
			}
			if _, err := transcon.Repo.UpdateUserReputation(uint(userId), totalReputation); err != nil {
				return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
			}
		}
		newBalance = transaction.User.Balance + transaction.Total
		if _, err := transcon.Repo.UpdateUserBalance(uint(userId), newBalance); err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		res, err := transcon.Repo.UpdateTransactionStatus(newTransaction)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		data := TransactionResponse{
			ID:           res.ID,
			UserID:       res.UserID,
			RestaurantID: res.RestaurantID,
			DateTime:     res.DateTime,
			Person:       res.Persons,
			Total:        res.Total,
		}
		response := TransactionResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    data,
		}
		return c.JSON(http.StatusOK, response)

	}
}
