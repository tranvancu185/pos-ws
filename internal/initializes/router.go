package initializes

import (
	"time"
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/internal/middlewares"
	"tranvancu185/vey-pos-ws/internal/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var router *gin.Engine
	var c *gin.Context
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		router = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
	}

	router = ExeStatic(router)

	// middleware
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(gin.Recovery())
	// cors
	corsConfig := cors.Config{
		AllowAllOrigins: true,                                                         // Cho phép tất cả các origin
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}, // Các phương thức được phép
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Authorization",
			"app-version",
		}, // Các header được phép
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(corsConfig))

	router.Use(middlewares.ErrorHandler())

	appRouter := routers.RouterGroupApp.App
	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	MainGroup := router.Group("/api")
	{
		MainGroup.GET("/checkStatus") // tracking status of server
		appRouter.InitAppRouter(c, MainGroup)
	}
	{
		userRouter.InitAuthRouter(c, MainGroup)
		userRouter.InitUserRouter(c, MainGroup)
		userRouter.InitProductRouter(c, MainGroup)
		userRouter.InitTableRouter(c, MainGroup)
	}
	{
		managerRouter.InitUserRouter(MainGroup)
		managerRouter.InitProductRouter(MainGroup)
	}

	return router
}

func ExeStatic(router *gin.Engine) *gin.Engine {
	// router.Use(static.Serve("/", static.LocalFile(global.Config.Path.PathWeb, true)))
	router.Use(static.Serve("/", static.LocalFile("", true)))
	return router
}
