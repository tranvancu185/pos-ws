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

func (mc *CustomerController) GetListCustomer(c *gin.Context) {
	var queryParams rq.GetListCustomerRequest
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		c.Error(err)
		return
	}

	result, errRS := mc.customerService.GetListCustomer(&queryParams)
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
		total, er := mc.customerService.GetTotalCustomer(&queryParams)
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

func (mc *CustomerController) CreateCustomer(c *gin.Context) {
	var params rq.CreateCustomerRequest
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

	id, err := mc.customerService.CreateCustomer(&params)
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

func (mc *CustomerController) UpdateCustomer(c *gin.Context) {
	var params rq.UpdateCustomerRequest
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

	err = mc.customerService.UpdateCustomerByID(idInt, &params)
	if err != nil {
		c.Error(err)
		return
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        nil,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (mc *CustomerController) GetCustomerByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	result, errRS := mc.customerService.GetCustomerByID(idInt)
	if errRS != nil {
		c.Error(errRS)
		return
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        result,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (mc *CustomerController) DeleteCustomerByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	err = mc.customerService.DeleteCustomerByID(idInt)
	if err != nil {
		c.Error(err)
		return
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (mc *CustomerController) GetTotalCustomer(c *gin.Context) {
	var queryParams rq.GetListCustomerRequest
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		c.Error(err)
		return
	}

	total, err := mc.customerService.GetTotalCustomer(&queryParams)
	if err != nil {
		c.Error(err)
		return
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        total,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}
