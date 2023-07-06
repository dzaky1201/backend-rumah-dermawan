package user

import (
	"gorm.io/gorm"
	"rumahdermawan/backedn-rdi/model/domain"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (repository *UserRepositoryImpl) SaveUser(user domain.User) (domain.User, error) {
	err := repository.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	err := repository.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindByID(ID uint) (domain.User, error) {
	var user domain.User
	err := repository.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
