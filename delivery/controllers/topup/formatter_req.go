package topup

type TopUpRequestFormat struct {
	Total int `json:"total" form:"total"`
}

type CallbackRequest struct {
	ExternalID string `json:"external_id"`
	Status     string `json:"status"`
}
