//go:build wireinject
// +build wireinject

package injection

import (
	"github.com/google/wire"
	"github.com/tamahavik/tasker-api/pkg/repository"
	"github.com/tamahavik/tasker-api/pkg/services"
	"gorm.io/gorm"
)

func UserInstance(db *gorm.DB) *services.UserService {
	wire.Build(repository.NewUserRepository)
	return nil
}
