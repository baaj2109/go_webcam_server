package router

import (
	"net/http"
	"time"

	"github.com/baaj2109/webcam_server/api"
	"github.com/baaj2109/webcam_server/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(engine *gin.Engine) {

	// engine.Use(gin.Logger())
	engine.Use(middleware.LoggerToFile())
	engine.Use(gin.Recovery())
	engine.Use(middleware.RateLimitMiddleware(2*time.Second, 40))

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	engine.GET("/home", api.GetHome)

	engine.GET("/start_webcam", api.StartWebCam)
	engine.GET("/list", api.ListAllCamera)
	engine.GET("/stop_webcam", api.StopWebCam)
	engine.POST("/set_cam/:webcam", api.SelectCamera)

	engine.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})

	authGroup := engine.Group("/auth")
	{
		authGroup.POST("/login", api.Login)
		authGroup.GET("/logout", api.Logout)
		authGroup.POST("/register", api.Register)
	}

}
