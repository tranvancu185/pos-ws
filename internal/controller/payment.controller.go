package controller

import (
	"tranvancu185/vey-pos-ws/internal/constants/messagecode"
	"tranvancu185/vey-pos-ws/internal/service"
	"tranvancu185/vey-pos-ws/pkg/response"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	paymentService service.IPaymentService
}

func NewPaymentController(
	paymentService service.IPaymentService,
) *PaymentController {
	return &PaymentController{
		paymentService: paymentService,
	}
}

func (pc *PaymentController) GetInFo(c *gin.Context) {
	result := pc.paymentService.GetInfo()
	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        result,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}
