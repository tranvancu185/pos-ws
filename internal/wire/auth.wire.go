//go:build wireinject

package wire

import (
	"tranvancu185/vey-pos-ws/internal/controller"
	"tranvancu185/vey-pos-ws/internal/repo"
	"tranvancu185/vey-pos-ws/internal/service"

	"github.com/google/wire"
)

func InitAuthRouterHandler() (*controller.AuthController, error) {
	wire.Build(
		repo.NewAuthRepo,
		repo.NewUserRepo,
		service.NeuAuthService,
		controller.NewAuthController,
	)

	return new(controller.AuthController), nil
}
