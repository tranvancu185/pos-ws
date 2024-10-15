//go:build wireinject

package wire

import (
	"tranvancu185/vey-pos-ws/internal/controller"
	"tranvancu185/vey-pos-ws/internal/repo"
	"tranvancu185/vey-pos-ws/internal/service"

	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepo,
		repo.NewAuthRepo,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
