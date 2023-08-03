package domain

import "time"

type OperationActivity struct {
	Id              uint
	DateNote        string
	Description     string
	Amount          string
	TypeTransaction string
	YearPeriodId    uint
	YearPeriod      YearPeriod
	UpdatedAt       time.Time
}

type FundingActivity struct {
	Id              uint
	DateNote        string
	Description     string
	Amount          string
	TypeTransaction string
	YearPeriodId    uint
	YearPeriod      YearPeriod
	UpdatedAt       time.Time
}
type InvestsActivity struct {
	Id              uint
	DateNote        string
	Description     string
	Amount          string
	TypeTransaction string
	YearPeriodId    uint
	YearPeriod      YearPeriod
	UpdatedAt       time.Time
}
