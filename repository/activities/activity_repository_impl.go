package activities

import (
	"fmt"
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

	err := repository.db.Preload("YearPeriod").Where("id = ?", Id).Find(&dataOperation).Error
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
	err := repository.db.Delete(&domain.OperationActivity{}, Id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository *ActivityRepositoryImpl) FindAllOperation(param activities.ActivityQueryParam) ([]domain.OperationActivity, int, int, error) {
	var operationList []domain.OperationActivity
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

	if param.Description != "" {
		desc := fmt.Sprintf("'''%s'''", param.Description)
		err := repository.db.Table("operation_activities").Preload("YearPeriod").Where("description @@ to_tsquery(?)", desc).Offset(offset).Limit(param.Limit).Order("updated_at desc").Find(&operationList).Error
		if err != nil {
			return operationList, 0, 0, err
		}
	} else {
		err := repository.db.Preload("YearPeriod").Limit(param.Limit).Offset(offset).Order("updated_at desc").Find(&operationList).Error
		if err != nil {
			return operationList, 0, 0, err
		}
	}
	totalCount := int64(0)
	repository.db.Table("operation_activities").Count(&totalCount)

	return operationList, int(totalCount), offset, nil
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
	err := repository.db.Preload("YearPeriod").Where("id = ?", Id).Find(&dataFunding).Error
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
	err := repository.db.Delete(&domain.FundingActivity{}, Id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository *ActivityRepositoryImpl) FindAllFunding(param activities.ActivityQueryParam) ([]domain.FundingActivity, int, int, error) {
	var fundingList []domain.FundingActivity
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

	if param.Description != "" {
		desc := fmt.Sprintf("'''%s'''", param.Description)
		err := repository.db.Table("funding_activities").Preload("YearPeriod").Where("description @@ to_tsquery(?)", desc).Offset(offset).Limit(param.Limit).Find(&fundingList).Error
		if err != nil {
			return fundingList, 0, 0, err
		}
	} else {
		err := repository.db.Preload("YearPeriod").Limit(param.Limit).Offset(offset).Find(&fundingList).Error
		if err != nil {
			return fundingList, 0, 0, err
		}
	}

	totalCount := int64(0)
	repository.db.Table("funding_activities").Count(&totalCount)

	return fundingList, int(totalCount), offset, nil
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
	err := repository.db.Preload("YearPeriod").Where("id = ?", Id).Find(&dataInvest).Error
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
	err := repository.db.Delete(&domain.InvestsActivity{}, Id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository *ActivityRepositoryImpl) FindAllInvest(param activities.ActivityQueryParam) ([]domain.InvestsActivity, int, int, error) {
	var investList []domain.InvestsActivity
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

	if param.Description != "" {
		desc := fmt.Sprintf("'''%s'''", param.Description)
		err := repository.db.Table("invests_activities").Preload("YearPeriod").Where("description @@ to_tsquery(?)", desc).Offset(offset).Limit(param.Limit).Find(&investList).Error
		if err != nil {
			return investList, 0, 0, err
		}
	} else {
		err := repository.db.Preload("YearPeriod").Limit(param.Limit).Offset(offset).Find(&investList).Error
		if err != nil {
			return investList, 0, 0, err
		}
	}
	totalCount := int64(0)
	repository.db.Table("invests_activities").Count(&totalCount)

	return investList, int(totalCount), offset, nil
}

func (repository *ActivityRepositoryImpl) ReportActivity(year string) ([]domain.ReportActivity, error) {
	var allData []domain.ReportActivity
	errData := repository.db.Raw("select info_period ->> 'Month' as month, sum(case when type_transaction = 'credit' then amount * -1 when type_transaction = 'debit' then amount end) as total from (select * from year_periods join operation_activities oa on year_periods.id = oa.year_period_id union select * from year_periods join invests_activities ia on year_periods.id = ia.year_period_id union select * from year_periods join funding_activities fa on year_periods.id = fa.year_period_id) as all_data where info_period ->> 'Year' = ? group by info_period ->> 'Month', year_period_id order by year_period_id", year).Scan(&allData).Error

	if errData != nil {
		return []domain.ReportActivity{}, errData
	}

	return allData, nil
}
