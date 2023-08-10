package user

import (
	"errors"
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

func (service *UserServiceImpl) SaveUser(request user2.UserCreateRequest, token string) (user2.UserResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return user2.UserResponse{}, err
	}

	passwordHash, errPass := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return user2.UserResponse{}, errPass
	}

	userReq := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(passwordHash),
	}

	userRes, errReg := service.repository.SaveUser(userReq)

	if errReg != nil {
		return user2.UserResponse{}, errors.New("email sudah ada")
	}

	return helper.ToUserResponse(userRes, token), nil

}

func (service *UserServiceImpl) FindByEmail(request user2.UserLoginRequest, token string) (user2.UserResponse, error) {
	userLogin, err := service.repository.FindByEmail(request.Email)
	if err != nil {
		return user2.UserResponse{}, err
	}

	if userLogin.Id == 0 {
		return user2.UserResponse{}, errors.New("no User Found on that email")
	}

	errPass := bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(request.Password))
	if errPass != nil {
		return user2.UserResponse{}, errors.New("password salah")
	}

	return helper.ToUserResponse(userLogin, token), nil
}

func (service *UserServiceImpl) FindById(Id uint) user2.UserResponse {
	userReq, err := service.repository.FindByID(Id)
	helper.PanicIfError(err)
	return helper.ToUserResponse(userReq, "")
}
