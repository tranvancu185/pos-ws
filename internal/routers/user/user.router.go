package user

import (
	"tranvancu185/vey-pos-ws/internal/constants"
	"tranvancu185/vey-pos-ws/internal/constants/messagecode"
	"tranvancu185/vey-pos-ws/internal/middlewares"
	"tranvancu185/vey-pos-ws/internal/wire"
	"tranvancu185/vey-pos-ws/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(c *gin.Context, Router *gin.RouterGroup) {

	userController, error := wire.InitUserRouterHandler()
	if error != nil {
		response.ErrorResponse(c, response.ParamsResponse{
			Status:      response.SatusInternalError,
			MessageCode: messagecode.CODE_INTERNAL_ERR,
		})
	}
	// Public route
	UserPublicRoute := Router.Group("/user")
	{
		UserPublicRoute.GET("/", userController.GetListUsers)
	}
	// Private route
	UserPrivateRoute := Router.Group("/user")
	{
		UserPrivateRoute.Use(middlewares.AuthMiddleware(constants.FLAG_BYPASS_ROLE))

		// UserPrivateRoute.POST("/", userController.)
		UserPrivateRoute.GET("/profile", userController.GetProfile)
		UserPrivateRoute.GET("/detail/:id", userController.GetUserByID)
		UserPrivateRoute.PUT("/:id/update-password", userController.UpdatePassword)
		UserPrivateRoute.PUT("/:id/update-profile", userController.UpdateProfile)
		UserPrivateRoute.PUT("/:id/update-avatar", userController.UpdateAvatar)
	}
}
