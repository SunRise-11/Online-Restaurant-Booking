package transactions

import (
	"Restobook/delivery/common"
	"Restobook/entities"
	"Restobook/repository/transactions"
	"net/http"
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
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))
		newTransactionReq := TransactionRequestFormat{}
		if err := c.Bind(&newTransactionReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		var dateTime, _ = time.Parse(time.RFC822, newTransactionReq.DateTime+" WIB")
		newTransaction := entities.Transaction{
			UserID:       uint(userID),
			RestaurantID: newTransactionReq.RestaurantID,
			DateTime:     dateTime,
			Persons:      newTransactionReq.Persons,
		}
		res, err := transcon.Repo.Create(newTransaction)
		if err != nil || res.ID == 0 {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		data := TransactionResponse{
			ID:           res.ID,
			UserID:       res.UserID,
			RestaurantID: res.RestaurantID,
			DateTime:     res.DateTime,
			Person:       res.Persons,
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
		data := []TransactionResponse{}
		for _, transaction := range transactions {
			data = append(
				data, TransactionResponse{
					ID:           transaction.ID,
					UserID:       transaction.UserID,
					RestaurantID: transaction.RestaurantID,
					Person:       transaction.Persons,
					DateTime:     transaction.DateTime,
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
