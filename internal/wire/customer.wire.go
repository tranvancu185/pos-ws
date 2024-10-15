//go:build wireinject

package wire

import (
	"tranvancu185/vey-pos-ws/internal/controller"
	"tranvancu185/vey-pos-ws/internal/repo"
	"tranvancu185/vey-pos-ws/internal/service"

	"github.com/google/wire"
)

func InitCustomerRouterHandler() (*controller.CustomerController, error) {
	wire.Build(
		repo.NewCustomerRepo,
		service.NewCustomerService,
		controller.NewCustomerController,
	)

	return new(controller.CustomerController), nil
}
