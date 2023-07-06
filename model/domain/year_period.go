package domain

type YearPeriod struct {
	Id         uint
	InfoPeriod []byte
}

type InfoPeriod struct {
	Month string
	Year  string
}
