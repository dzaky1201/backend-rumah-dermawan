package activities

import (
	"rumahdermawan/backedn-rdi/model/web/activities"
)

type ActivityService interface {
	Save(request activities.ActivityCreateRequest, createType string) (activities.ActivityResponse, error)
	Update(request activities.ActivityCreateRequest, pathId int, updateType string) (activities.ActivityResponse, error)
	FindById(activityId int, findType string) (activities.ActivityResponse, error)
	Delete(activityId int, deleteType string) error
	FindAll(param activities.ActivityQueryParam, findAllType string) ([]activities.ActivityResponse, error)
	ReportActivityAll(year string) ([]activities.ActivityReportResponse, error)
}
