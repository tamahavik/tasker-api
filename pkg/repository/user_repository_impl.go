package repository

import (
	"github.com/tamahavik/tasker-api/pkg/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repository *UserRepositoryImpl) Save(user models.User) models.User {
	repository.DB.Create(&user)
	return user
}

func (repository *UserRepositoryImpl) Update(user models.User) models.User {
	panic("Implement me")
}
func (repository *UserRepositoryImpl) Delete(id string) {
	panic("Implement me")
}
func (repository *UserRepositoryImpl) FindById(id string) models.User {
	panic("Implement me")
}
func (repository *UserRepositoryImpl) FindAll() []models.User {
	var users []models.User
	repository.DB.Find(&users)
	return users
}
