package period

import (
	"gorm.io/gorm"
	"rumahdermawan/backedn-rdi/model/domain"
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

func (repository *PeriodRepositoryImpl) FindAll() []domain.YearPeriod {
	//TODO implement me
	panic("implement me")
}

func (repository *PeriodRepositoryImpl) FindById(period domain.YearPeriod) (domain.YearPeriod, error) {
	var dataPeriod domain.YearPeriod
	err := repository.db.First(&dataPeriod, period.Id).Error
	if err != nil {
		return period, err
	}

	return dataPeriod, nil
}
