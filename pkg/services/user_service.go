package services

import (
	"github.com/tamahavik/tasker-api/pkg/models"
	"github.com/tamahavik/tasker-api/pkg/repository"
	"github.com/tamahavik/tasker-api/pkg/utils"
)

type UserService interface {
	Create(user models.User) models.User
	Update()
	Delete()
	FindById()
	FindAll() []models.User
}

type userServiceImpl struct {
	repository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userServiceImpl{repository: userRepository}
}

func (userService userServiceImpl) Create(user models.User) models.User {
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		panic(err)
	}
	user.Password = hash
	createdUser := userService.repository.Save(user)
	return createdUser
}

func (userService userServiceImpl) Update() {
	//TODO implement me
	panic("implement me")
}

func (userService userServiceImpl) Delete() {
	//TODO implement me
	panic("implement me")
}

func (userService userServiceImpl) FindById() {
	//TODO implement me
	panic("implement me")
}

func (userService userServiceImpl) FindAll() []models.User {
	return userService.repository.FindAll()
}
