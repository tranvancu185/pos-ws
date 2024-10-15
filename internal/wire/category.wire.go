//go:build wireinject

package wire

import (
	"tranvancu185/vey-pos-ws/internal/controller"
	"tranvancu185/vey-pos-ws/internal/repo"
	"tranvancu185/vey-pos-ws/internal/service"

	"github.com/google/wire"
)

func InitCategoryRouterHandler() (*controller.CategoryController, error) {
	wire.Build(
		repo.NewCategoryRepo,
		service.NewCategoryService,
		controller.NewCategoryController,
	)

	return new(controller.CategoryController), nil
}
