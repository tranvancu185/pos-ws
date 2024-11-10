package controller

import (
	"tranvancu185/vey-pos-ws/internal/service"
	"tranvancu185/vey-pos-ws/internal/uconst/messagecode"
	"tranvancu185/vey-pos-ws/pkg/response"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService service.IOrderService
}

func NewOrderController(
	orderService service.IOrderService,
) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

func (oc *OrderController) GetInFo(c *gin.Context) {
	result := oc.orderService.GetInfo()
	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        result,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}
