package controller

import (
	"path/filepath"
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/internal/constants/messagecode"
	"tranvancu185/vey-pos-ws/internal/service"
	"tranvancu185/vey-pos-ws/pkg/electron"
	"tranvancu185/vey-pos-ws/pkg/request"
	"tranvancu185/vey-pos-ws/pkg/response"

	"github.com/gin-gonic/gin"
)

type AppController struct {
	appService service.IAppService
}

func NewAppController(
	appService service.IAppService,
) *AppController {
	return &AppController{
		appService: appService,
	}
}

func (ac *AppController) GetListApp(c *gin.Context) {
	var queryParams request.GetListAppRequest

	if err := c.ShouldBindQuery(&queryParams); err != nil {
		c.Error(err)
		return
	}

	result, err := ac.appService.GetListApp(queryParams)
	if err != nil {
		c.Error(err)
		return
	}

	data := response.GetListResponse{
		Page:     queryParams.Page,
		PageSize: queryParams.PageSize,
		Total:    queryParams.Total,
		Rows:     result,
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        data,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (ac *AppController) GetVersion(c *gin.Context) {

	electron.SendElecEvent(electron.ELEC_EVENT_GET_VERSION, nil)

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        nil,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (ac *AppController) SetAppInfo(c *gin.Context) {
	var params request.SetAppRequest

	if err := c.ShouldBindJSON(&params); err != nil {
		c.Error(err)
		return
	}

	id, err := ac.appService.SetAppInfo(params)
	if err != nil {
		c.Error(err)
		return
	}

	data := interface{}(id)

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        data,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (ac *AppController) UpdateAppInfo(c *gin.Context) {
	//TODO: update app info
	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        nil,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (ac *AppController) QuitApp(c *gin.Context) {
	electron.SendElecEvent(electron.ELEC_EVENT_QUIT_APP, nil)

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        nil,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (ac *AppController) GetImage(c *gin.Context) {
	filename := c.Param("filename")
	filePath := filepath.Join(global.Config.Path.PathAvatar, filename)
	c.File(filePath)
}

func (ac *AppController) UploadFile(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.Error(err)
		return
	}

	filePath := filepath.Join(global.Config.Path.PathFile, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.Error(err)
		return
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        nil,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}

func (ac *AppController) UploadImage(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.Error(err)
		return
	}

	filePath := filepath.Join(global.Config.Path.PathImage, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.Error(err)
		return
	}

	response.SuccessResponse(c, response.ParamsResponse{
		Status:      response.StatusCodeSuccess,
		Data:        nil,
		MessageCode: messagecode.CODE_SUCCESS,
	})
}
