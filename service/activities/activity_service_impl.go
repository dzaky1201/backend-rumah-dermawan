package activities

import (
	"github.com/go-playground/validator/v10"
	"rumahdermawan/backedn-rdi/helper"
	"rumahdermawan/backedn-rdi/model/domain"
	activities2 "rumahdermawan/backedn-rdi/model/web/activities"
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

func (service *ActivityServiceImpl) Update(request activities2.ActivityCreateRequest, pathId int, updateType string) (activities2.ActivityResponse, error) {
	if updateType == "operation" {
		data, err := service.ActivityRepository.FindByIdOperation(pathId)
		if err != nil {
			return activities2.ActivityResponse{}, err
		}

		data.Amount = request.Amount
		data.Description = request.Description
		data.PeriodeID = request.PeriodID
		data.TypeTransaction = request.TypeTransaction
		data.DateNote = request.InputDate

		update, errUpdate := service.ActivityRepository.UpdateOperation(data)
		if errUpdate != nil {
			return activities2.ActivityResponse{}, errUpdate
		}

		return helper.ToOperationActivityResponse(update, update.PeriodeID), nil
	} else if updateType == "invest" {
		data, err := service.ActivityRepository.FindByIdInvest(pathId)
		if err != nil {
			return activities2.ActivityResponse{}, err
		}

		data.Amount = request.Amount
		data.Description = request.Description
		data.PeriodeID = request.PeriodID
		data.TypeTransaction = request.TypeTransaction
		data.DateNote = request.InputDate

		update, errUpdate := service.ActivityRepository.UpdateInvest(data)
		if errUpdate != nil {
			return activities2.ActivityResponse{}, errUpdate
		}

		return helper.ToInvestActivityResponse(update, update.PeriodeID), nil
	} else if updateType == "funding" {
		data, err := service.ActivityRepository.FindByIdFunding(pathId)
		if err != nil {
			return activities2.ActivityResponse{}, err
		}

		data.Amount = request.Amount
		data.Description = request.Description
		data.PeriodeID = request.PeriodID
		data.TypeTransaction = request.TypeTransaction
		data.DateNote = request.InputDate

		update, errUpdate := service.ActivityRepository.UpdateFunding(data)
		if errUpdate != nil {
			return activities2.ActivityResponse{}, errUpdate
		}

		return helper.ToFundingActivityResponse(update, update.PeriodeID), nil
	}

	return activities2.ActivityResponse{}, nil
}

func (service *ActivityServiceImpl) FindById(activityId int, findType string) (activities2.ActivityResponse, error) {
	if findType == "operation" {
		data, err := service.ActivityRepository.FindByIdOperation(activityId)

		if err != nil {
			return activities2.ActivityResponse{}, err
		}

		return helper.ToOperationActivityResponse(data, data.PeriodeID), nil
	} else if findType == "invest" {
		data, err := service.ActivityRepository.FindByIdInvest(activityId)

		if err != nil {
			return activities2.ActivityResponse{}, err
		}

		return helper.ToInvestActivityResponse(data, data.PeriodeID), nil
	} else if findType == "funding" {
		data, err := service.ActivityRepository.FindByIdFunding(activityId)

		if err != nil {
			return activities2.ActivityResponse{}, err
		}

		return helper.ToFundingActivityResponse(data, data.PeriodeID), nil
	}

	return activities2.ActivityResponse{}, nil
}

func (service *ActivityServiceImpl) Delete(activityId int, deleteType string) error {
	//TODO implement me
	panic("implement me")
}

func (service *ActivityServiceImpl) FindAll(page int, limit int, year string, month string, findAllType string) ([]activities2.ActivityResponse, error) {
	//TODO implement me
	panic("implement me")
}
