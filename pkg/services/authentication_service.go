package services

import (
	"errors"
	"github.com/tamahavik/tasker-api/pkg/dto"
	"github.com/tamahavik/tasker-api/pkg/models"
	"github.com/tamahavik/tasker-api/pkg/repository"
	"github.com/tamahavik/tasker-api/pkg/utils"
)

type AuthenticationService interface {
	Login(request *dto.LoginRequest) error
	Register(request *dto.RegisterRequest) (models.User, error)
}

type authenticationServiceImpl struct {
	repository repository.AuthenticationRepository
}

func NewAuthenticationService(authenticationRepository repository.AuthenticationRepository) AuthenticationService {
	return &authenticationServiceImpl{repository: authenticationRepository}
}

func (a *authenticationServiceImpl) Login(d *dto.LoginRequest) error {
	user := a.repository.FindByEmailAndUsername(d.Auth, d.Auth)
	if user.ID == 0 {
		return errors.New("credentials not valid")
	}
	if !utils.ComparePassword(d.Password, user.Password) {
		return errors.New("credentials not valid")
	}
	return nil
}

func (a *authenticationServiceImpl) Register(r *dto.RegisterRequest) (models.User, error) {
	user := a.repository.FindByEmailAndUsername(r.Email, r.Username)
	if user.ID != 0 {
		return user, errors.New("email or username has beer registered")
	}

	hash, err := utils.HashPassword(r.Password)
	if err != nil {
		return user, err
	}

	u := models.User{
		Username: r.Username,
		Email:    r.Email,
		Password: hash,
		Name:     r.Name,
	}

	a.repository.Register(&u)
	return u, nil
}
