package services

import (
	"errors"
	"fmt"
	"github.com/tamahavik/tasker-api/pkg/repository"
	"github.com/tamahavik/tasker-api/pkg/utils"
)

type AuthenticationService interface {
	Login(auth string, password string) error
}

type authenticationServiceImpl struct {
	Repository repository.AuthenticationRepository
}

func NewAuthenticationService(authenticationRepository repository.AuthenticationRepository) *authenticationServiceImpl {
	return &authenticationServiceImpl{Repository: authenticationRepository}
}

func (a *authenticationServiceImpl) Login(auth string, password string) error {
	user := a.Repository.Login(auth)
	if user.ID == 0 {
		return errors.New("credentials not valid")
	}

	fmt.Println(utils.ComparePassword(user.Password, password))
	if !utils.ComparePassword(password, user.Password) {
		return errors.New("credentials not valid")
	}
	return nil
}
