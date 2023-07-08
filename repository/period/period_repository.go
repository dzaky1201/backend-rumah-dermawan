package period

import (
	"rumahdermawan/backedn-rdi/model/domain"
	"rumahdermawan/backedn-rdi/model/web/period"
)

type PeriodRepository interface {
	Save(period domain.YearPeriod) (domain.YearPeriod, error)
	FindById(period domain.YearPeriod) (domain.YearPeriod, error)
	Update(period domain.YearPeriod) (domain.YearPeriod, error)
	Delete(Id int) error
	FindAll(param period.PeriodQueryParam) ([]domain.YearPeriod, error)
}
