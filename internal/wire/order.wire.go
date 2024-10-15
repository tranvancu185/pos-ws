//go:build wireinject

package wire

import (
	"tranvancu185/vey-pos-ws/internal/controller"
	"tranvancu185/vey-pos-ws/internal/repo"
	"tranvancu185/vey-pos-ws/internal/service"

	"github.com/google/wire"
)

func InitOrderRouterHandler() (*controller.OrderController, error) {
	wire.Build(
		repo.NewOrderRepo,
		service.NewOrderService,
		controller.NewOrderController,
	)

	return new(controller.OrderController), nil
}
