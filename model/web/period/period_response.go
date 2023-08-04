package period

type PeriodResponse struct {
	Id           uint   `json:"id"`
	Year         string `json:"year"`
	Month        string `json:"month"`
	Label        string `json:"label"`
	HaveRelation int    `json:"have_relation"`
}
