package topup

import (
	"Restobook/delivery/common"
	"Restobook/entities"
	"Restobook/repository/topup"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

type TopUpController struct {
	Repo topup.TopUpInterface
}

func NewTopUpControllers(turep topup.TopUpInterface) *TopUpController {
	return &TopUpController{Repo: turep}
}

func (tc TopUpController) TopUp() echo.HandlerFunc {

	return func(c echo.Context) error {
		topuprequest := TopUpRequestFormat{}

		if err := c.Bind(&topuprequest); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))

		invoiceId := strings.ToUpper(strings.ReplaceAll(uuid.New().String(), "-", ""))

		data := entities.TopUp{
			UserID:    uint(userID),
			InvoiceID: invoiceId,
			Total:     topuprequest.Total,
			Status:    "PENDING",
		}

		topUpData, err := tc.Repo.Create(data)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		topUpPayment, err := CreateInvoice(topUpData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		updateData := entities.TopUp{}
		updateData.PaymentUrl = topUpPayment.PaymentUrl
		tc.Repo.Update(invoiceId, updateData)

		responsedata := TopUpResponse{
			ID:         topUpData.ID,
			UserID:     uint(userID),
			InvoiceID:  invoiceId,
			PaymentUrl: topUpPayment.PaymentUrl,
			Total:      topUpData.Total,
			Status:     "PENDING",
		}

		response := TopUpResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    responsedata,
		}

		return c.JSON(http.StatusOK, response)
	}
}

func (tc TopUpController) GetAllWaiting() echo.HandlerFunc {

	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))

		if res, err := tc.Repo.GetAllWaiting(uint(userID)); err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {
			topUpDatas := []TopUpResponse{}
			for _, x := range res {
				topUpDatas = append(topUpDatas, TopUpResponse{
					ID:         x.ID,
					UserID:     x.UserID,
					InvoiceID:  x.InvoiceID,
					PaymentUrl: x.PaymentUrl,
					Total:      x.Total,
					Status:     x.Status,
				})
			}

			response := TopUpResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Data:    topUpDatas,
			}

			return c.JSON(http.StatusOK, response)
		}
	}
}

func (tc TopUpController) GetAllPaid() echo.HandlerFunc {

	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))

		if res, err := tc.Repo.GetAllPaid(uint(userID)); err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {
			topUpDatas := []TopUpResponse{}
			for _, x := range res {
				topUpDatas = append(topUpDatas, TopUpResponse{
					ID:         x.ID,
					UserID:     x.UserID,
					InvoiceID:  x.InvoiceID,
					PaymentUrl: x.PaymentUrl,
					Total:      x.Total,
					Status:     x.Status,
				})
			}

			response := TopUpResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Data:    topUpDatas,
			}

			return c.JSON(http.StatusOK, response)
		}
	}
}

func (tc TopUpController) Callback() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		headers := req.Header

		xCallbackToken := headers.Get("X-Callback-Token")

		if xCallbackToken != common.XENDIT_CALLBACK_TOKEN {
			return c.JSON(http.StatusNotAcceptable, common.NewStatusNotAcceptable())
		}

		var callbackRequest CallbackRequest
		if err := c.Bind(&callbackRequest); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		invoice, err := tc.Repo.GetByInvoice(callbackRequest.ExternalID)
		if err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

		var data entities.TopUp
		data.Status = callbackRequest.Status

		_, err = tc.Repo.Update(callbackRequest.ExternalID, data)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		if data.Status == "PAID" {
			user, _ := tc.Repo.GetUser(int(invoice.UserID))

			newBalance := entities.User{
				Balance: (user.Balance + invoice.Total),
			}

			tc.Repo.UpdateUserBalance(int(invoice.UserID), newBalance)
		}

		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}
}

//FUNC FOR XENDIT
func CreateInvoice(topUp entities.TopUp) (entities.TopUp, error) {
	xendit.Opt.SecretKey = os.Getenv("XENDIT_SECRET_KEY")

	data := invoice.CreateParams{
		ExternalID:  topUp.InvoiceID,
		Amount:      float64(topUp.Total),
		Description: "Invoice " + topUp.InvoiceID,
	}

	resp, err := invoice.Create(&data)
	if err != nil {
		return topUp, err
	}

	topUpSuccess := entities.TopUp{
		UserID:     topUp.UserID,
		InvoiceID:  topUp.InvoiceID,
		PaymentUrl: resp.InvoiceURL,
		Total:      int(resp.Amount),
		Status:     resp.Status,
	}

	return topUpSuccess, nil
}
