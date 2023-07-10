package activities

type ActivityResponse struct {
	InputDate       string      `json:"input_date"`
	Description     string      `json:"description"`
	Amount          uint        `json:"amount"`
	TypeTransaction string      `json:"type_transaction"`
	Period          interface{} `json:"period"`
}
