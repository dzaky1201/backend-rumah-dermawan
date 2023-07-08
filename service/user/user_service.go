package user

import (
	"rumahdermawan/backedn-rdi/model/web/user"
)

type UserService interface {
	SaveUser(request user.UserCreateRequest, token string) (user.UserResponse, error)
	FindByEmail(request user.UserLoginRequest, token string) (user.UserResponse, error)
	FindById(Id uint) user.UserResponse
}
