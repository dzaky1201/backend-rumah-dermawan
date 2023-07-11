package domain

type OperationActivity struct {
	Id              uint
	DateNote        string
	Description     string
	Amount          string
	TypeTransaction string
	YearPeriodId    uint
	YearPeriod      YearPeriod
}

type FundingActivity struct {
	Id              uint
	DateNote        string
	Description     string
	Amount          string
	TypeTransaction string
	YearPeriodId    uint
	YearPeriod      YearPeriod
}
type InvestsActivity struct {
	Id              uint
	DateNote        string
	Description     string
	Amount          string
	TypeTransaction string
	YearPeriodId    uint
	YearPeriod      YearPeriod
}
