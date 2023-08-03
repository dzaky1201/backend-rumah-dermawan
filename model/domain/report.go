package domain

type ReportActivity struct {
	Month string
	Total int
}

type AllActivity struct {
	Id              uint
	DateNote        string
	Description     string
	Amount          string
	TypeTransaction string
	YearPeriodId    uint
	YearPeriod      YearPeriod
}
