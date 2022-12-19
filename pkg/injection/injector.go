//go:build wireinject
// +build wireinject

package injection

import (
	"github.com/google/wire"
	"github.com/tamahavik/tasker-api/pkg/controllers"
	"github.com/tamahavik/tasker-api/pkg/repository"
	"github.com/tamahavik/tasker-api/pkg/services"
	"gorm.io/gorm"
)

var (
	provideUser wire.ProviderSet = wire.NewSet(
		repository.NewUserRepository,
		services.NewUserService,
		controllers.NewUserController)
	provideAuthentication wire.ProviderSet = wire.NewSet(
		repository.NewAuthenticationRepository,
		services.NewAuthenticationService,
		controllers.NewAuthenticationController)
)

func InstanceUser(db *gorm.DB) controllers.UserController {
	wire.Build(provideUser)
	return nil
}

func InstanceAuthentication(db *gorm.DB) controllers.AuthenticationController {
	wire.Build(provideAuthentication)
	return nil
}
