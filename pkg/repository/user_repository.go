package repository

import (
	"github.com/tamahavik/tasker-api/pkg/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user models.User) models.User
	Update(user models.User) models.User
	Delete(id string)
	FindById(id string) models.User
	FindAll() []models.User
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepositoryImpl {
	return &userRepositoryImpl{
		db: db,
	}
}

func (repository *userRepositoryImpl) Save(user models.User) models.User {
	repository.db.Create(&user)
	return user
}

func (repository *userRepositoryImpl) Update(user models.User) models.User {
	panic("Implement me")
}
func (repository *userRepositoryImpl) Delete(id string) {
	panic("Implement me")
}
func (repository *userRepositoryImpl) FindById(id string) models.User {
	panic("Implement me")
}
func (repository *userRepositoryImpl) FindAll() []models.User {
	var users []models.User
	repository.db.Find(&users)
	return users
}
