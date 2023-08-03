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
	AllData interface{} `json:"allData"`
	Month   interface{} `json:"month"`
	Total   interface{} `json:"total"`
}

type ActivityReportAll struct {
	Month string `json:"month"`
	Total int    `json:"total"`
}

type ActivityReportMonth struct {
	Month string `json:"name"`
}

type ActivityReportTotal struct {
	Total int `json:"total"`
}
