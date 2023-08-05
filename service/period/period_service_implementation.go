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

func (service *PeriodServiceImpl) Save(request period2.PeriodCreateRequest) (period2.PeriodResponse, error) {
	errReq := service.Validate.Struct(request)
	if errReq != nil {
		return period2.PeriodResponse{}, errReq
	}

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
		return period2.PeriodResponse{}, errRes
	}

	return helper.ToPeriodResponse(yearRes), nil

}

func (service *PeriodServiceImpl) Update(request period2.PeriodCreateRequest, pathId domain.YearPeriod) (period2.PeriodResponse, error) {
	getDataDetail, err := service.PeriodRepository.FindById(pathId)
	if err != nil {
		return period2.PeriodResponse{}, err
	}

	// data dari periode yang mau diupdate
	periodConvert := helper.ToPeriodResponse(getDataDetail)

	// jika bulan dan tahun diupdate salah satu aja
	if request.Year == "" {
		request.Year = periodConvert.Year
	} else if request.Month == "" {
		request.Month = periodConvert.Month
	}

	objectPeriod := domain.InfoPeriod{
		Month: request.Month,
		Year:  request.Year,
	}

	// converter to make jsonb postgres SQL
	var jsonData, errConv = json.Marshal(objectPeriod)

	if errConv != errConv {
		return period2.PeriodResponse{}, errConv
	}

	periodReq := domain.YearPeriod{
		InfoPeriod: jsonData,
		Id:         pathId.Id,
	}

	data, ErrUpdate := service.PeriodRepository.Update(periodReq)
	if ErrUpdate != nil {
		return period2.PeriodResponse{}, ErrUpdate
	}

	return helper.ToPeriodResponse(data), nil
}

func (service *PeriodServiceImpl) Delete(periodId int) error {
	err := service.PeriodRepository.Delete(periodId)
	if err != nil {
		return err
	}
	return nil
}

func (service *PeriodServiceImpl) FindAll(page int, limit int, year string) ([]period2.PeriodResponse, error) {
	param := period2.PeriodQueryParam{Page: page, Limit: limit, Year: year}
	data, err := service.PeriodRepository.FindAll(param)
	if err != nil {
		return helper.ToPeriodResponseList(data), err
	}

	return helper.ToPeriodResponseList(data), nil
}

func (service *PeriodServiceImpl) FindById(request domain.YearPeriod) (period2.PeriodResponse, error) {
	data, err := service.PeriodRepository.FindById(request)

	if err != nil {
		return period2.PeriodResponse{}, err
	}

	return helper.ToPeriodResponse(data), nil
}

func (service *PeriodServiceImpl) FindAllYear() ([]period2.AllYearResponse, error) {
	data, err := service.PeriodRepository.FindAllYear()
	if err != nil {
		return helper.ToYearResponseList(data), err
	}

	return helper.ToYearResponseList(data), nil
}
