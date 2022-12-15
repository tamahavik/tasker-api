//go:build wireinject
// +build wireinject

package injection

import (
	"github.com/google/wire"
	"github.com/tamahavik/tasker-api/pkg/controllers"
	"github.com/tamahavik/tasker-api/pkg/repository"
	"github.com/tamahavik/tasker-api/pkg/services"
	"github.com/tamahavik/tasker-api/pkg/utils"
	"gorm.io/gorm"
)

var (
	provideConfig wire.ProviderSet = wire.NewSet(utils.NewConfiguration)
	provideUser   wire.ProviderSet = wire.NewSet(
		repository.NewUserRepository,
		services.NewUserService,
		controllers.NewUserController)
	provideAuthentication wire.ProviderSet = wire.NewSet(
		repository.NewAuthenticationRepository,
		services.NewAuthenticationService,
		controllers.NewAuthenticationController)
)

func InstanceConfig() utils.Configuration {
	wire.Build(provideConfig)
	return nil
}

func InstanceUser(db *gorm.DB) controllers.UserController {
	wire.Build(provideUser)
	return nil
}

func InstanceAuthentication(db *gorm.DB) controllers.AuthenticationController {
	wire.Build(provideAuthentication)
	return nil
}
