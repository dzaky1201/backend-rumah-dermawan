package period

import (
	"rumahdermawan/backedn-rdi/model/web/period"
)

type PeriodService interface {
	Save(request period.PeriodCreateRequest) period.PeriodResponse
	Update(request period.PeriodCreateRequest) period.PeriodResponse
	Delete(periodId int)
	FindAll() []period.PeriodResponse
}
