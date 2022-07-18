package controller

import (
	"github.com/gin-gonic/gin"
	"goserver/app/v1/service"
)

type CaptchaRouter struct {}

func (c *CaptchaRouter) CaptchaRouterInit(router *gin.RouterGroup) {
	{
		router.GET("/captcha", service.GenerateCaptchaHandler)
	}
}
