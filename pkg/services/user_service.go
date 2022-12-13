package services

import (
	"github.com/tamahavik/tasker-api/pkg/models"
)

type UserService interface {
	Create(user models.User) models.User
	Update()
	Delete()
	FindById()
	FindAll() []models.User
}
