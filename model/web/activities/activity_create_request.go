package activities

type ActivityCreateRequest struct {
	InputDate       string `validate:"required" json:"input_date"`
	Description     string `validate:"required" json:"description"`
	Amount          string `validate:"required" json:"amount"`
	TypeTransaction string `validate:"required" json:"type_transaction"`
	YearPeriodId    uint   `validate:"required" json:"year_period_id"`
}
