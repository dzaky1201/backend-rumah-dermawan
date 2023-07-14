package activities

type ActivityResponse struct {
	ID              int         `json:"id"`
	InputDate       string      `json:"input_date"`
	Description     string      `json:"description"`
	Amount          string      `json:"amount"`
	TypeTransaction string      `json:"type_transaction"`
	Period          interface{} `json:"period"`
}

type ActivityReportResponse struct {
	Month string `json:"month"`
	Total string `json:"total"`
}
