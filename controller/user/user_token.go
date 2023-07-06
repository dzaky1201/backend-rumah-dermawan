package user

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

type UserToken interface {
	GenerateToken(userID uint) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type UserTokenImpl struct {
}

func NewServiceUserToken() *UserTokenImpl {
	return &UserTokenImpl{}
}

func (u *UserTokenImpl) GenerateToken(userID uint) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(os.Getenv("API_SECRET")))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (u *UserTokenImpl) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid Token")
		}

		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
