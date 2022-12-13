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

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (userService UserServiceImpl) Create(user models.User) models.User {
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		panic(err)
	}
	user.Password = hash
	createdUser := userService.UserRepository.Save(user)
	return createdUser
}

func (userService UserServiceImpl) Update() {
	//TODO implement me
	panic("implement me")
}

func (userService UserServiceImpl) Delete() {
	//TODO implement me
	panic("implement me")
}

func (userService UserServiceImpl) FindById() {
	//TODO implement me
	panic("implement me")
}

func (userService UserServiceImpl) FindAll() []models.User {
	return userService.UserRepository.FindAll()
}
