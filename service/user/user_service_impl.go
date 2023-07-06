package user

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"rumahdermawan/backedn-rdi/helper"
	"rumahdermawan/backedn-rdi/model/domain"
	user2 "rumahdermawan/backedn-rdi/model/web/user"
	"rumahdermawan/backedn-rdi/repository/user"
)

type UserServiceImpl struct {
	repository user.UserRepository
	Validate   *validator.Validate
}

func NewUserService(repository user.UserRepository, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{repository: repository, Validate: validate}
}

func (service *UserServiceImpl) SaveUser(request user2.UserCreateRequest, token string) user2.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	passwordHash, errPass := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	helper.PanicIfError(errPass)

	userReq := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(passwordHash),
	}

	userRes, _ := service.repository.SaveUser(userReq)

	return helper.ToUserResponse(userRes, token)

}

func (service *UserServiceImpl) FindByEmail(request user2.UserLoginRequest, token string) user2.UserResponse {
	userLogin, err := service.repository.FindByEmail(request.Email)
	helper.PanicIfError(err)

	errPass := bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(request.Password))
	helper.PanicIfError(errPass)

	return helper.ToUserResponse(userLogin, token)
}

func (service *UserServiceImpl) FindById(Id uint) user2.UserResponse {
	userReq, err := service.repository.FindByID(Id)
	helper.PanicIfError(err)
	return helper.ToUserResponse(userReq, "")
}
