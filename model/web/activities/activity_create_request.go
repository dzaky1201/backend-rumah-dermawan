package activities

type ActivityCreateRequest struct {
	InputDate       string `validate:"required" json:"input_date"`
	Description     string `validate:"required" json:"description"`
	Amount          uint   `validate:"required" json:"amount"`
	TypeTransaction string `validate:"required" json:"type_transaction"`
	PeriodID        uint   `validate:"required" json:"period_id"`
}
