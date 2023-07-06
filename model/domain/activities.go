package domain

type Activities struct {
	Id              uint
	InputDate       string
	Description     string
	Amount          uint
	TypeTransaction string
	PeriodID        uint
}
