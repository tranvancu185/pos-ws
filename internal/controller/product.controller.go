package controller

import (
	"tranvancu185/vey-pos-ws/internal/constants/messagecode"
	"tranvancu185/vey-pos-ws/internal/service"
	"tranvancu185/vey-pos-ws/pkg/response"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService service.IProductService
}

func NewProductController(
	productService service.IProductService,
) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (pc *ProductController) GetInFo(c *gin.Context) {
	result := pc.productService.GetInFo()
	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        result,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}
