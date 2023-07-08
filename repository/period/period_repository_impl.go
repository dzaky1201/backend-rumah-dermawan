package period

import (
	"gorm.io/gorm"
	"rumahdermawan/backedn-rdi/model/domain"
	"rumahdermawan/backedn-rdi/model/web/period"
)

type PeriodRepositoryImpl struct {
	db *gorm.DB
}

func NewPeriodRepository(db *gorm.DB) *PeriodRepositoryImpl {
	return &PeriodRepositoryImpl{db: db}
}

func (repository *PeriodRepositoryImpl) Save(period domain.YearPeriod) (domain.YearPeriod, error) {
	err := repository.db.Create(&period).Error

	if err != nil {
		return period, err
	}

	return period, nil
}

func (repository *PeriodRepositoryImpl) Update(period domain.YearPeriod) (domain.YearPeriod, error) {
	err := repository.db.Model(domain.YearPeriod{}).Where("id = ?", period.Id).Updates(domain.YearPeriod{InfoPeriod: period.InfoPeriod}).Error
	if err != nil {
		return period, err
	}

	return period, nil
}

func (repository *PeriodRepositoryImpl) Delete(iD int) error {
	err := repository.db.Delete(&domain.YearPeriod{}, iD).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository *PeriodRepositoryImpl) FindAll(param period.PeriodQueryParam) ([]domain.YearPeriod, error) {
	var periods []domain.YearPeriod
	if param.Page == 0 {
		param.Page = 1
	}
	switch {
	case param.Limit > 20:
		param.Limit = 20
	case param.Limit <= 0:
		param.Limit = 10
	}
	offset := (param.Page - 1) * param.Limit

	if param.Year != "" {
		err := repository.db.Raw("select * from year_periods where year_periods.info_period ->> 'Year' @@ to_tsquery(?) offset ? limit ?", param.Year, offset, param.Limit).Scan(&periods).Error
		if err != nil {
			return periods, err
		}
	} else {
		err := repository.db.Limit(param.Limit).Offset(offset).Find(&periods).Error
		if err != nil {
			return periods, err
		}
	}

	return periods, nil
}

func (repository *PeriodRepositoryImpl) FindById(period domain.YearPeriod) (domain.YearPeriod, error) {
	var dataPeriod domain.YearPeriod
	err := repository.db.First(&dataPeriod, period.Id).Error
	if err != nil {
		return period, err
	}

	return dataPeriod, nil
}
