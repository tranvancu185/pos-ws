//go:build wireinject

package wire

import (
	"tranvancu185/vey-pos-ws/internal/controller"
	"tranvancu185/vey-pos-ws/internal/repo"
	"tranvancu185/vey-pos-ws/internal/service"

	"github.com/google/wire"
)

func InitProductRouterHandler() (*controller.ProductController, error) {
	wire.Build(
		repo.NewProductRepo,
		service.NewProductService,
		controller.NewProductController,
	)

	return new(controller.ProductController), nil
}
