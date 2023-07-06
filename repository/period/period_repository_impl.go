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
	//TODO implement me
	panic("implement me")
}

func (repository *PeriodRepositoryImpl) Delete(period domain.YearPeriod) {
	//TODO implement me
	panic("implement me")
}

func (repository *PeriodRepositoryImpl) FindAll() []domain.YearPeriod {
	//TODO implement me
	panic("implement me")
}
