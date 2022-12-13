package repository

import "github.com/tamahavik/tasker-api/pkg/models"

type AuthenticationRepository interface {
	Login(auth string) models.User
}
