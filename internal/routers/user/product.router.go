package user

import (
	"tranvancu185/vey-pos-ws/internal/middlewares"
	"tranvancu185/vey-pos-ws/internal/uconst"
	"tranvancu185/vey-pos-ws/internal/uconst/messagecode"
	"tranvancu185/vey-pos-ws/internal/wire"
	"tranvancu185/vey-pos-ws/pkg/response"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct{}

func (pr *ProductRouter) InitProductRouter(c *gin.Context, Router *gin.RouterGroup) {

	productController, error := wire.InitProductRouterHandler()
	if error != nil {
		response.ErrorResponse(c, response.ParamsResponse{
			Status:      response.SatusInternalError,
			MessageCode: messagecode.CODE_INTERNAL_ERR,
		})
	}
	// Public router
	productPublicRouter := Router.Group("/product")
	{
		productPublicRouter.GET("/search")
		productPublicRouter.GET("/detail/:id")
		productPublicRouter.GET("/list", productController.GetListProduct)
	}

	// Private router
	productPrivateRouter := Router.Group("/product")
	{
		productPrivateRouter.Use(middlewares.AuthMiddleware(uconst.FLAG_BYPASS_ROLE))
		productPrivateRouter.POST("/create", productController.CreateProduct)
		productPrivateRouter.GET("/detail/:id", productController.GetProductByID)
		productPrivateRouter.PUT("/update/:id", productController.UpdateProduct)
		productPrivateRouter.DELETE("/delete/:id", productController.DeleteProductByID)
	}

}
