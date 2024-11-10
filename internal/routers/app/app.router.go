package app

import (
	"tranvancu185/vey-pos-ws/internal/uconst/messagecode"
	"tranvancu185/vey-pos-ws/internal/wire"
	"tranvancu185/vey-pos-ws/pkg/response"

	"github.com/gin-gonic/gin"
)

type MyAppRouter struct{}

func (er *MyAppRouter) InitAppRouter(c *gin.Context, Router *gin.RouterGroup) {

	appController, err := wire.InitAppRouterHandler()
	if err != nil {
		response.ErrorResponse(c, response.ParamsResponse{
			Status:      response.SatusInternalError,
			MessageCode: messagecode.CODE_INTERNAL_ERR,
		})
	}
	// Public route
	AppPublicRoute := Router.Group("/app")
	{
		AppPublicRoute.GET("/version", appController.GetVersion)
		AppPublicRoute.POST("/info", appController.SetAppInfo)
	}

	ImagePublicRoute := Router.Group("/images")
	{
		ImagePublicRoute.GET("/:filename", appController.GetImage)
	}

	UploadFileRoute := Router.Group("/upload")
	{
		UploadFileRoute.POST("/file", appController.UploadFile)
		UploadFileRoute.POST("/image", appController.UploadImage)
	}

}
