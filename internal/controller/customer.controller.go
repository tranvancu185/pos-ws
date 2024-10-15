package controller

import (
	"tranvancu185/vey-pos-ws/internal/constants/messagecode"
	"tranvancu185/vey-pos-ws/internal/service"
	"tranvancu185/vey-pos-ws/pkg/response"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	customerService service.ICustomerService
}

func NewCustomerController(
	customerService service.ICustomerService,
) *CustomerController {
	return &CustomerController{
		customerService: customerService,
	}
}

func (cusc *CustomerController) GetInFo(c *gin.Context) {
	result := cusc.customerService.GetInfo()
	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        result,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}
