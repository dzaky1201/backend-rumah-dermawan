package period

import (
	"rumahdermawan/backedn-rdi/model/domain"
	"rumahdermawan/backedn-rdi/model/web/period"
)

type PeriodService interface {
	Save(request period.PeriodCreateRequest) (period.PeriodResponse, error)
	Update(request period.PeriodCreateRequest, pathId domain.YearPeriod) (period.PeriodResponse, error)
	FindById(request domain.YearPeriod) (period.PeriodResponse, error)
	Delete(periodId int) error
	FindAll(page int, limit int, year string) ([]period.PeriodResponse, error)
}
