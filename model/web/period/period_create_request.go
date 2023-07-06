package period

type PeriodCreateRequest struct {
	Year  string `validate:"required" json:"year"`
	Month string `validate:"required" json:"month"`
}
