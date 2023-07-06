package user

import "rumahdermawan/backedn-rdi/model/domain"

type UserRepository interface {
	SaveUser(user domain.User) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	FindByID(ID uint) (domain.User, error)
}
