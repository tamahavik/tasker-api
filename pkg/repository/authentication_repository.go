package repository

import (
	"github.com/tamahavik/tasker-api/pkg/models"
	"gorm.io/gorm"
)

type AuthenticationRepository interface {
	Login(auth string) models.User
}

type authenticationRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthenticationRepository(db *gorm.DB) AuthenticationRepository {
	return &authenticationRepositoryImpl{db: db}
}

func (a *authenticationRepositoryImpl) Login(auth string) models.User {
	var user models.User
	a.db.Where("username = ? OR email = ?", auth, auth).Find(&user)
	return user
}
