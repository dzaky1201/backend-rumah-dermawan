package helper

import (
	"encoding/json"
	"fmt"
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
		Id:           yearModel.Id,
		Year:         data.Year,
		Month:        data.Month,
		Label:        fmt.Sprintf("%s %s", data.Month, data.Year),
		HaveRelation: yearModel.HaveRelation,
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

func ToActivityReportResponse(data []domain.ReportActivity) activities.ActivityReportResponse {
	var responseMonth []activities.ActivityReportMonth
	var responseTotal []activities.ActivityReportTotal
	var responseAllData []activities.ActivityReportAll

	for _, activity := range data {
		responseMonth = append(responseMonth, activities.ActivityReportMonth{Month: activity.Month})
		responseTotal = append(responseTotal, activities.ActivityReportTotal{Total: activity.Total})
		responseAllData = append(responseAllData, activities.ActivityReportAll{
			Month: activity.Month,
			Total: activity.Total,
		})
	}
	return activities.ActivityReportResponse{
		Month: responseMonth,
		Total: responseTotal, AllData: responseAllData,
	}
}

func ToYearResponseList(data []domain.YearList) []period.AllYearResponse {
	var response []period.AllYearResponse
	for _, yearList := range data {
		response = append(response, period.AllYearResponse{Year: yearList.Year})
	}
	return response
}
