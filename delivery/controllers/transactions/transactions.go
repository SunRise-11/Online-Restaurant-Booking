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
		var balance, total int
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))
		newTransactionReq := TransactionRequestFormat{}
		if err := c.Bind(&newTransactionReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		var dateTime, _ = time.Parse(time.RFC822, newTransactionReq.DateTime+" WIB")
		if res, err := transcon.Repo.GetBalanceAndPriceResto(uint(userID), newTransactionReq.RestaurantID); err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
			// fmt.Println(result)
		} else {
			total = newTransactionReq.Persons * res.PriceResto
			balance = res.Balance - total
			if balance < 0 {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"code":    http.StatusInternalServerError,
					"message": "Your Money Not Enough For Booking",
				})
			}
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
		if res, err := transcon.Repo.Create(newTransaction); err != nil || res.ID == 0 {
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

func (transcon TransactionsController) GetAllWaitingCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))
		// transactions2, err2 := transcon.Repo.GetTransactionById(1)
		// fmt.Println(transactions2, err2)
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
		transactions, err := transcon.Repo.GetAllWaiting(uint(restoID))
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
		if restoID == 0 {
			return c.JSON(http.StatusUnauthorized, common.DefaultResponse{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
			})
		}
		newTransactionReq := TransactionRequestFormat{}
		if err := c.Bind(&newTransactionReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
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
func (transcon TransactionsController) RejectTransactionCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newBalance int
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restoID := int(claims["restoid"].(float64))
		if restoID == 0 {
			return c.JSON(http.StatusUnauthorized, common.DefaultResponse{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
			})
		}
		newTransactionReq := TransactionRequestFormat{}
		if err := c.Bind(&newTransactionReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		res, err := transcon.Repo.GetTransactionById(newTransactionReq.ID)
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
