package user

import (
	"tranvancu185/vey-pos-ws/internal/constants"
	"tranvancu185/vey-pos-ws/internal/constants/messagecode"
	"tranvancu185/vey-pos-ws/internal/middlewares"
	"tranvancu185/vey-pos-ws/internal/wire"
	"tranvancu185/vey-pos-ws/pkg/response"

	"github.com/gin-gonic/gin"
)

type TableRouter struct{}

func (tr *TableRouter) InitTableRouter(c *gin.Context, Router *gin.RouterGroup) {
	tableController, error := wire.InitTableRouterHandler()
	if error != nil {
		response.ErrorResponse(c, response.ParamsResponse{
			Status:      response.SatusInternalError,
			MessageCode: messagecode.CODE_INTERNAL_ERR,
		})
	}
	// Public route
	TablePublicRoute := Router.Group("/table")
	{
		TablePublicRoute.GET("/", tableController.GetListTable)
	}
	// Private route
	TablePrivateRoute := Router.Group("/table")
	{
		TablePrivateRoute.Use(middlewares.AuthMiddleware(constants.FLAG_BYPASS_ROLE))

		TablePrivateRoute.POST("/create", tableController.CreateTable)
		TablePrivateRoute.GET("/detail/:id", tableController.GetTableByID)
		TablePrivateRoute.PUT("/update/:id", tableController.UpdateTable)
		TablePrivateRoute.DELETE("/delete/:id", tableController.DeleteTableByID)
	}

}
