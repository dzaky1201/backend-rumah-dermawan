package domain

type OperationActivity struct {
	Id              uint
	DateNote        string
	Description     string
	Amount          uint
	TypeTransaction string
	PeriodeID       uint
}

type FundingActivity struct {
	Id              uint
	DateNote        string
	Description     string
	Amount          uint
	TypeTransaction string
	PeriodeID       uint
}
type InvestsActivity struct {
	Id              uint
	DateNote        string
	Description     string
	Amount          uint
	TypeTransaction string
	PeriodeID       uint
}
