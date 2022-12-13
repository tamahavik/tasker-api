package repository

import (
	"github.com/tamahavik/tasker-api/pkg/models"
)

type UserRepository interface {
	Save(user models.User) models.User
	Update(user models.User) models.User
	Delete(id string)
	FindById(id string) models.User
	FindAll() []models.User
}
