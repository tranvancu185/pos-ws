package user

import (
	"tranvancu185/vey-pos-ws/internal/constants"
	"tranvancu185/vey-pos-ws/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct{}

func (pr *ProductRouter) InitProductRouter(c *gin.Context, Router *gin.RouterGroup) {

	// productController, error := wire.InitProductRouterHandler()
	// Public router
	productPublicRouter := Router.Group("/product")
	{
		productPublicRouter.GET("/search")
		productPublicRouter.GET("/detail/:id")
	}

	// Private router
	productPrivateRouter := Router.Group("/product", middlewares.AuthMiddleware(constants.USER_ROLEID_MANAGER))
	{
		productPrivateRouter.POST("/")
		// productPrivateRouter.POST("/", productController.CreateProduct)
		// productPrivateRouter.PUT("/:id", productController.UpdateProduct)
		// productPrivateRouter.DELETE("/:id", productController.DeleteProduct)
	}

}
