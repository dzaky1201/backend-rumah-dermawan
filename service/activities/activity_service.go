package activities

import (
	"rumahdermawan/backedn-rdi/model/web/activities"
	"rumahdermawan/backedn-rdi/model/web/period"
)

type ActivityService interface {
	Save(request activities.ActivityCreateRequest, createType string) (activities.ActivityResponse, error)
	Update(request activities.ActivityCreateRequest, pathId int) (activities.ActivityResponse, error)
	FindById(request int) (period.PeriodResponse, error)
	Delete(activityId int) error
	FindAll(page int, limit int, year string, month string) ([]activities.ActivityResponse, error)
}
