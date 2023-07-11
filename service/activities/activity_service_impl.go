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
			YearPeriodId:    request.YearPeriodId,
		}
		response, errRes := service.ActivityRepository.SaveOperation(createRequest)
		if errRes != nil {
			return activities2.ActivityResponse{}, errRes
		}
		return helper.ToOperationActivityResponse(response, response.YearPeriodId), nil
	} else if createType == "invest" {
		createRequest := domain.InvestsActivity{
			DateNote:        request.InputDate,
			Description:     request.Description,
			Amount:          request.Amount,
			TypeTransaction: request.TypeTransaction,
			YearPeriodId:    request.YearPeriodId,
		}
		response, errRes := service.ActivityRepository.SaveInvest(createRequest)
		if errRes != nil {
			return activities2.ActivityResponse{}, errRes
		}
		return helper.ToInvestActivityResponse(response, response.YearPeriodId), nil
	} else if createType == "funding" {
		createRequest := domain.FundingActivity{
			DateNote:        request.InputDate,
			Description:     request.Description,
			Amount:          request.Amount,
			TypeTransaction: request.TypeTransaction,
			YearPeriodId:    request.YearPeriodId,
		}
		response, errRes := service.ActivityRepository.SaveFunding(createRequest)
		if errRes != nil {
			return activities2.ActivityResponse{}, errRes
		}
		return helper.ToFundingActivityResponse(response, response.YearPeriodId), nil
	}

	return activities2.ActivityResponse{}, nil

}

func (service *ActivityServiceImpl) Update(request activities2.ActivityCreateRequest, pathId int, updateType string) (activities2.ActivityResponse, error) {
	if updateType == "operation" {
		data, err := service.ActivityRepository.FindByIdOperation(pathId)
		if err != nil {
			return activities2.ActivityResponse{}, err
		}

		if request.InputDate == "" {
			request.InputDate = data.DateNote
		}
		if request.Amount == "" {
			request.Amount = data.Amount
		}
		if request.Description == "" {
			request.Description = data.Description
		}
		if request.TypeTransaction == "" {
			request.TypeTransaction = data.TypeTransaction
		}
		if request.YearPeriodId == 0 {
			request.YearPeriodId = data.YearPeriodId
		}

		activityReq := domain.OperationActivity{
			Id:              uint(pathId),
			DateNote:        request.InputDate,
			Description:     request.Description,
			Amount:          request.Amount,
			TypeTransaction: request.TypeTransaction,
			YearPeriodId:    request.YearPeriodId,
		}

		update, errUpdate := service.ActivityRepository.UpdateOperation(activityReq)
		if errUpdate != nil {
			return activities2.ActivityResponse{}, errUpdate
		}

		return helper.ToOperationActivityResponse(update, update.YearPeriodId), nil
	}
	if updateType == "invest" {
		data, err := service.ActivityRepository.FindByIdInvest(pathId)
		if err != nil {
			return activities2.ActivityResponse{}, err
		}

		if request.InputDate == "" {
			request.InputDate = data.DateNote
		}
		if request.Amount == "" {
			request.Amount = data.Amount
		}
		if request.Description == "" {
			request.Description = data.Description
		}
		if request.TypeTransaction == "" {
			request.TypeTransaction = data.TypeTransaction
		}
		if request.YearPeriodId == 0 {
			request.YearPeriodId = data.YearPeriodId
		}

		activityReq := domain.InvestsActivity{
			Id:              uint(pathId),
			DateNote:        request.InputDate,
			Description:     request.Description,
			Amount:          request.Amount,
			TypeTransaction: request.TypeTransaction,
			YearPeriodId:    request.YearPeriodId,
		}

		update, errUpdate := service.ActivityRepository.UpdateInvest(activityReq)
		if errUpdate != nil {
			return activities2.ActivityResponse{}, errUpdate
		}

		return helper.ToInvestActivityResponse(update, update.YearPeriodId), nil
	}
	if updateType == "funding" {
		data, err := service.ActivityRepository.FindByIdFunding(pathId)
		if err != nil {
			return activities2.ActivityResponse{}, err
		}

		if request.InputDate == "" {
			request.InputDate = data.DateNote
		}
		if request.Amount == "" {
			request.Amount = data.Amount
		}
		if request.Description == "" {
			request.Description = data.Description
		}
		if request.TypeTransaction == "" {
			request.TypeTransaction = data.TypeTransaction
		}
		if request.YearPeriodId == 0 {
			request.YearPeriodId = data.YearPeriodId
		}

		activityReq := domain.FundingActivity{
			Id:              uint(pathId),
			DateNote:        request.InputDate,
			Description:     request.Description,
			Amount:          request.Amount,
			TypeTransaction: request.TypeTransaction,
			YearPeriodId:    request.YearPeriodId,
		}

		update, errUpdate := service.ActivityRepository.UpdateFunding(activityReq)
		if errUpdate != nil {
			return activities2.ActivityResponse{}, errUpdate
		}

		return helper.ToFundingActivityResponse(update, update.YearPeriodId), nil
	}

	return activities2.ActivityResponse{}, nil
}

func (service *ActivityServiceImpl) FindById(activityId int, findType string) (activities2.ActivityResponse, error) {
	if findType == "operation" {
		data, err := service.ActivityRepository.FindByIdOperation(activityId)

		if err != nil {
			return activities2.ActivityResponse{}, err
		}

		return helper.ToOperationActivityResponse(data, data.YearPeriodId), nil
	} else if findType == "invest" {
		data, err := service.ActivityRepository.FindByIdInvest(activityId)

		if err != nil {
			return activities2.ActivityResponse{}, err
		}

		return helper.ToInvestActivityResponse(data, data.YearPeriodId), nil
	} else if findType == "funding" {
		data, err := service.ActivityRepository.FindByIdFunding(activityId)

		if err != nil {
			return activities2.ActivityResponse{}, err
		}

		return helper.ToFundingActivityResponse(data, data.YearPeriodId), nil
	}

	return activities2.ActivityResponse{}, nil
}

func (service *ActivityServiceImpl) Delete(activityId int, deleteType string) error {

	if deleteType == "operation" {
		err := service.ActivityRepository.DeleteOperation(activityId)
		if err != nil {
			return err
		}
		return nil
	} else if deleteType == "invest" {
		err := service.ActivityRepository.DeleteInvest(activityId)
		if err != nil {
			return err
		}
		return nil
	} else if deleteType == "funding" {
		err := service.ActivityRepository.DeleteFunding(activityId)
		if err != nil {
			return err
		}
		return nil
	}

	return nil

}

func (service *ActivityServiceImpl) FindAll(page int, limit int, year string, month string, findAllType string) ([]activities2.ActivityResponse, error) {
	//TODO implement me
	panic("implement me")
}
