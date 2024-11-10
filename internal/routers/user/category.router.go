package user

import (
	"tranvancu185/vey-pos-ws/internal/middlewares"
	"tranvancu185/vey-pos-ws/internal/uconst"
	"tranvancu185/vey-pos-ws/internal/uconst/messagecode"
	"tranvancu185/vey-pos-ws/internal/wire"
	"tranvancu185/vey-pos-ws/pkg/response"

	"github.com/gin-gonic/gin"
)

type CategoryRouter struct{}

func (tr *CategoryRouter) InitTableRouter(c *gin.Context, Router *gin.RouterGroup) {
	categoryController, error := wire.InitTableRouterHandler()
	if error != nil {
		response.ErrorResponse(c, response.ParamsResponse{
			Status:      response.SatusInternalError,
			MessageCode: messagecode.CODE_INTERNAL_ERR,
		})
	}
	// Public route
	TablePublicRoute := Router.Group("/table")
	{
		TablePublicRoute.GET("/", categoryController.GetListTable)
	}
	// Private route
	TablePrivateRoute := Router.Group("/table")
	{
		TablePrivateRoute.Use(middlewares.AuthMiddleware(uconst.FLAG_BYPASS_ROLE))

		TablePrivateRoute.POST("/create", categoryController.CreateTable)
		TablePrivateRoute.GET("/detail/:id", categoryController.GetTableByID)
		TablePrivateRoute.PUT("/update/:id", categoryController.UpdateTable)
		TablePrivateRoute.DELETE("/delete/:id", categoryController.DeleteTableByID)
	}

}
