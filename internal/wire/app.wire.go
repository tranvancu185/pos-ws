//go:build wireinject

package wire

import (
	"tranvancu185/vey-pos-ws/internal/controller"
	"tranvancu185/vey-pos-ws/internal/repo"
	"tranvancu185/vey-pos-ws/internal/service"

	"github.com/google/wire"
)

func InitAppRouterHandler() (*controller.AppController, error) {
	wire.Build(
		repo.NewAppRepo,
		service.NeuAppService,
		controller.NewAppController,
	)
	return new(controller.AppController), nil
}
