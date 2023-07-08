package period

import (
	"rumahdermawan/backedn-rdi/model/domain"
)

type PeriodRepository interface {
	Save(period domain.YearPeriod) (domain.YearPeriod, error)
	FindById(period domain.YearPeriod) (domain.YearPeriod, error)
	Update(period domain.YearPeriod) (domain.YearPeriod, error)
	Delete(Id int) error
	FindAll() []domain.YearPeriod
}
