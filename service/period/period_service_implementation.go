package period

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"rumahdermawan/backedn-rdi/helper"
	"rumahdermawan/backedn-rdi/model/domain"
	period2 "rumahdermawan/backedn-rdi/model/web/period"
	"rumahdermawan/backedn-rdi/repository/period"
)

type PeriodServiceImpl struct {
	PeriodRepository period.PeriodRepository
	Validate         *validator.Validate
}

func NewPeriodService(repository period.PeriodRepository, validate *validator.Validate) *PeriodServiceImpl {
	return &PeriodServiceImpl{
		PeriodRepository: repository,
		Validate:         validate,
	}
}

func (service *PeriodServiceImpl) Save(request period2.PeriodCreateRequest) period2.PeriodResponse {
	errReq := service.Validate.Struct(request)
	helper.PanicIfError(errReq)

	objectPeriod := domain.InfoPeriod{
		Month: request.Month,
		Year:  request.Year,
	}

	// converter to make jsonb postgres SQL
	var jsonData, err = json.Marshal(objectPeriod)

	if err != nil {
		fmt.Println(err.Error())
	}

	periodReq := domain.YearPeriod{
		InfoPeriod: jsonData,
	}

	yearRes, errRes := service.PeriodRepository.Save(periodReq)
	if errRes != nil {
		return period2.PeriodResponse{}
	}

	return helper.ToPeriodResponse(yearRes)

}

func (service *PeriodServiceImpl) Update(request period2.PeriodCreateRequest) period2.PeriodResponse {
	//TODO implement me
	panic("implement me")
}

func (service *PeriodServiceImpl) Delete(periodId int) {
	//TODO implement me
	panic("implement me")
}

func (service *PeriodServiceImpl) FindAll() []period2.PeriodResponse {
	//TODO implement me
	panic("implement me")
}
