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

type TableController struct {
	tableService service.ITableService
}

func NewTableController(
	tableService service.ITableService,
) *TableController {
	return &TableController{
		tableService: tableService,
	}
}

func (mc *TableController) GetListTable(c *gin.Context) {
	var queryParams rq.GetListTableRequest
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		c.Error(err)
		return
	}

	result, errRS := mc.tableService.GetListTable(&queryParams)
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
		total, er := mc.tableService.GetTotalTable(&queryParams)
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

func (mc *TableController) CreateTable(c *gin.Context) {
	var queryParams rq.CreateTableRequest

	if err := c.ShouldBindJSON(&queryParams); err != nil {
		c.Error(err)
		return
	}

	// Validate username, password
	validate := validator.New()
	if err := validate.Struct(queryParams); err != nil {
		c.Error(err)
		return
	}

	result, er := mc.tableService.CreateTable(&queryParams)
	if er != nil {
		c.Error(er)
		return
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        result,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (mc *TableController) UpdateTable(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	var queryParams rq.UpdateTableRequest

	if err := c.ShouldBindJSON(&queryParams); err != nil {
		c.Error(err)
		return
	}

	// Validate username, password
	validate := validator.New()
	if err := validate.Struct(queryParams); err != nil {
		c.Error(err)
		return
	}

	errRS := mc.tableService.UpdateTable(idInt, &queryParams)
	if errRS != nil {
		c.Error(errRS)
		return
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (mc *TableController) GetTableByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	result, errRS := mc.tableService.GetTableByID(idInt)
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

func (mc *TableController) DeleteTableByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	errRS := mc.tableService.DeleteTableByID(idInt)
	if errRS != nil {
		c.Error(errRS)
		return
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}
