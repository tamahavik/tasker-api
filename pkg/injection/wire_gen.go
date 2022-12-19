// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injection

import (
	"github.com/google/wire"
	"github.com/tamahavik/tasker-api/pkg/controllers"
	"github.com/tamahavik/tasker-api/pkg/repository"
	"github.com/tamahavik/tasker-api/pkg/services"
	"gorm.io/gorm"
)

// Injectors from injector.go:

func InstanceUser(db *gorm.DB) controllers.UserController {
	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	return userController
}

func InstanceAuthentication(db *gorm.DB) controllers.AuthenticationController {
	authenticationRepository := repository.NewAuthenticationRepository(db)
	authenticationService := services.NewAuthenticationService(authenticationRepository)
	authenticationController := controllers.NewAuthenticationController(authenticationService)
	return authenticationController
}

// injector.go:

var (
	provideUser           wire.ProviderSet = wire.NewSet(repository.NewUserRepository, services.NewUserService, controllers.NewUserController)
	provideAuthentication wire.ProviderSet = wire.NewSet(repository.NewAuthenticationRepository, services.NewAuthenticationService, controllers.NewAuthenticationController)
)
