package repository

import (
	"github.com/tamahavik/tasker-api/pkg/models"
	"gorm.io/gorm"
)

type AuthenticationRepository interface {
	FindByEmailAndUsername(email string, username string) models.User
	Register(user *models.User) *models.User
}

type authenticationRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthenticationRepository(db *gorm.DB) AuthenticationRepository {
	return &authenticationRepositoryImpl{db: db}
}

func (a *authenticationRepositoryImpl) FindByEmailAndUsername(email string, username string) models.User {
	var user models.User
	a.db.Where("username = ? OR email = ?", username, email).Find(&user)
	return user
}

func (a *authenticationRepositoryImpl) Register(user *models.User) *models.User {
	a.db.Create(user)
	return user
}
