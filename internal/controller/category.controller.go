package controller

import (
	"tranvancu185/vey-pos-ws/internal/constants/messagecode"
	"tranvancu185/vey-pos-ws/internal/service"
	"tranvancu185/vey-pos-ws/pkg/response"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService service.ICategoryService
}

func NewCategoryController(
	categoryService service.ICategoryService,
) *CategoryController {
	return &CategoryController{
		categoryService: categoryService,
	}
}

func (cas *CategoryController) GetInFo(c *gin.Context) {
	result := cas.categoryService.GetInfo()
	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        result,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}
