package activities

import (
	"gorm.io/gorm"
	"rumahdermawan/backedn-rdi/model/domain"
	"rumahdermawan/backedn-rdi/model/web/activities"
)

type ActivityRepositoryImpl struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) *ActivityRepositoryImpl {
	return &ActivityRepositoryImpl{db}
}

func (repository *ActivityRepositoryImpl) SaveOperation(activity domain.OperationActivity) (domain.OperationActivity, error) {
	err := repository.db.Create(&activity).Error

	if err != nil {
		return activity, err
	}

	return activity, nil
}

func (repository *ActivityRepositoryImpl) FindByIdOperation(activity domain.OperationActivity) (domain.OperationActivity, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *ActivityRepositoryImpl) UpdateOperation(activity domain.OperationActivity) (domain.OperationActivity, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *ActivityRepositoryImpl) DeleteOperation(Id int) error {
	//TODO implement me
	panic("implement me")
}

func (repository *ActivityRepositoryImpl) FindAllOperation(param activities.ActivityQueryParam) ([]domain.OperationActivity, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *ActivityRepositoryImpl) SaveFunding(activity domain.FundingActivity) (domain.FundingActivity, error) {
	err := repository.db.Create(&activity).Error

	if err != nil {
		return activity, err
	}

	return activity, nil
}

func (repository *ActivityRepositoryImpl) FindByIdFunding(activity domain.FundingActivity) (domain.FundingActivity, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *ActivityRepositoryImpl) UpdateFunding(activity domain.FundingActivity) (domain.FundingActivity, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *ActivityRepositoryImpl) DeleteFunding(Id int) error {
	//TODO implement me
	panic("implement me")
}

func (repository *ActivityRepositoryImpl) FindAllFunding(param activities.ActivityQueryParam) ([]domain.FundingActivity, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *ActivityRepositoryImpl) SaveInvest(activity domain.InvestsActivity) (domain.InvestsActivity, error) {
	err := repository.db.Create(&activity).Error

	if err != nil {
		return activity, err
	}

	return activity, nil
}

func (repository *ActivityRepositoryImpl) FindByIdInvest(activity domain.InvestsActivity) (domain.InvestsActivity, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *ActivityRepositoryImpl) UpdateInvest(activity domain.InvestsActivity) (domain.InvestsActivity, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *ActivityRepositoryImpl) DeleteInvest(Id int) error {
	//TODO implement me
	panic("implement me")
}

func (repository *ActivityRepositoryImpl) FindAllInvest(param activities.ActivityQueryParam) ([]domain.InvestsActivity, error) {
	//TODO implement me
	panic("implement me")
}
