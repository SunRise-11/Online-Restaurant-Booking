package helpers

import (
	"Restobook/entities"

	"github.com/xendit/xendit-go/invoice"
)

func CreateInvoice(topUp entities.TopUp) (entities.TopUp, error) {
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
