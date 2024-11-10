package user

import (
	"tranvancu185/vey-pos-ws/internal/middlewares"
	"tranvancu185/vey-pos-ws/internal/uconst"
	"tranvancu185/vey-pos-ws/internal/uconst/messagecode"
	"tranvancu185/vey-pos-ws/internal/wire"
	"tranvancu185/vey-pos-ws/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

func (ur *AuthRouter) InitAuthRouter(c *gin.Context, Router *gin.RouterGroup) {

	authController, err := wire.InitAuthRouterHandler()
	if err != nil {
		response.ErrorResponse(c, response.ParamsResponse{
			Status:      response.SatusInternalError,
			MessageCode: messagecode.CODE_INTERNAL_ERR,
		})
	}
	// Public route
	AuthPublicRoute := Router.Group("/auth")
	{
		AuthPublicRoute.POST("/login", authController.Login)
	}

	AuthPrivateRoute := Router.Group("/auth")
	{
		AuthPrivateRoute.POST("/register", middlewares.AuthMiddleware(uconst.USER_ROLEID_ADMIN), authController.Register)
	}
}
