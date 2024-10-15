//go:build wireinject

package wire

import (
	"tranvancu185/vey-pos-ws/internal/controller"
	"tranvancu185/vey-pos-ws/internal/repo"
	"tranvancu185/vey-pos-ws/internal/service"

	"github.com/google/wire"
)

func InitPaymentRouterHandler() (*controller.PaymentController, error) {
	wire.Build(
		repo.NewPaymentRepo,
		service.NewPaymentService,
		controller.NewPaymentController,
	)

	return new(controller.PaymentController), nil
}
