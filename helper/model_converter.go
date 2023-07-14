package helper

import (
	"encoding/json"
	"rumahdermawan/backedn-rdi/model/domain"
	"rumahdermawan/backedn-rdi/model/web/activities"
	"rumahdermawan/backedn-rdi/model/web/period"
	"rumahdermawan/backedn-rdi/model/web/user"
)

func ToUserResponse(userModel domain.User, token string) user.UserResponse {
	return user.UserResponse{
		Id:    userModel.Id,
		Email: userModel.Email,
		Token: token,
	}
}

func ToPeriodResponse(yearModel domain.YearPeriod) period.PeriodResponse {

	var jsonData = yearModel.InfoPeriod

	var data domain.InfoPeriod

	var err = json.Unmarshal(jsonData, &data)

	PanicIfError(err)

	return period.PeriodResponse{
		Id:    yearModel.Id,
		Year:  data.Year,
		Month: data.Month,
	}
}

func ToPeriodResponseList(data []domain.YearPeriod) []period.PeriodResponse {
	var response []period.PeriodResponse
	for _, yearPeriod := range data {
		response = append(response, ToPeriodResponse(yearPeriod))
	}
	return response
}

func ToOperationActivityResponse(data domain.OperationActivity) activities.ActivityResponse {
	dataPeriod := ToPeriodResponse(data.YearPeriod)
	return activities.ActivityResponse{
		ID:              int(data.Id),
		InputDate:       data.DateNote,
		Description:     data.Description,
		Amount:          data.Amount,
		TypeTransaction: data.TypeTransaction,
		Period:          dataPeriod,
	}
}

func ToOperationActivityResponseList(data []domain.OperationActivity) []activities.ActivityResponse {
	var response []activities.ActivityResponse
	for _, responseList := range data {
		response = append(response, ToOperationActivityResponse(responseList))
	}
	return response
}

func ToFundingActivityResponse(data domain.FundingActivity) activities.ActivityResponse {
	dataPeriod := ToPeriodResponse(data.YearPeriod)
	return activities.ActivityResponse{
		ID:              int(data.Id),
		InputDate:       data.DateNote,
		Description:     data.Description,
		Amount:          data.Amount,
		TypeTransaction: data.TypeTransaction,
		Period:          dataPeriod,
	}
}

func ToFundingActivityResponseList(data []domain.FundingActivity) []activities.ActivityResponse {
	var response []activities.ActivityResponse
	for _, responseList := range data {
		response = append(response, ToFundingActivityResponse(responseList))
	}
	return response
}

func ToInvestActivityResponse(data domain.InvestsActivity) activities.ActivityResponse {
	dataPeriod := ToPeriodResponse(data.YearPeriod)
	return activities.ActivityResponse{
		ID:              int(data.Id),
		InputDate:       data.DateNote,
		Description:     data.Description,
		Amount:          data.Amount,
		TypeTransaction: data.TypeTransaction,
		Period:          dataPeriod,
	}
}

func ToInvestActivityResponseList(data []domain.InvestsActivity) []activities.ActivityResponse {
	var response []activities.ActivityResponse
	for _, responseList := range data {
		response = append(response, ToInvestActivityResponse(responseList))
	}
	return response
}

func ToActivityReportResponseObject(data domain.ReportActivity) activities.ActivityReportResponse {
	return activities.ActivityReportResponse{
		Month: data.Month,
		Total: data.Total,
	}
}

func ToActivityReportResponse(data []domain.ReportActivity) []activities.ActivityReportResponse {
	var response []activities.ActivityReportResponse

	for _, activity := range data {
		response = append(response, ToActivityReportResponseObject(activity))
	}
	return response
}
