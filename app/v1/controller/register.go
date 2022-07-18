package controller

import (
	"github.com/gin-gonic/gin"
	"goserver/app/v1/service"
)

type RegisterRouter struct {}
var registerService service.RegisterService

func (r RegisterRouter) RegisterRouterInit(router *gin.RouterGroup) {
	{
		router.POST("/register", registerService.Register)
	}
}