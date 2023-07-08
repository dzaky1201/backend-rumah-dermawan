package helper

import (
	"encoding/json"
	"rumahdermawan/backedn-rdi/model/domain"
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
