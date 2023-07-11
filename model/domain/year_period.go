package domain

type YearPeriod struct {
	Id                uint
	InfoPeriod        []byte
	OperationActivity []OperationActivity
	FundingActivity   []FundingActivity
	InvestsActivity   []InvestsActivity
}

type InfoPeriod struct {
	Month string
	Year  string
}
