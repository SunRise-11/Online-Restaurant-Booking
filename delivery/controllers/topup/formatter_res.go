package topup

type TopUpResponseFormat struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type TopUpResponse struct {
	ID         uint   `json:"id"`
	UserID     uint   `json:"user_id"`
	InvoiceID  string `json:"invoice_id"`
	PaymentUrl string `json:"payment_url"`
	Total      int    `json:"total"`
	Status     string `json:"status"`
}
