package routers

import (
	"context"
	"go-line-message-analyzer/internal/app"
	v1 "go-line-message-analyzer/internal/router/api/v1"
	"go-line-message-analyzer/internal/router/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouter(ctx context.Context, app *app.Application) *gin.Engine {
	// docs.SwaggerInfo.BasePath = "/api/v1"

	if viper.GetBool("release") {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())

	// r.Use(log.GinLogger(log.Logger), log.GinRecovery(log.Logger, true))

	//use Cors
	r.Use(middleware.Cors())

	// use swagger (close in ReleaseMode)
	// if gin.Mode() != gin.ReleaseMode {
	// 	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 	// http://127.0.0.1:7777/swagger/index.html
	// }

	// for handler Response
	// responser := NewResponder(ctx)

	// gin.SetMode(setting.RunMode)
	setUpRouter(r, app)

	return r
}

func setUpRouter(router *gin.Engine, app *app.Application) {
	api := router.Group("/api")
	{
		v1.RegisterRouter(api, app)
		//v2.RegisterRouter(api)
	}
}
