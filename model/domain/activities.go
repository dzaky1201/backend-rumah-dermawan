package domain

type OperationActivity struct {
	Id              uint
	DateNote        string
	Description     string
	Amount          string
	TypeTransaction string
	YearPeriodId    uint
}

type FundingActivity struct {
	Id              uint
	DateNote        string
	Description     string
	Amount          string
	TypeTransaction string
	YearPeriodId    uint
}
type InvestsActivity struct {
	Id              uint
	DateNote        string
	Description     string
	Amount          string
	TypeTransaction string
	YearPeriodId    uint
}
