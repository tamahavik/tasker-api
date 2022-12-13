package repository

import (
	"github.com/tamahavik/tasker-api/pkg/models"
	"gorm.io/gorm"
)

type authenticationRepositoryImpl struct {
	*gorm.DB
}

func NewAuthenticationRepository(db *gorm.DB) *authenticationRepositoryImpl {
	return &authenticationRepositoryImpl{DB: db}
}

func (a *authenticationRepositoryImpl) Login(auth string) models.User {
	var user models.User
	a.DB.Where("username = ? OR email = ?", auth, auth).Find(&user)
	return user
}
