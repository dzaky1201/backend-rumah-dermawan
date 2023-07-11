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

func (repository *ActivityRepositoryImpl) FindByIdOperation(Id int) (domain.OperationActivity, error) {
	var dataOperation domain.OperationActivity
	err := repository.db.First(&dataOperation, Id).Error
	if err != nil {
		return dataOperation, err
	}

	return dataOperation, nil
}

func (repository *ActivityRepositoryImpl) UpdateOperation(activity domain.OperationActivity) (domain.OperationActivity, error) {
	err := repository.db.Model(domain.OperationActivity{}).Where("id = ?", activity.Id).Updates(activity).Error
	if err != nil {
		return activity, err
	}

	return activity, nil
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

func (repository *ActivityRepositoryImpl) FindByIdFunding(Id int) (domain.FundingActivity, error) {
	var dataFunding domain.FundingActivity
	err := repository.db.First(&dataFunding, Id).Error
	if err != nil {
		return dataFunding, err
	}

	return dataFunding, nil
}

func (repository *ActivityRepositoryImpl) UpdateFunding(activity domain.FundingActivity) (domain.FundingActivity, error) {
	err := repository.db.Model(domain.FundingActivity{}).Where("id = ?", activity.Id).Updates(activity).Error
	if err != nil {
		return activity, err
	}

	return activity, nil
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

func (repository *ActivityRepositoryImpl) FindByIdInvest(Id int) (domain.InvestsActivity, error) {
	var dataInvest domain.InvestsActivity
	err := repository.db.First(&dataInvest, Id).Error
	if err != nil {
		return dataInvest, err
	}

	return dataInvest, nil
}

func (repository *ActivityRepositoryImpl) UpdateInvest(activity domain.InvestsActivity) (domain.InvestsActivity, error) {
	err := repository.db.Model(domain.InvestsActivity{}).Where("id = ?", activity.Id).Updates(activity).Error
	if err != nil {
		return activity, err
	}

	return activity, nil
}

func (repository *ActivityRepositoryImpl) DeleteInvest(Id int) error {
	//TODO implement me
	panic("implement me")
}

func (repository *ActivityRepositoryImpl) FindAllInvest(param activities.ActivityQueryParam) ([]domain.InvestsActivity, error) {
	//TODO implement me
	panic("implement me")
}
