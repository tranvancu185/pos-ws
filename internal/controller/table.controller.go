package controller

import (
	"tranvancu185/vey-pos-ws/internal/constants/messagecode"
	"tranvancu185/vey-pos-ws/internal/service"
	"tranvancu185/vey-pos-ws/pkg/request"
	"tranvancu185/vey-pos-ws/pkg/response"

	"github.com/gin-gonic/gin"
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
	var queryParams request.GetListTableRequest

	if err := c.ShouldBindQuery(&queryParams); err != nil {
		c.Error(err)
		return
	}

	result, errRS := mc.tableService.GetListTable(&queryParams)
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
