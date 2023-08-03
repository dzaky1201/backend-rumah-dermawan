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

		dataPeriod, _ := service.ActivityRepository.FindByIdOperation(int(response.Id))
		response.YearPeriod = dataPeriod.YearPeriod

		return helper.ToOperationActivityResponse(response), nil
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

		dataPeriod, _ := service.ActivityRepository.FindByIdInvest(int(response.Id))
		response.YearPeriod = dataPeriod.YearPeriod

		return helper.ToInvestActivityResponse(response), nil
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

		dataPeriod, _ := service.ActivityRepository.FindByIdFunding(int(response.Id))
		response.YearPeriod = dataPeriod.YearPeriod

		return helper.ToFundingActivityResponse(response), nil
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
		// for show month and year name
		update.YearPeriod = data.YearPeriod

		return helper.ToOperationActivityResponse(update), nil
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

		update.YearPeriod = data.YearPeriod

		return helper.ToInvestActivityResponse(update), nil
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

		update.YearPeriod = data.YearPeriod

		return helper.ToFundingActivityResponse(update), nil
	}

	return activities2.ActivityResponse{}, nil
}

func (service *ActivityServiceImpl) FindById(activityId int, findType string) (activities2.ActivityResponse, error) {
	if findType == "operation" {
		data, err := service.ActivityRepository.FindByIdOperation(activityId)

		if err != nil {
			return activities2.ActivityResponse{}, err
		}

		return helper.ToOperationActivityResponse(data), nil
	} else if findType == "invest" {
		data, err := service.ActivityRepository.FindByIdInvest(activityId)

		if err != nil {
			return activities2.ActivityResponse{}, err
		}

		return helper.ToInvestActivityResponse(data), nil
	} else if findType == "funding" {
		data, err := service.ActivityRepository.FindByIdFunding(activityId)

		if err != nil {
			return activities2.ActivityResponse{}, err
		}

		return helper.ToFundingActivityResponse(data), nil
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

func (service *ActivityServiceImpl) FindAll(param activities2.ActivityQueryParam, findAllType string) ([]activities2.ActivityResponse, int, int, error) {
	switch {
	case findAllType == "operation":
		data, count, offset, err := service.ActivityRepository.FindAllOperation(param)
		if err != nil {
			return helper.ToOperationActivityResponseList(data), 0, 0, err
		}

		return helper.ToOperationActivityResponseList(data), count, offset, nil
	case findAllType == "funding":
		data, count, offset, err := service.ActivityRepository.FindAllFunding(param)
		if err != nil {
			return helper.ToFundingActivityResponseList(data), 0, 0, err
		}

		return helper.ToFundingActivityResponseList(data), count, offset, nil
	case findAllType == "invest":
		data, count, offset, err := service.ActivityRepository.FindAllInvest(param)
		if err != nil {
			return helper.ToInvestActivityResponseList(data), 0, 0, err
		}

		return helper.ToInvestActivityResponseList(data), count, offset, nil
	}

	return []activities2.ActivityResponse{}, 0, 0, nil
}

func (service *ActivityServiceImpl) ReportActivityAll(year string) (activities2.ActivityReportResponse, error) {
	data, err := service.ActivityRepository.ReportActivity(year)
	if err != nil {
		return helper.ToActivityReportResponse(data), err
	}
	return helper.ToActivityReportResponse(data), nil
}
