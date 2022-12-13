package services

import (
	"github.com/google/uuid"
	"github.com/tamahavik/tasker-api/pkg/models"
	"github.com/tamahavik/tasker-api/pkg/repository"
	"time"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (userService UserServiceImpl) Create(user models.User) models.User {
	user.Id = uuid.New().String()
	user.CreatedAt = time.Now()
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
