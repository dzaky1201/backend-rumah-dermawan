package activities

import (
	"github.com/go-playground/validator/v10"
	"rumahdermawan/backedn-rdi/helper"
	"rumahdermawan/backedn-rdi/model/domain"
	activities2 "rumahdermawan/backedn-rdi/model/web/activities"
	"rumahdermawan/backedn-rdi/model/web/period"
	"rumahdermawan/backedn-rdi/repository/activities"
)

type ActivityServiceImpl struct {
	ActivityRepository activities.ActivitiyRepository
	Validate           *validator.Validate
}

func NewActivityService(repository activities.ActivitiyRepository, validate *validator.Validate) *ActivityServiceImpl {
	return &ActivityServiceImpl{
		ActivityRepository: repository,
		Validate:           validate,
	}
}

func (service *ActivityServiceImpl) Save(request activities2.ActivityCreateRequest, createType string) (activities2.ActivityResponse, error) {
	errReq := service.Validate.Struct(request)

	if errReq != nil {
		return activities2.ActivityResponse{}, errReq
	}

	if createType == "operation" {
		createRequest := domain.OperationActivity{
			DateNote:        request.InputDate,
			Description:     request.Description,
			Amount:          request.Amount,
			TypeTransaction: request.TypeTransaction,
			PeriodeID:       request.PeriodID,
		}
		response, errRes := service.ActivityRepository.SaveOperation(createRequest)
		if errRes != nil {
			return activities2.ActivityResponse{}, errRes
		}
		return helper.ToOperationActivityResponse(response, response.PeriodeID), nil
	} else if createType == "invest" {
		createRequest := domain.InvestsActivity{
			DateNote:        request.InputDate,
			Description:     request.Description,
			Amount:          request.Amount,
			TypeTransaction: request.TypeTransaction,
			PeriodeID:       request.PeriodID,
		}
		response, errRes := service.ActivityRepository.SaveInvest(createRequest)
		if errRes != nil {
			return activities2.ActivityResponse{}, errRes
		}
		return helper.ToInvestActivityResponse(response, response.PeriodeID), nil
	} else if createType == "funding" {
		createRequest := domain.FundingActivity{
			DateNote:        request.InputDate,
			Description:     request.Description,
			Amount:          request.Amount,
			TypeTransaction: request.TypeTransaction,
			PeriodeID:       request.PeriodID,
		}
		response, errRes := service.ActivityRepository.SaveFunding(createRequest)
		if errRes != nil {
			return activities2.ActivityResponse{}, errRes
		}
		return helper.ToFundingActivityResponse(response, response.PeriodeID), nil
	}

	return activities2.ActivityResponse{}, nil

}

func (service *ActivityServiceImpl) Update(request activities2.ActivityCreateRequest, pathId int) (activities2.ActivityResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (service *ActivityServiceImpl) FindById(request int) (period.PeriodResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (service *ActivityServiceImpl) Delete(activityId int) error {
	//TODO implement me
	panic("implement me")
}

func (service *ActivityServiceImpl) FindAll(page int, limit int, year string, month string) ([]activities2.ActivityResponse, error) {
	//TODO implement me
	panic("implement me")
}
