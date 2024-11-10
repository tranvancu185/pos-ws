package controller

import (
	"strconv"
	"tranvancu185/vey-pos-ws/internal/model/rq"
	"tranvancu185/vey-pos-ws/internal/model/rs"
	"tranvancu185/vey-pos-ws/internal/service"
	"tranvancu185/vey-pos-ws/internal/uconst/messagecode"
	"tranvancu185/vey-pos-ws/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func (mc *CategoryController) GetListCategory(c *gin.Context) {
	var queryParams rq.GetListCategoryRequest
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		c.Error(err)
		return
	}

	result, errRS := mc.categoryService.GetListCategories(&queryParams)
	if errRS != nil {
		c.Error(errRS)
		return
	}

	data := rs.GetListResponse{
		Page:     queryParams.Page,
		PageSize: queryParams.PageSize,
		Rows:     result,
	}

	if queryParams.Total != 0 {
		total, er := mc.categoryService.GetTotalCategories(&queryParams)
		if er != nil {
			c.Error(er)
			return
		}
		data.Total = total
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        data,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (mc *CategoryController) CreateCategory(c *gin.Context) {
	var params rq.CreateCategoryRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		c.Error(err)
		return
	}

	// Validate username, password
	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		c.Error(err)
		return
	}

	id, err := mc.categoryService.CreateCategory(&params)
	if err != nil {
		c.Error(err)
		return
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        id,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (mc *CategoryController) UpdateCategory(c *gin.Context) {
	var params rq.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		c.Error(err)
		return
	}

	// Validate username, password
	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		c.Error(err)
		return
	}

	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	err = mc.categoryService.UpdateCategory(idInt, &params)
	if err != nil {
		c.Error(err)
		return
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (mc *CategoryController) GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	result, err := mc.categoryService.GetCategoryByID(idInt)
	if err != nil {
		c.Error(err)
		return
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        result,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (cc *CategoryController) DeleteCategoryByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	errR := cc.categoryService.DeleteCategoryByID(idInt)
	if errR != nil {
		c.Error(errR)
		return
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}
