package initialize

import (
	"github.com/gin-gonic/gin"
	"goserver/middleware"
	"goserver/router"
	"goserver/utils"
)

func RoutersInit() {
	gin.SetMode(utils.AppMode)
	engine := gin.New()
	routers := new(router.GroupRouters)

	engine.Use(middleware.Logger())
	engine.Use(middleware.Cors())
	//engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user := engine.Group("/api/v1")
	user.Use(middleware.JwtToken())
	{
		routers.UserRoutersInit(user)
		routers.UploadRouterInit(user)
	}
	server := engine.Group("/api/v1")
	{
		routers.LoginRouterInit(server)
		routers.RegisterRouterInit(server)
		routers.CaptchaRouterInit(server)
	}

	err := engine.Run(utils.Host)
	if err != nil {
		return 
	}
}