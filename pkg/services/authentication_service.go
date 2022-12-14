package services

import (
	"errors"
	"github.com/tamahavik/tasker-api/pkg/repository"
	"github.com/tamahavik/tasker-api/pkg/utils"
)

type AuthenticationService interface {
	Login(auth string, password string) error
}

type authenticationServiceImpl struct {
	repository repository.AuthenticationRepository
}

func NewAuthenticationService(authenticationRepository repository.AuthenticationRepository) AuthenticationService {
	return &authenticationServiceImpl{repository: authenticationRepository}
}

func (a *authenticationServiceImpl) Login(auth string, password string) error {
	user := a.repository.Login(auth)
	if user.ID == 0 {
		return errors.New("credentials not valid")
	}
	if !utils.ComparePassword(password, user.Password) {
		return errors.New("credentials not valid")
	}
	return nil
}
