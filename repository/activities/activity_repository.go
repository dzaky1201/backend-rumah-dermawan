package activities

import (
	"rumahdermawan/backedn-rdi/model/domain"
	"rumahdermawan/backedn-rdi/model/web/activities"
)

type ActivitiyRepository interface {
	SaveOperation(activity domain.OperationActivity) (domain.OperationActivity, error)
	FindByIdOperation(activity domain.OperationActivity) (domain.OperationActivity, error)
	UpdateOperation(activity domain.OperationActivity) (domain.OperationActivity, error)
	DeleteOperation(Id int) error
	FindAllOperation(param activities.ActivityQueryParam) ([]domain.OperationActivity, error)

	SaveFunding(activity domain.FundingActivity) (domain.FundingActivity, error)
	FindByIdFunding(activity domain.FundingActivity) (domain.FundingActivity, error)
	UpdateFunding(activity domain.FundingActivity) (domain.FundingActivity, error)
	DeleteFunding(Id int) error
	FindAllFunding(param activities.ActivityQueryParam) ([]domain.FundingActivity, error)

	SaveInvest(activity domain.InvestsActivity) (domain.InvestsActivity, error)
	FindByIdInvest(activity domain.InvestsActivity) (domain.InvestsActivity, error)
	UpdateInvest(activity domain.InvestsActivity) (domain.InvestsActivity, error)
	DeleteInvest(Id int) error
	FindAllInvest(param activities.ActivityQueryParam) ([]domain.InvestsActivity, error)
}
