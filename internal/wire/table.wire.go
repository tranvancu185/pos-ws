//go:build wireinject

package wire

import (
	"tranvancu185/vey-pos-ws/internal/controller"
	"tranvancu185/vey-pos-ws/internal/repo"
	"tranvancu185/vey-pos-ws/internal/service"

	"github.com/google/wire"
)

func InitTableRouterHandler() (*controller.TableController, error) {
	wire.Build(
		repo.NewTableRepo,
		service.NewTableService,
		controller.NewTableController,
	)

	return new(controller.TableController), nil
}
